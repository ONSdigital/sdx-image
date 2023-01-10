package main

import (
	"image"
	"image/jpeg"
	"os"
	"sdxImage/pkg/drawing"
)

func main() {
	width := 1241
	canvas := drawing.NewCanvas(width)

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
	bottom.BorderColor = drawing.WHITE
	bottom.Padding = 20

	canvas.AddText("A very long line of text that needs wrapping", 48, bottom)

	err := saveJPG("images/test.jpg", canvas.Draw(), 100)
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

func saveJPG(path string, im image.Image, quality int) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var opt jpeg.Options
	opt.Quality = quality

	return jpeg.Encode(file, im, &opt)
}
