package components

import (
	"image"
	"sdxImage/pkg/drawing"
)

var width = 1241
var minHeight = 1754

type Page struct {
	canvas *drawing.Canvas
	outer  *drawing.Container
}

func CreatePage(header Header) *Page {
	canvas := drawing.NewCanvas(width)

	outer := canvas.AddTopLevelContainer(float64(width), 0)
	outer.SetLayout(
		drawing.LayoutColumn,
		drawing.JustifyStart,
		drawing.AlignStart).SetPaddingAll(140)

	createHeading(header, canvas, outer)

	return &Page{canvas, outer}
}

func (page *Page) AddSection(title string) *Section {
	return newSection(title, page.canvas, page.outer)
}

func (page *Page) Draw() image.Image {
	return page.canvas.Draw(float64(minHeight))
}
