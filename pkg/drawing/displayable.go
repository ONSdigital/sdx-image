package drawing

import "github.com/fogleman/gg"

type Displayable interface {
	GetWidth() float64
	GetHeight() float64
	Render(location Rectangle, context *gg.Context) Rectangle
}

func Display(displayable Displayable, left, top float64, context *gg.Context) {
	outer := Rectangle{left: left, top: top, width: displayable.GetWidth(), height: displayable.GetHeight()}
	displayable.Render(outer, context)
}
