package color

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	Cyan.Print("Prints text in cyan.")
	Blue.Print("Prints text in blue.")

	// Chain SGR paramaters
	Green.Add(Bold).Println("Green with bold")
	Red.Add(BgWhite, Underline).Printf("Red with White background and underscore: %s\n", "format too!")

	// Create and reuse color objects
	c := Cyan.Add(Underline)
	c.Println("Prints bold cyan.")
	c.Printf("Thir prints bold cyan %s\n", "too!.")

	// Create custom color objects:
	d := New(FgGreen, BgCyan, Italic)
	d.Print("Italic green with cyan backround")

	// You can use set custom objects too
	Cyan.Set()
	fmt.Println("Existing text in your codebase will be now in Cyan")
	fmt.Printf("This one %s\n", "too")
	Unset() // don't forget to unset

	// You can use set custom objects too
	New(FgBlack, BgWhite, Bold).Set()
	defer Unset() // use it in your function

	fmt.Println("All text will be now bold red with white background.")

}
