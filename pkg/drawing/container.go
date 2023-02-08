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

type Padding struct {
	left, right, top, bottom float64
}

type Container struct {
	*Div
	children       []Displayable
	layout         Layout
	justifyContent Justification
	alignItems     Alignment
	padding        Padding
}

func newContainer(width, height Length) *Container {
	return &Container{
		Div:            newDiv(width, height),
		layout:         LayoutColumn,
		justifyContent: JustifyStart,
		alignItems:     AlignStart,
		padding:        Padding{0, 0, 0, 0},
	}
}

func (container *Container) SetLayout(
	layout Layout,
	justifyContent Justification,
	alignItems Alignment) *Container {

	container.layout = layout
	container.justifyContent = justifyContent
	container.alignItems = alignItems
	return container
}

func (container *Container) SetPadding(left, top, right, bottom float64) *Container {
	container.padding.left = left
	container.padding.top = top
	container.padding.right = right
	container.padding.bottom = bottom
	return container
}

func (container *Container) SetPaddingAll(padding float64) *Container {
	return container.SetPadding(padding, padding, padding, padding)
}

func (container *Container) GetHeight(parent Dimension) float64 {
	height := container.Div.GetHeight(parent)
	if height == 0 {
		internal := container.Div.getInternalDim(parent)
		internal.width -= container.padding.left + container.padding.right
		internal.height -= container.padding.top + container.padding.bottom
		if container.layout == LayoutColumn {
			height = container.getTotalChildHeight(internal)
		} else {
			height = container.getLargestChildHeight(internal)
		}
		height += 2 * container.borderWeight
		height += container.padding.top + container.padding.bottom
	}
	return height
}

func (container *Container) Render(area Rectangle, context *gg.Context) {
	internalArea := container.Div.getInternalArea(area)
	internalArea.left += container.padding.left
	internalArea.top += container.padding.top
	internalArea.width -= container.padding.left + container.padding.right
	internalArea.height -= container.padding.top + container.padding.bottom
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
