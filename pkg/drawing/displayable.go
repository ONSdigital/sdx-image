package drawing

type Displayable interface {
	GetWidth(location Rectangle) float64
	GetHeight(location Rectangle) float64
	Render(location Rectangle) Rectangle
}

func Display(displayable Displayable, left, top float64) {
	outer := Rectangle{
		left:   left,
		top:    top,
		width:  displayable.GetWidth(Rectangle{0, 0, 0, 0}),
		height: displayable.GetHeight(Rectangle{0, 0, 0, 0})}

	displayable.Render(outer)
}
