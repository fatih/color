# Color [![GoDoc](https://godoc.org/github.com/fatih/color?status.png)](http://godoc.org/github.com/fatih/color)

Color let you use colorized outputs in terms of [ASCI Escape
Codes](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors). The api can be
used in several way, pick one one that suits you.


## Install

```bash
go get github.com/fatih/color
```

## Examples

### Standard colors

```go
// Print with default foreground colors
color.Cyan.Print("Prints text in cyan.")
color.Blue.Print("Prints text in blue.")

// Chain SGR paramaters
color.Green.Add(color.Bold).Println("Green with bold")
color.Red.Add(color.BgWhite, color.Underline).Printf("Red with Black background and underscore: %s\n", "format too!")
```

### Custom colors

```go
// Create and reuse color objects
c := color.Cyan.Add(color.Underline)
c.Println("Prints cyan text with an underline.")
c.Printf("Thir prints bold cyan %s\n", "too!.")

// Create custom color objects:
d := color.New(color.FgWhite, color.BgGreen)
d.Println("White with green backround")
```

### Plug into your existing code

```go
// Use handy standard colors.
color.Yellow.Set()
fmt.Println("Existing text in your codebase will be now in Yellow")
fmt.Printf("This one %s\n", "too")
color.Unset() // don't forget to unset

// You can set custom objects too
color.New(color.FgMagenta, color.Bold).Set()
defer color.Unset() // use it in your function

fmt.Println("All text will be now bold magenta.")
```

## Credits

 * [Fatih Arslan](https://github.com/fatih)

## License

The MIT License (MIT) - see LICENSE.md for more details


