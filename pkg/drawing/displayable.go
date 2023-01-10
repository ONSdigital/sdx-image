package drawing

import "github.com/fogleman/gg"

type Displayable interface {
	GetWidth(parent Dimension) float64
	GetHeight(parent Dimension) float64
	Render(area Rectangle, context *gg.Context)
}
