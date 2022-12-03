package color

import (
	"os"

	"golang.org/x/sys/windows"
)

func init() {
	// Opt-in for ansi color support for current process.
	// https://learn.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences#output-sequences
	var outMode uint32
	out := windows.Handle(os.Stdout.Fd())
	if err := windows.GetConsoleMode(out, &outMode); err == nil {
		_ = windows.SetConsoleMode(out, outMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	}
}
