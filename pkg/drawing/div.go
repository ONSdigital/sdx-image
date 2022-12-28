package drawing

import (
	"github.com/fogleman/gg"
	"image/color"
)

type Div struct {
	*Base
	backgroundColor color.Color
	borderColor     color.Color
	borderWeight    float64
}

func newDiv(width, height float64, context *gg.Context) *Div {
	return &Div{
		Base:            newBase(width, height, context),
		backgroundColor: nil,
		borderColor:     nil,
		borderWeight:    0.0}
}

func (div *Div) Render(location Rectangle) Rectangle {
	context := div.context
	l := location.left
	t := location.top
	w := location.width
	h := location.height
	b := div.borderWeight

	if div.backgroundColor != nil {
		context.DrawRectangle(l, t, w, h)
		context.SetColor(color.White)
		context.FillPreserve()
		context.SetColor(div.backgroundColor)
		context.Fill()
	}

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
