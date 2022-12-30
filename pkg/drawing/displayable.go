package drawing

type Displayable interface {
	GetWidth(parent Dimension) float64
	GetHeight(parent Dimension) float64
	Render(area Rectangle)
}

func Display(displayable Displayable, left, top float64) {
	outer := newRectangle(
		left,
		top,
		displayable.GetWidth(Dimension{0, 0}),
		displayable.GetHeight(Dimension{0, 0}))

	displayable.Render(outer)
}
