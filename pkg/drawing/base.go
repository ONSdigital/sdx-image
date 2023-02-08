package drawing

import "github.com/fogleman/gg"

type Location struct {
	left, top float64
}

type Dimension struct {
	width, height float64
}

type LengthType int

const (
	Content LengthType = iota
	Proportion
	Pixels
)

type Length struct {
	Value float64
	LengthType
}

func PX(value float64) Length {
	return Length{value, Pixels}
}

func ProportionOfParent(value float64) Length {
	return Length{value, Proportion}
}

func FitContent() Length {
	return Length{0, Content}
}

func MatchParent() Length {
	return Length{1, Proportion}
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
	width, height Length
}

func newBase(width, height Length) *Base {
	return &Base{width, height}
}

func (base *Base) GetWidth(parent Dimension) float64 {
	if base.width.LengthType == Proportion {
		return parent.width * base.width.Value
	}
	return base.width.Value
}

func (base *Base) GetHeight(parent Dimension) float64 {
	if base.height.LengthType == Content {
		return base.height.Value
	} else if base.height.LengthType == Proportion {
		return parent.height * base.height.Value
	}
	return base.height.Value
}
