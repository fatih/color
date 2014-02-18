// Package color is an ANSI color package to output colorized or SGR defined
// output to the standard output.
package color

import (
	"io"
	"os"
	"strconv"
	"strings"

	"fmt"
)

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params []Parameter
}

// Parameter defines a single SGR Code
type Parameter int

const escape = "\x1b"

// All attributes off
const (
	Reset Parameter = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Parameter = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Background text colors
const (
	BgBlack Parameter = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// New returns a newly created color object.
func New(value ...Parameter) *Color {
	c := &Color{params: make([]Parameter, 0)}
	c.Add(value...)
	return c
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.
func Set(p ...Parameter) *Color {
	c := New(p...)
	c.Set()
	return c
}

// Unset resets all escape attributes and clears the output. Usualy should
// be called after Set().
func Unset() {
	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

// Add is used to chain SGR parameters. Use as many as paramters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *Color) Add(value ...Parameter) *Color {
	c.params = append(c.params, value...)
	return c
}

func (c *Color) prepend(value Parameter) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Output defines the standard output of the print functions. By default
// os.Stdout is used.
var Output io.Writer = os.Stdout

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// This is the standard fmt.Printf() method wrapped with the given color.
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprintf(Output, format, a...)
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprint(Output, a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Fprintln(Output, a...)
}

// sequence returns a formated SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// Set sets the SGR sequence.
func (c *Color) Set() *Color {
	fmt.Fprintf(Output, "%s[%sm", escape, c.sequence())
	return c
}

// Black is an convenient helper function to print with black foreground. A
// newline is appended to format by default.
func Black(format string, a ...interface{}) { printColor(format, FgBlack, a...) }

// Red is an convenient helper function to print with red foreground. A
// newline is appended to format by default.
func Red(format string, a ...interface{}) { printColor(format, FgRed, a...) }

// Green is an convenient helper function to print with green foreground. A
// newline is appended to format by default.
func Green(format string, a ...interface{}) { printColor(format, FgGreen, a...) }

// Yellow is an convenient helper function to print with yello foreground.
// A newline is appended to format by default.
func Yellow(format string, a ...interface{}) { printColor(format, FgYellow, a...) }

// Blue is an convenient helper function to print with blue foreground. A
// newline is appended to format by default.
func Blue(format string, a ...interface{}) { printColor(format, FgBlue, a...) }

// Magenta is an convenient helper function to print with magenta foreground.
// A newline is appended to format by default.
func Magenta(format string, a ...interface{}) { printColor(format, FgMagenta, a...) }

// Cyan is an convenient helper function to print with cyan foreground. A
// newline is appended to format by default.
func Cyan(format string, a ...interface{}) { printColor(format, FgCyan, a...) }

// White is an convenient helper function to print with white foreground. A
// newline is appended to format by default.
func White(format string, a ...interface{}) { printColor(format, FgWhite, a...) }

func printColor(format string, p Parameter, a ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	c := &Color{params: []Parameter{p}}
	c.Printf(format, a...)
}
