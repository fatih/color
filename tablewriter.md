# Tabwriter alignment plan

## Context
- Reported issue in this repo: #241 "Library not compatible with text/tabwriter" (Oct 3, 2024).
- Goal: provide a drop-in tabwriter that ignores ANSI escape sequences when measuring column widths, without changing the existing color APIs.

## Drop-in replacements found (external)
1) go.followtheprocess.codes/hue/tabwriter
   - Drop-in replacement for text/tabwriter that ignores ANSI codes for width calculations.
   - Usage should be identical to text/tabwriter (import path change only).

2) github.com/juju/ansiterm (TabWriter)
   - NewTabWriter returns a tabwriter that strips ANSI escape codes for width calculations.
   - API is very similar to text/tabwriter but uses ansiterm.NewTabWriter.

3) github.com/kravlad/ansiterm (TabWriter)
   - Same TabWriter API as juju/ansiterm with ANSI-stripping width logic.

## Plan: add an in-repo fork at github.com/fatih/color/tabwriter
1) Fork stdlib text/tabwriter from Go 1.24.1 into a new package directory: tabwriter/.
2) Keep the API identical to stdlib so callers only change the import path.
3) Modify width calculations to ignore ANSI escape sequences while preserving output bytes:
   - Skip CSI sequences: ESC [ ... final (0x40-0x7E).
   - Skip OSC sequences: ESC ] ... terminated by BEL or ESC \\.
   - Treat malformed/incomplete sequences as visible bytes to avoid hiding user data.
4) Tests:
   - ANSI SGR does not affect alignment (colored vs plain columns align).
   - Mixed ANSI/plain output remains stable.
   - Incomplete escape sequence at end does not panic.
   - Output still includes the original ANSI bytes (only width calc changes).
5) Docs:
   - Note the new package in README and TABWRITER_ALIGNMENT_ISSUE.md.
   - Provide a short usage example.

## Example usage (drop-in replacements)
```go
package main

import (
    "fmt"
    "os"

    // Option A: planned in-repo fork
    ctab "github.com/fatih/color/tabwriter"

    // Option B: external drop-in
    // htab "go.followtheprocess.codes/hue/tabwriter"

    // Option C: external drop-in with a different constructor name
    // "github.com/juju/ansiterm"
)

func main() {
    w := ctab.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
    // or: w := htab.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
    // or: w := ansiterm.NewTabWriter(os.Stdout, 0, 0, 2, ' ', 0)

    fmt.Fprintln(w, "ID\tSTATE\tPROGRESS")
    fmt.Fprintln(w, "abc123\t\x1b[32mcompleted\x1b[0m\t\x1b[32m100%\x1b[0m")
    fmt.Fprintln(w, "def456\t\x1b[33mstopped\x1b[0m\t50.0%")
    _ = w.Flush()
}
```
