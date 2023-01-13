package main

import (
	"image"
	"image/jpeg"
	"os"
	"sdxImage/pkg/components"
)

func main() {
	header := components.Header{
		SurveyName:  "Annual Business Survey",
		FormType:    "1802",
		RuRef:       "49800108249D",
		SubmittedAt: "11 January 2023 14:49:33"}

	page := components.CreatePage(header)

	reportingSection := page.AddSection("Reporting")
	reportingSection.AddAnswer("11", "To", "01/01/2021")
	reportingSection.AddAnswer("12", "From", "01/01/2022")

	incomeSection := page.AddSection("Income")
	incomeSection.AddAnswer("123", "What was your total turnover?", "56000")
	incomeSection.AddAnswer("124", "What did you spend on goods?", "34000")

	err := saveJPG("images/test.jpg", page.Draw(), 100)
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
