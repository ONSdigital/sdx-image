package drawing

import "github.com/fogleman/gg"

type Location struct {
	left, top float64
}

type Dimension struct {
	width, height float64
}

type Rectangle struct {
	Location
	Dimension
}

func newRectangle(left, top, width, height float64) Rectangle {
	return Rectangle{Location{left, top}, Dimension{width, height}}
}

type Displayable interface {
	GetWidth(parent Dimension) float64
	GetHeight(parent Dimension) float64
	Render(area Rectangle, context *gg.Context)
}

type Base struct {
	Dimension // 0 fit content, 0 < x <= 1 proportion of parent, 1 < x length in px
}

func newBase(width, height float64) *Base {
	return &Base{Dimension: Dimension{width, height}}
}

func (base *Base) GetWidth(parent Dimension) float64 {
	if base.width <= 1 {
		return parent.width * base.width
	}
	return base.width
}

func (base *Base) GetHeight(parent Dimension) float64 {
	if base.height == 0 {
		return base.height
	} else if base.height <= 1 {
		return parent.height * base.height
	}
	return base.height
}
