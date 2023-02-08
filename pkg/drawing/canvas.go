package drawing

import (
	"github.com/fogleman/gg"
	"image"
)

type Canvas struct {
	width   float64
	body    *Container
	context *gg.Context
}

func NewCanvas(width int) *Canvas {
	w := float64(width)
	body := newContainer(PX(w), FitContent())
	body.layout = LayoutRow
	return &Canvas{width: w, body: body}
}

func (canvas *Canvas) AddTopLevelContainer(width, height Length) *Container {
	return canvas.AddContainer(width, height, canvas.body)
}

func (canvas *Canvas) AddContainer(width, height Length, container *Container) *Container {
	c := newContainer(width, height)
	container.children = append(container.children, c)
	return c
}

func (canvas *Canvas) AddDiv(width, height Length, container *Container) *Div {
	div := newDiv(width, height)
	container.children = append(container.children, div)
	return div
}

func (canvas *Canvas) AddText(value string, size int, container *Container) *Text {
	tempContext := gg.NewContext(int(canvas.width), 1000)
	text := newText(value, size, false, tempContext)
	container.children = append(container.children, text)
	return text
}

func (canvas *Canvas) AddBoldText(value string, size int, container *Container) *Text {
	tempContext := gg.NewContext(int(canvas.width), 1000)
	text := newText(value, size, true, tempContext)
	container.children = append(container.children, text)
	return text
}

func (canvas *Canvas) Draw(minHeight float64) image.Image {
	width := canvas.width
	height := canvas.body.GetHeight(Dimension{width: canvas.width, height: 0})
	if height < minHeight {
		height = minHeight
	}

	context := gg.NewContext(int(width), int(height))
	area := newRectangle(0, 0, width, height)

	setColour(WHITE, context)
	context.Clear()
	canvas.body.Render(area, context)

	return context.Image()
}
