package drawing

import "github.com/fogleman/gg"

type Rectangle struct {
	left, top, width, height float64
}

type Base struct {
	width, height float64 //0 fit content, 0 < x <= 1, proportion of parent, 1 < x length in px
	context       *gg.Context
}

func newBase(width, height float64, context *gg.Context) *Base {
	return &Base{width: width, height: height, context: context}
}

func (base *Base) GetWidth(location Rectangle) float64 {
	if base.width <= 1 {
		return location.width * base.width
	}
	return base.width
}

func (base *Base) GetHeight(location Rectangle) float64 {
	if base.height <= 1 {
		return location.height * base.height
	}
	return base.height
}

func (base *Base) Render(location Rectangle) Rectangle {
	return location
}
