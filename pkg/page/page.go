package page

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

func createPage(header Header) *Page {
	canvas := drawing.NewCanvas(width)

	outer := canvas.AddTopLevelContainer(drawing.PX(float64(width)), drawing.FitContent())
	outer.SetLayout(
		drawing.LayoutColumn,
		drawing.JustifyStart,
		drawing.AlignStart).SetPaddingAll(140)

	createHeading(header, canvas, outer)

	return &Page{canvas, outer}
}

func (page *Page) addSection(title string) *Section {
	return newSection(title, page.canvas, page.outer)
}

func (page *Page) draw() image.Image {
	return page.canvas.Draw(float64(minHeight))
}
