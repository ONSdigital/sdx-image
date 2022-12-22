package drawing

import (
	"github.com/fogleman/gg"
)

type Text struct {
	*Div
	value string
	size  int
}

func newText(value string, size int, width, height float64) *Text {
	return &Text{value: value, size: size, Div: newDiv(width, height)}
}

func (text *Text) GetWidth() float64 {
	return text.width
}

func (text *Text) GetHeight() float64 {
	return text.height
}

func (text *Text) Render(location Rectangle, context *gg.Context) Rectangle {
	return location
}
