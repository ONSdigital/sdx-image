package drawing

import "github.com/fogleman/gg"

type Layout int

const (
	LayoutRow Layout = iota
	LayoutColumn
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

func newContainer(width, height float64) *Container {
	return &Container{
		Div:            newDiv(width, height),
		layout:         LayoutColumn,
		justifyContent: JustifyStart,
		alignItems:     AlignStart,
		padding:        0.0,
	}
}

func (container *Container) SetLayout(layout Layout, justifyContent Justification, alignItems Alignment) {
	container.layout = layout
	container.justifyContent = justifyContent
	container.alignItems = alignItems
}

func (container *Container) SetPadding(padding float64) {
	container.padding = padding
}

func (container *Container) GetHeight(parent Dimension) float64 {
	height := container.Div.GetHeight(parent)
	if height == 0 {
		internal := container.Div.getInternalDim(parent)
		internal.width -= 2 * container.padding
		internal.height -= 2 * container.padding
		if container.layout == LayoutColumn {
			height = container.getTotalChildHeight(internal)
		} else {
			height = container.getLargestChildHeight(internal)
		}
		height += 2 * container.borderWeight
		height += 2 * container.padding
	}
	return height
}

func (container *Container) Render(area Rectangle, context *gg.Context) {
	internalArea := container.Div.getInternalArea(area)
	internalArea.left += container.padding
	internalArea.top += container.padding
	internalArea.width -= 2 * container.padding
	internalArea.height -= 2 * container.padding
	container.Div.Render(area, context)
	container.renderChildren(internalArea, context)
}

func (container *Container) renderChildren(area Rectangle, context *gg.Context) {
	left := area.left
	top := area.top

	w := 0.0
	h := 0.0
	wGap := 0.0
	hGap := 0.0

	if container.justifyContent == JustifyEnd {
		if container.layout == LayoutRow {
			totalWidth := container.getTotalChildWidth(area.Dimension)
			w = area.width - totalWidth
		} else if container.layout == LayoutColumn {
			totalHeight := container.getTotalChildHeight(area.Dimension)
			h = area.height - totalHeight
		}
	} else if container.justifyContent == JustifySpaced {
		if container.layout == LayoutRow {
			totalWidth := container.getTotalChildWidth(area.Dimension)
			wGap = (area.width - totalWidth) / float64(len(container.children)+1)
			w = wGap
		} else if container.layout == LayoutColumn {
			totalHeight := container.getTotalChildHeight(area.Dimension)
			hGap = (area.height - totalHeight) / float64(len(container.children)+1)
			h = hGap
		}
	}

	for _, child := range container.children {
		width := child.GetWidth(area.Dimension)
		height := child.GetHeight(area.Dimension)

		if container.alignItems == AlignEnd {
			if container.layout == LayoutRow {
				h = area.height - height
			} else if container.layout == LayoutColumn {
				w = area.width - width
			}
		} else if container.alignItems == AlignCenter {
			if container.layout == LayoutRow {
				h = area.height/2 - height/2
			} else if container.layout == LayoutColumn {
				w = area.width/2 - width/2
			}
		}

		childArea := newRectangle(left+w, top+h, width, height)
		child.Render(childArea, context)
		if container.layout == LayoutRow {
			w += width + wGap
		} else if container.layout == LayoutColumn {
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
