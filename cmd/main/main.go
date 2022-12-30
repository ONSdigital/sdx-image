package main

import (
	"github.com/fogleman/gg"
	"sdxImage/pkg/drawing"
)

func main() {
	width := 1241
	height := 1754
	canvas := drawing.NewCanvas(width, height)

	outer := canvas.AddContainer(700, 1000, canvas.GetBody())
	outer.Padding = 20
	outer.Layout = drawing.Column
	outer.BackgroundColor = drawing.RED
	outer.BorderColor = drawing.BLACK
	outer.JustifyContent = drawing.JustifyStart

	top := canvas.AddContainer(500, 0, outer)
	header := canvas.AddText("Hello", 48, top)
	header.TextAlign = drawing.TextCenter

	middle := canvas.AddContainer(500, 500, outer)
	middle.Layout = drawing.Row
	middle.BackgroundColor = drawing.CYAN
	middle.BorderColor = drawing.BLACK
	middle.BorderWeight = 20
	middle.JustifyContent = drawing.JustifySpaced
	middle.AlignItems = drawing.AlignCenter

	createChildren(3, middle, canvas)

	bottom := canvas.AddContainer(500, 0, outer)
	bottom.Layout = drawing.Row
	bottom.BorderColor = drawing.BLACK
	bottom.BorderWeight = 2

	canvas.AddText("A very long line of text that needs wrapping", 48, bottom)

	err := gg.SaveJPG("images/test.jpg", canvas.Draw(), 100)
	if err != nil {
		return
	}
}

func createChildren(n int, container *drawing.Container, canvas *drawing.Canvas) {
	for i := 0; i < n; i++ {
		c := canvas.AddContainer(0.3, 0.3, container)
		c.BackgroundColor = drawing.WHITE
		c.BorderColor = drawing.BLACK
		c.BorderWeight = 1
	}
}
