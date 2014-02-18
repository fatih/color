/*
Package color is an ANSI color package to output colorized or SGR defined
output to the standard output. The api can be used in several way, pick one
one that suits you.

Use simple and default helper functions with predefined foreground colors:

    color.Cyan("Prints text in cyan.")
    color.Blue("Prints %s in blue.", "text") // a newline will be appended automatically

    // These are using by default foreground colors.
    color.Red("We have red")
    color.Yellow("Yellow color too!")
    color.Magenta("And many others ..")


However there are times where custom color mixes are required. Below are some
examples to create custom color objects and use the print functions of each
seperate color object.

    // Create a new color object
    c := color.New(color.FgCyan).Add(color.Underline)
    c.Println("Prints cyan text with an underline.")

    // Or just add them to New()
    d := color.New(color.FgCyan, color.Bold)
    d.Printf("This prints bold cyan %s\n", "too!.")


    // Mix up foreground and background colors, create new mixes!
    red := color.New(color.FgRed)

    boldRed := red.Add(color.Bold)
    boldRed.Println("This will print text in bold red.")

    whiteBackground := red.Add(color.BgWhite)
    whiteBackground.Println("Red text with White background.")


Using with existing color is possible too. Just use the Set() method to set
the standart output to the given parameters. That way a rewrite of an existing
code is not required.

    // Use handy standard colors.
    color.Set(collor.FgYellow)
    fmt.Println("Existing text in your codebase will be now in Yellow")
    fmt.Printf("This one %s\n", "too")
    color.Unset() // don't forget to unset

    // You can mix up parameters
    color.Set(color.FgMagenta, color.Bold)
    defer color.Unset() // use it in your function

    fmt.Println("All text will be now bold magenta.")
*/

package color
