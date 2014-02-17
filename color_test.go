package color

import (
	"fmt"
	"testing"
)

// The example from the standart library doesn't work unfortunaly.
func TestColor(t *testing.T) {
	fmt.Printf("black")
	New(BgBlack).Println("         ")
	fmt.Printf("red\t")
	New(BgRed).Println("         ")
	fmt.Printf("green\t")
	New(BgGreen).Println("         ")
	fmt.Printf("yellow\t")
	New(BgYellow).Println("         ")
	fmt.Printf("blue\t")
	New(BgBlue).Println("         ")
	fmt.Printf("magenta\t")
	New(BgMagenta).Println("         ")
	fmt.Printf("cyan\t")
	New(BgCyan).Println("         ")
	fmt.Printf("white\t")
	New(BgWhite).Println("         ")

	fmt.Println("")

	Cyan.Print("Prints text in cyan. ")
	Blue.Print("Prints text in blue. ")

	// Chain SGR paramaters
	Green.Add(Bold).Println("Green with bold")
	Red.Add(BgWhite, Underline).Println("Red with White background and underscore")

	// Create and reuse color objects
	c := Red.Add(Underline)
	fmt.Println(c.params)
	c.Println("Prints cyan text with an underline.")

	// Create custom color objects:
	d := New(FgWhite, BgGreen)
	d.Println("White with green backround")

	// You can use set custom objects too
	Yellow.Set()
	fmt.Println("Existing text in your codebase will be now in Yellow")
	fmt.Printf("This one %s\n", "too")
	Unset() // don't forget to unset

	// You can use set custom objects too
	New(FgMagenta, Bold).Set()
	defer Unset() // use it in your function

	fmt.Println("All text will be now bold red with white background.")

}
