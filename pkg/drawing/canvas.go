package drawing

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
)

var WHITE = color.White
var BLACK = color.Black
var RED = color.RGBA{R: 255, G: 1, B: 1}
var CYAN = color.RGBA{R: 161, G: 211, B: 225}

type Canvas struct {
	Dimension
	body    *Container
	context *gg.Context
}

func NewCanvas(width, height int) *Canvas {
	context := gg.NewContext(width, height)
	w := float64(width)
	h := float64(height)
	body := newContainer(w, h, context)
	body.Layout = Row
	return &Canvas{Dimension: Dimension{w, h}, body: body, context: context}
}

func (canvas *Canvas) GetBody() *Container {
	return canvas.body
}

func (canvas *Canvas) AddContainer(width, height float64, container *Container) *Container {
	c := newContainer(width, height, canvas.context)
	container.children = append(container.children, c)
	return c
}

func (canvas *Canvas) AddDiv(width, height float64, container *Container) *Div {
	div := newDiv(width, height, canvas.context)
	container.children = append(container.children, div)
	return div
}

func (canvas *Canvas) AddText(value string, size int, container *Container) *Text {
	text := newText(value, size, 1.0, canvas.context)
	container.children = append(container.children, text)
	return text
}

func (canvas *Canvas) Draw() image.Image {
	area := newRectangle(0, 0, canvas.width, canvas.height)
	canvas.context.SetColor(WHITE)
	canvas.context.Clear()
	canvas.body.Render(area)
	return canvas.context.Image()
}
