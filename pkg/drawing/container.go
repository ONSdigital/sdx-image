package drawing

import (
	"github.com/fogleman/gg"
)

type Layout int

const (
	ROW Layout = iota
	COLUMN
)

type Justification int

const (
	START Justification = iota
	END
	SPACED
)

type Alignment int

const (
	ALIGN_START Alignment = iota
	ALIGN_END
	ALIGN_CENTER
)

type Container struct {
	*Div
	children       []Displayable
	layout         Layout
	justifyContent Justification
	alignItems     Alignment
	padding        float64
}

func newContainer(width, height float64) *Container {
	return &Container{Div: newDiv(width, height), layout: ROW, justifyContent: START, padding: 0.0}
}

func (container *Container) GetWidth() float64 {
	return container.width
}

func (container *Container) GetHeight() float64 {
	return container.height
}

func (container *Container) Render(location Rectangle, context *gg.Context) Rectangle {
	rect := container.Div.Render(location, context)
	container.renderChildren(rect, context)
	return rect
}

func (container *Container) renderChildren(location Rectangle, context *gg.Context) {
	left := location.left
	top := location.top

	w := 0.0
	h := 0.0
	wGap := 0.0
	hGap := 0.0

	if container.justifyContent == END {
		if container.layout == ROW {
			totalWidth := container.getTotalChildWidth(location)
			w = location.width - totalWidth
		} else if container.layout == COLUMN {
			totalHeight := container.getTotalChildHeight(location)
			h = location.height - totalHeight
		}
	} else if container.justifyContent == SPACED {
		if container.layout == ROW {
			totalWidth := container.getTotalChildWidth(location)
			wGap = (location.width - totalWidth) / float64(len(container.children)+1)
			w = wGap
		} else if container.layout == COLUMN {
			totalHeight := container.getTotalChildHeight(location)
			hGap = (location.height - totalHeight) / float64(len(container.children)+1)
			h = hGap
		}
	}

	for _, child := range container.children {
		width := getChildWidth(location, child)
		height := getChildHeight(location, child)

		if container.alignItems == ALIGN_END {
			if container.layout == ROW {
				h = location.height - height
			} else if container.layout == COLUMN {
				w = location.width - width
			}
		} else if container.alignItems == ALIGN_CENTER {
			if container.layout == ROW {
				h = location.height/2 - height/2
			} else if container.layout == COLUMN {
				w = location.width/2 - width/2
			}
		}

		childLocation := Rectangle{left + container.padding + w, top + container.padding + h, width, height}
		child.Render(childLocation, context)
		if container.layout == ROW {
			w += width + wGap
		} else if container.layout == COLUMN {
			h += height + hGap
		}
	}
}

func getChildWidth(location Rectangle, child Displayable) float64 {
	if child.GetWidth() <= 1 {
		return location.width * child.GetWidth()
	} else {
		return child.GetWidth()
	}
}

func getChildHeight(location Rectangle, child Displayable) float64 {
	if child.GetHeight() <= 1 {
		return location.height * child.GetHeight()
	} else {
		return child.GetHeight()
	}
}

func (container *Container) getTotalChildWidth(location Rectangle) float64 {
	width := 0.0
	for _, child := range container.children {
		width += getChildWidth(location, child)
	}
	return width
}

func (container *Container) getTotalChildHeight(location Rectangle) float64 {
	height := 0.0
	for _, child := range container.children {
		height += getChildHeight(location, child)
	}
	return height
}
