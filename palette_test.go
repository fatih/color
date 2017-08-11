package color

import (
	"fmt"
	"testing"
)

func TestPalette(t *testing.T) {
	p := NewPalette(
		FgRed,
		FgGreen,
		FgBlue,
	)

	colors := []*Color{
		New(FgRed),
		New(FgGreen),
		New(FgBlue),
		New(FgRed),
		New(FgGreen),
		New(FgBlue),
	}

	for _, c := range colors {
		got := p.Next()
		expected := c

		fmt.Printf("%s %s\n", got.Sprintf("got"), expected.Sprintf("expected"))
		if !expected.Equals(got) {
			t.Errorf("Palette's Next returned bad value. Expected: %s, Got: %s", expected, got)
		}
	}
}

func TestPaletteLen(t *testing.T) {
	p := NewPalette()
	if p.Len() != 0 {
		t.Errorf("Palette should report 0 length if initialized with no attributes")
	}

	p = NewPalette(
		FgRed,
		FgGreen,
		FgBlue,
	)
	if p.Len() != 3 {
		t.Errorf("Palette should have the correct length")
	}
}
