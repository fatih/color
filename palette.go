package color

// Pallete provides a way to manage a collection of colors. It allows you to do
// things like iterate over the colors. This can be useful for coloring the
// output of commands to differentiate between different streams.
type Palette struct {
  curIndex int
	colors   []*Color
}

// NewPalette creates a new palette from the attributes passed in.
func NewPalette(attrs ...Attribute) *Palette {
  colors := make([]*Color, len(attrs))
  for i, a := range attrs {
    colors[i] = New(a)
  }

  return &Palette{
    curIndex: 0,
    colors: colors,
  }
}

// Next returns the next color in the palette.
func (p *Palette) Next() *Color {
	defer func() {
		p.curIndex++
		if p.curIndex >= len(p.colors) {
			p.curIndex = 0
		}
	}()

	return p.colors[p.curIndex]
}

// Length returns the number of colors in the palette.
func (p *Palette) Len() int {
  return len(p.colors)
}