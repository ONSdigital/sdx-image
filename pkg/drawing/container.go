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
	return &Container{Div: newDiv(width, height, context), layout: Row, justifyContent: JustifyStart, padding: 0.0}
}

func (container *Container) Render(location Rectangle) Rectangle {
	rect := container.Div.Render(location)
	container.renderChildren(rect)
	return rect
}

func (container *Container) renderChildren(location Rectangle) {
	left := location.left
	top := location.top

	w := 0.0
	h := 0.0
	wGap := 0.0
	hGap := 0.0

	if container.justifyContent == JustifyEnd {
		if container.layout == Row {
			totalWidth := container.getTotalChildWidth(location)
			w = location.width - totalWidth
		} else if container.layout == Column {
			totalHeight := container.getTotalChildHeight(location)
			h = location.height - totalHeight
		}
	} else if container.justifyContent == JustifySpaced {
		if container.layout == Row {
			totalWidth := container.getTotalChildWidth(location)
			wGap = (location.width - totalWidth) / float64(len(container.children)+1)
			w = wGap
		} else if container.layout == Column {
			totalHeight := container.getTotalChildHeight(location)
			hGap = (location.height - totalHeight) / float64(len(container.children)+1)
			h = hGap
		}
	}

	for _, child := range container.children {
		width := child.GetWidth(location)
		height := child.GetHeight(location)

		if container.alignItems == AlignEnd {
			if container.layout == Row {
				h = location.height - height
			} else if container.layout == Column {
				w = location.width - width
			}
		} else if container.alignItems == AlignCenter {
			if container.layout == Row {
				h = location.height/2 - height/2
			} else if container.layout == Column {
				w = location.width/2 - width/2
			}
		}

		childLocation := Rectangle{left + container.padding + w, top + container.padding + h, width, height}
		child.Render(childLocation)
		if container.layout == Row {
			w += width + wGap
		} else if container.layout == Column {
			h += height + hGap
		}
	}
}

func (container *Container) getTotalChildWidth(location Rectangle) float64 {
	width := 0.0
	for _, child := range container.children {
		width += child.GetWidth(location)
	}
	return width
}

func (container *Container) getTotalChildHeight(location Rectangle) float64 {
	height := 0.0
	for _, child := range container.children {
		height += child.GetHeight(location)
	}
	return height
}
