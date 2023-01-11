package components

import (
	"image/color"
	"sdxImage/pkg/drawing"
)

func CreateHeading(
	surveyName, formType, ruRef, submittedAt string,
	canvas *drawing.Canvas,
	container *drawing.Container) {

	padding := 8.0

	headerBox := canvas.AddContainer(1, 0, container)
	headerBox.SetBorder(drawing.BLACK, 4)

	titleBox := canvas.AddContainer(1, 0, headerBox)
	titleBox.SetBackgroundColor(drawing.CYAN)
	titleBox.SetBorder(color.Black, 1)
	titleBox.SetPadding(padding)

	titleText := canvas.AddText(surveyName, 34, titleBox)
	titleText.SetTextAlign(drawing.TextCenter)

	detailsBox := canvas.AddContainer(1, 0, headerBox)
	detailsBox.SetLayout(drawing.LayoutRow, drawing.JustifyStart, drawing.AlignStart)
	detailsLeft := canvas.AddContainer(0.5, 0, detailsBox)
	detailsRight := canvas.AddContainer(0.5, 0, detailsBox)

	for _, key := range []string{"Form Type", "Respondent", "Submitted At"} {
		c := canvas.AddContainer(1, 0, detailsLeft)
		c.SetBorder(color.Black, 1)
		c.SetPadding(padding)
		t := canvas.AddText(key, 24, c)
		t.SetTextAlign(drawing.TextLeft)
	}

	for _, value := range []string{formType, ruRef, submittedAt} {
		c := canvas.AddContainer(1, 0, detailsRight)
		c.SetBorder(color.Black, 1)
		c.SetPadding(padding)
		t := canvas.AddText(value, 24, c)
		t.SetTextAlign(drawing.TextLeft)
	}
}
