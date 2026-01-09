package tabwriter

import (
	"bytes"
	"strings"
	"testing"
	stdtabwriter "text/tabwriter"
)

func stripANSI(input string) string {
	var out strings.Builder
	out.Grow(len(input))
	state := ansiNone
	for i := 0; i < len(input); i++ {
		ch := input[i]
		if state != ansiNone {
			switch state {
			case ansiEsc:
				switch ch {
				case '[':
					state = ansiCSI
				case ']', 'P', '^', '_':
					state = ansiString
				default:
					state = ansiNone
				}
			case ansiCSI:
				if ch >= 0x40 && ch <= 0x7e {
					state = ansiNone
				}
			case ansiString:
				if ch == 0x07 {
					state = ansiNone
				} else if ch == ansiEscape {
					state = ansiStringEsc
				}
			case ansiStringEsc:
				if ch == '\\' {
					state = ansiNone
				} else {
					state = ansiString
				}
			}
			continue
		}

		if ch == ansiEscape {
			state = ansiEsc
			continue
		}
		out.WriteByte(ch)
	}
	return out.String()
}

func TestANSIAlignmentMatchesStripped(t *testing.T) {
	chunks := []string{
		"Status\tMessage\tCode\n",
		"\x1b[",
		"31mFAIL",
		"\x1b[0m\t",
		"something went wrong\t",
		"500\n",
		"OK\tall good\t200\n",
	}

	var got bytes.Buffer
	w := NewWriter(&got, 0, 0, 2, ' ', 0)
	for _, chunk := range chunks {
		if _, err := w.Write([]byte(chunk)); err != nil {
			t.Fatalf("write failed: %v", err)
		}
	}
	if err := w.Flush(); err != nil {
		t.Fatalf("flush failed: %v", err)
	}

	if !strings.Contains(got.String(), "\x1b[31m") || !strings.Contains(got.String(), "\x1b[0m") {
		t.Fatalf("expected ANSI sequences to remain in output")
	}

	var want bytes.Buffer
	w2 := stdtabwriter.NewWriter(&want, 0, 0, 2, ' ', 0)
	plain := stripANSI(strings.Join(chunks, ""))
	if _, err := w2.Write([]byte(plain)); err != nil {
		t.Fatalf("write (stdlib) failed: %v", err)
	}
	if err := w2.Flush(); err != nil {
		t.Fatalf("flush (stdlib) failed: %v", err)
	}

	gotStripped := stripANSI(got.String())
	if gotStripped != want.String() {
		t.Fatalf("alignment mismatch\n--- got ---\n%s\n--- want ---\n%s", gotStripped, want.String())
	}
}
