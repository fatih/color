package color

import (
	"fmt"
	"testing"
)

// The example from the standart library doesn't work unfortunaly.
func TestColor(t *testing.T) {

	New(FgRed).Printf("red\t")
	New(BgRed).Print("         ")
	New(FgRed, Bold).Println(" red")

	New(FgGreen).Printf("green\t")
	New(BgGreen).Print("         ")
	New(FgGreen, Bold).Println(" green")

	New(FgYellow).Printf("yellow\t")
	New(BgYellow).Print("         ")
	New(FgYellow, Bold).Println(" yellow")

	New(FgBlue).Printf("blue\t")
	New(BgBlue).Print("         ")
	New(FgBlue, Bold).Println(" blue")

	New(FgMagenta).Printf("magenta\t")
	New(BgMagenta).Print("         ")
	New(FgMagenta, Bold).Println(" magenta")

	New(FgCyan).Printf("cyan\t")
	New(BgCyan).Print("         ")
	New(FgCyan, Bold).Println(" cyan")

	New(FgWhite).Printf("white\t")
	New(BgWhite).Print("         ")
	New(FgWhite, Bold).Println(" white")

	fmt.Println("")

	Cyan.Print("Prints text in cyan. ")
	Blue.Print("Prints text in blue. ")

	// Chain SGR paramaters
	Green.Add(Bold).Println("Green with bold")
	Red.Add(BgWhite, Underline).Println("Red with White background and underscore")

	// Create and reuse color objects
	c := New(FgRed).Add(Underline)
	c.Println("Prints red text with an underline.")

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

	fmt.Println("All text will be now bold magenta.")

}
