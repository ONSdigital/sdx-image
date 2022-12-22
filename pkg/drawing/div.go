package drawing

import (
	"github.com/fogleman/gg"
	"image/color"
)

type Div struct {
	width, height   float64 //0 fit content, 0 < x <= 1, proportion of parent, 1 < x length in px
	backgroundColor color.Color
	borderColor     color.Color
	borderWeight    float64
}

func newDiv(width, height float64) *Div {
	return &Div{width: width, height: height, backgroundColor: color.White, borderColor: color.White}
}

func (div *Div) GetWidth() float64 {
	return div.width
}

func (div *Div) GetHeight() float64 {
	return div.height
}

func (div *Div) Render(location Rectangle, context *gg.Context) Rectangle {
	l := location.left
	t := location.top
	w := location.width
	h := location.height
	b := div.borderWeight

	context.DrawRectangle(l, t, w, h)
	context.SetColor(color.White)
	context.FillPreserve()
	context.SetColor(div.backgroundColor)
	context.Fill()

	//draw border
	if div.borderWeight > 0 {
		context.SetColor(div.borderColor)
		context.DrawRectangle(l, t, b, h)
		context.Fill()
		context.DrawRectangle(l+w-b, t, b, h)
		context.Fill()
		context.DrawRectangle(l, t, w, b)
		context.Fill()
		context.DrawRectangle(l, t+h-b, w, b)
		context.Fill()
	}

	return Rectangle{l + b, t + b, w - 2*b, h - 2*b}
}
