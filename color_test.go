package color

import (
	"fmt"
	"testing"
)

// The example from the standart library doesn't work unfortunaly.
func TestColor(t *testing.T) {
	Cyan.Print("Prints text in cyan.")
	Blue.Print("Prints text in blue.")

	// Chain SGR paramaters
	Green.Add(Bold).Println("Green with bold")
	Red.Add(BgWhite, Underline).Printf("Red with White background and underscore: %s\n", "format too!")

	// Create and reuse color objects
	c := Cyan.Add(Underline)
	c.Println("Prints cyan text with an underline.")
	c.Printf("Thir prints bold cyan %s\n", "too!.")

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
