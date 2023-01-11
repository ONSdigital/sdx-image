package drawing

import (
	"github.com/fogleman/gg"
)

type Div struct {
	*Base
	backgroundColor Colour
	borderColor     Colour
	borderWeight    float64
}

func newDiv(width, height float64) *Div {
	return &Div{
		Base:            newBase(width, height),
		backgroundColor: nil,
		borderColor:     nil,
		borderWeight:    0.0}
}

func (div *Div) SetBackgroundColor(colour Colour) {
	div.backgroundColor = colour
}

func (div *Div) SetBorder(colour Colour, weight float64) {
	div.borderColor = colour
	div.borderWeight = weight
}

func (div *Div) getInternalDim(parent Dimension) Dimension {
	w := div.GetWidth(parent)
	h := div.GetHeight(parent)
	b := div.borderWeight
	return Dimension{w - 2*b, h - 2*b}
}

func (div *Div) getInternalArea(area Rectangle) Rectangle {
	l := area.left
	t := area.top
	w := area.width
	h := area.height
	b := div.borderWeight
	return newRectangle(l+b, t+b, w-2*b, h-2*b)
}

func (div *Div) Render(area Rectangle, context *gg.Context) {
	l := area.left
	t := area.top
	w := area.width
	h := area.height
	b := div.borderWeight

	if div.backgroundColor != nil {
		context.DrawRectangle(l, t, w, h)
		context.SetRGB255(255, 255, 255)
		context.FillPreserve()
		setColour(div.backgroundColor, context)
		context.Fill()
	}

	//draw border
	if div.borderWeight > 0 {
		setColour(div.borderColor, context)
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
