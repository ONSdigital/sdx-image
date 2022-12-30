package drawing

import (
	"github.com/fogleman/gg"
	"image/color"
)

type Div struct {
	*Base
	BackgroundColor color.Color
	BorderColor     color.Color
	BorderWeight    float64
}

func newDiv(width, height float64, context *gg.Context) *Div {
	return &Div{
		Base:            newBase(width, height, context),
		BackgroundColor: nil,
		BorderColor:     nil,
		BorderWeight:    0.0}
}

func (div *Div) getInternalDim(parent Dimension) Dimension {
	w := div.GetWidth(parent)
	h := div.GetHeight(parent)
	b := div.BorderWeight
	return Dimension{w - 2*b, h - 2*b}
}

func (div *Div) getInternalArea(area Rectangle) Rectangle {
	l := area.left
	t := area.top
	w := area.width
	h := area.height
	b := div.BorderWeight
	return newRectangle(l+b, t+b, w-2*b, h-2*b)
}

func (div *Div) Render(area Rectangle) {
	context := div.context
	l := area.left
	t := area.top
	w := area.width
	h := area.height
	b := div.BorderWeight

	if div.BackgroundColor != nil {
		context.DrawRectangle(l, t, w, h)
		context.SetColor(color.White)
		context.FillPreserve()
		context.SetColor(div.BackgroundColor)
		context.Fill()
	}

	//draw border
	if div.BorderWeight > 0 {
		context.SetColor(div.BorderColor)
		context.DrawRectangle(l, t, b, h)
		context.Fill()
		context.DrawRectangle(l+w-b, t, b, h)
		context.Fill()
		context.DrawRectangle(l, t, w, b)
		context.Fill()
		context.DrawRectangle(l, t+h-b, w, b)
		context.Fill()
	}
}
