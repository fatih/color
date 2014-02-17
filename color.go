package color

import (
	"strconv"
	"strings"

	"fmt"
)

type Color struct {
	params []Parameter
}

// Parameter defines a single SGR Code
type Parameter int

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

var (
	Black   = &Color{params: []Parameter{FgBlack}}
	Red     = &Color{params: []Parameter{FgRed}}
	Green   = &Color{params: []Parameter{FgGreen}}
	Yellow  = &Color{params: []Parameter{FgYellow}}
	Blue    = &Color{params: []Parameter{FgBlue}}
	Magenta = &Color{params: []Parameter{FgMagenta}}
	Cyan    = &Color{params: []Parameter{FgCyan}}
	White   = &Color{params: []Parameter{FgWhite}}
)

func New(value ...Parameter) *Color {
	c := &Color{params: make([]Parameter, 0)}
	c.Add(value...)
	return c
}

func (c *Color) Bold() *Color {
	c.Add(Bold)
	return c
}

func (c *Color) Add(value ...Parameter) *Color {
	c.params = append(c.params, value...)
	return c
}

func (c *Color) prepend(value Parameter) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Printf(format, a...)
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered.
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Print(a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered.
func (c *Color) Println(a ...interface{}) (n int, err error) {
	c.Set()
	defer Unset()

	return fmt.Println(a...)
}

// sequence returns a formated SGR sequence to be plugged into a "\033[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// Set sets the SGR sequence.
func (c *Color) Set() {
	fmt.Printf("\033[%sm", c.sequence())
}

func Unset() {
	fmt.Print("\033[0m")
}
