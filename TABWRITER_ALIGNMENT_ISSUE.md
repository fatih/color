# ANSI Color Codes Break tabwriter Alignment

## The Problem

When using `text/tabwriter` with ANSI-colored strings, columns become misaligned because tabwriter counts the ANSI escape sequences as visible characters.

### Example of the Bug

```go
package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Fprintln(w, "ID\tSTATE\tPROGRESS")
	fmt.Fprintf(w, "%s\t%s\t%s\n", "abc123", green("completed"), green("100%"))
	fmt.Fprintf(w, "%s\t%s\t%s\n", "def456", yellow("stopped"), "50.0%")
	fmt.Fprintf(w, "%s\t%s\t%s\n", "ghi789", "pending", "-")

	w.Flush()
}
```

### Expected Output (aligned)
```
ID      STATE      PROGRESS
abc123  completed  100%
def456  stopped    50.0%
ghi789  pending    -
```

### Actual Output (misaligned)
```
ID      STATE          PROGRESS
abc123  completed      100%
def456  stopped    50.0%
ghi789  pending    -
```

The "completed" row has extra padding because tabwriter thinks `\x1b[32mcompleted\x1b[0m` (19 bytes) is longer than `stopped` (7 bytes), when visually they're 9 and 7 characters respectively.

## Why This Happens

ANSI escape sequences for colors look like:
- `\x1b[32m` - Start green color (5 bytes)
- `\x1b[0m` - Reset color (4 bytes)

So a green "completed" string is actually:
```
\x1b[32mcompleted\x1b[0m
```
That's 9 visible characters but 18 bytes total.

`tabwriter` uses byte length to calculate column widths, not visible character count. It has no awareness of ANSI escape sequences.

## The Fix (Manual Padding)

The workaround is to bypass tabwriter and calculate padding manually, accounting for the ANSI overhead:

```go
package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	// Print header with fixed widths
	fmt.Printf("%-10s %-12s %-10s\n", "ID", "STATE", "PROGRESS")

	// For each row, calculate padding that accounts for ANSI codes
	rows := []struct {
		id       string
		stateRaw string
		stateFn  func(a ...interface{}) string
		progress string
	}{
		{"abc123", "completed", green, "100%"},
		{"def456", "stopped", yellow, "50.0%"},
		{"ghi789", "pending", nil, "-"},
	}

	for _, r := range rows {
		stateStr := r.stateRaw
		if r.stateFn != nil {
			stateStr = r.stateFn(r.stateRaw)
		}

		// Calculate padding: desired_width - visible_length + actual_length
		// This compensates for the invisible ANSI bytes
		statePad := 12 - len(r.stateRaw) + len(stateStr)

		fmt.Printf("%-10s %-*s %-10s\n", r.id, statePad, stateStr, r.progress)
	}
}
```

### How the Padding Math Works

For a column width of 12:
- Plain "pending" (7 chars): `%-12s` pads to 12, adds 5 spaces ✓
- Colored "completed" (9 visible, but 18 bytes):
  - `%-12s` would only add 0 spaces (12-18 < 0, no padding)
  - We need: `%-21s` (12 - 9 + 18 = 21) to get correct 3 spaces of padding

Formula: `padWidth = desiredWidth - visibleLength + byteLength`

## Potential Solutions for the color Library

### Option 1: Add a PaddedSprint Function

```go
// PaddedSprint returns a colored string padded to the specified visible width
func (c *Color) PaddedSprint(width int, a ...interface{}) string {
    raw := fmt.Sprint(a...)
    colored := c.Sprint(a...)

    visibleLen := len(raw)
    if visibleLen >= width {
        return colored
    }

    padding := strings.Repeat(" ", width-visibleLen)
    return colored + padding
}

// Usage:
green.PaddedSprint(12, "completed") // "completed   " with green color
```

### Option 2: Add a helper to calculate pad width

```go
// PadWidth returns the width needed for fmt's %-*s to achieve the desired
// visible width when the string contains ANSI codes
func (c *Color) PadWidth(desiredWidth int, a ...interface{}) int {
    raw := fmt.Sprint(a...)
    colored := c.Sprint(a...)
    return desiredWidth - len(raw) + len(colored)
}

// Usage:
width := green.PadWidth(12, "completed")
fmt.Printf("%-*s", width, green.Sprint("completed"))
```

### Option 3: ANSI-aware tabwriter wrapper

Create a wrapper that strips ANSI codes when calculating widths but preserves them in output. This is more complex but would be a drop-in replacement for tabwriter.

### Option 4: Utility function to strip ANSI codes

```go
// VisibleLength returns the length of a string excluding ANSI escape sequences
func VisibleLength(s string) int {
    // Regex: \x1b\[[0-9;]*m
    ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
    return len(ansiRegex.ReplaceAllString(s, ""))
}
```

## Real-World Code That Was Fixed

From `rsssyncer` project - the table display for download transfers:

```go
// Before (broken):
w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
fmt.Fprintf(w, "%s\t%s\t%s\n", id, green("completed"), green("100%"))
w.Flush()

// After (fixed):
stateRaw := "completed"
stateStr := green(stateRaw)
statePad := 12 - len(stateRaw) + len(stateStr)
fmt.Printf("%-10s %-*s %-10s\n", id, statePad, stateStr, progress)
```

## References

- Go tabwriter package: https://pkg.go.dev/text/tabwriter
- ANSI escape codes: https://en.wikipedia.org/wiki/ANSI_escape_code
- Related issue in other projects: This is a common problem when mixing tabwriter with any ANSI coloring library
