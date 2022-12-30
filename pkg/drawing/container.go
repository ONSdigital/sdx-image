package drawing

import "github.com/fogleman/gg"

type Layout int

const (
	Row Layout = iota
	Column
)

type Justification int

const (
	JustifyStart Justification = iota
	JustifyEnd
	JustifySpaced
)

type Alignment int

const (
	AlignStart Alignment = iota
	AlignEnd
	AlignCenter
)

type Container struct {
	*Div
	children       []Displayable
	layout         Layout
	justifyContent Justification
	alignItems     Alignment
	padding        float64
}

func newContainer(width, height float64, context *gg.Context) *Container {
	return &Container{
		Div:            newDiv(width, height, context),
		layout:         Row,
		justifyContent: JustifyStart,
		alignItems:     AlignStart,
		padding:        0.0}
}

func (container *Container) GetHeight(parent Dimension) float64 {
	height := container.Div.GetHeight(parent)
	if height == 0 {
		internal := container.Div.getInternalDim(parent)
		if container.layout == Column {
			height = container.getTotalChildHeight(internal)
		} else {
			height = container.getLargestChildHeight(internal)
		}
		height += 2 * container.borderWeight
	}
	return height
}

func (container *Container) Render(area Rectangle) {
	internalArea := container.Div.getInternalArea(area)
	container.Div.Render(area)
	container.renderChildren(internalArea)
}

func (container *Container) renderChildren(area Rectangle) {
	left := area.left
	top := area.top

	w := 0.0
	h := 0.0
	wGap := 0.0
	hGap := 0.0

	if container.justifyContent == JustifyEnd {
		if container.layout == Row {
			totalWidth := container.getTotalChildWidth(area.Dimension)
			w = area.width - totalWidth
		} else if container.layout == Column {
			totalHeight := container.getTotalChildHeight(area.Dimension)
			h = area.height - totalHeight
		}
	} else if container.justifyContent == JustifySpaced {
		if container.layout == Row {
			totalWidth := container.getTotalChildWidth(area.Dimension)
			wGap = (area.width - totalWidth) / float64(len(container.children)+1)
			w = wGap
		} else if container.layout == Column {
			totalHeight := container.getTotalChildHeight(area.Dimension)
			hGap = (area.height - totalHeight) / float64(len(container.children)+1)
			h = hGap
		}
	}

	for _, child := range container.children {
		width := child.GetWidth(area.Dimension)
		height := child.GetHeight(area.Dimension)

		if container.alignItems == AlignEnd {
			if container.layout == Row {
				h = area.height - height
			} else if container.layout == Column {
				w = area.width - width
			}
		} else if container.alignItems == AlignCenter {
			if container.layout == Row {
				h = area.height/2 - height/2
			} else if container.layout == Column {
				w = area.width/2 - width/2
			}
		}

		childArea := newRectangle(left+container.padding+w, top+container.padding+h, width, height)
		child.Render(childArea)
		if container.layout == Row {
			w += width + wGap
		} else if container.layout == Column {
			h += height + hGap
		}
	}
}

func (container *Container) getTotalChildWidth(dimension Dimension) float64 {
	width := 0.0
	for _, child := range container.children {
		width += child.GetWidth(dimension)
	}
	return width
}

func (container *Container) getTotalChildHeight(dimension Dimension) float64 {
	height := 0.0
	for _, child := range container.children {
		height += child.GetHeight(dimension)
	}
	return height
}

func (container *Container) getLargestChildHeight(dimension Dimension) float64 {
	largest := 0.0
	for _, child := range container.children {
		height := child.GetHeight(dimension)
		if height > largest {
			largest = height
		}
	}
	return largest
}
