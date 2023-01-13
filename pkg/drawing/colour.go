package drawing

import (
	"github.com/fogleman/gg"
)

type Colour interface {
	R() int
	G() int
	B() int
}

type Color struct {
	r, g, b int
}

func (color Color) R() int {
	return color.r
}

func (color Color) G() int {
	return color.g
}

func (color Color) B() int {
	return color.b
}

func CreateColor(r, g, b int) Colour {
	return Color{r, g, b}
}

func setColour(c Colour, context *gg.Context) {
	context.SetRGB255(c.R(), c.G(), c.B())
}

var WHITE = Color{255, 255, 255}
var BLACK = Color{0, 0, 0}
