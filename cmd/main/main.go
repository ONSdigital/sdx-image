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
	minHeight := 1754.0
	canvas := drawing.NewCanvas(width)

	outer := canvas.AddTopLevelContainer(float64(width), 0)
	outer.SetLayout(
		drawing.LayoutColumn,
		drawing.JustifyStart,
		drawing.AlignStart).SetPaddingAll(140)

	components.CreateHeading(
		"Annual Business Survey",
		"1802",
		"49800108249D",
		"11 January 2023 14:49:33",
		canvas,
		outer)

	reportingSection := components.NewSection("Reporting", canvas, outer)
	reportingSection.AddAnswer("11", "To", "01/01/2021")
	reportingSection.AddAnswer("12", "From", "01/01/2022")

	incomeSection := components.NewSection("Income", canvas, outer)
	incomeSection.AddAnswer("123", "What was your total turnover?", "56000")
	incomeSection.AddAnswer("124", "What did you spend on goods?", "34000")

	err := saveJPG("images/test.jpg", canvas.Draw(minHeight), 100)
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
