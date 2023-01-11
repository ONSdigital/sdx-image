package main

import (
	"image"
	"image/jpeg"
	"os"
	"sdxImage/pkg/components"
	"sdxImage/pkg/drawing"
)

func main() {
	width := 1241
	canvas := drawing.NewCanvas(width)

	outer := canvas.AddTopLevelContainer(float64(width), 1500)
	outer.SetLayout(drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)
	outer.SetPadding(140)

	components.CreateHeading(
		"Annual Business Survey",
		"1802",
		"49800108249D",
		"11 January 2023 14:49:33",
		canvas,
		outer)

	canvas.AddDiv(1, 140, outer)

	components.CreateAnswer("11", "What was your total turnover?", "56000", canvas, outer)

	err := saveJPG("images/test.jpg", canvas.Draw(), 100)
	if err != nil {
		return
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
