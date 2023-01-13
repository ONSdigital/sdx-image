package components

import (
	"sdxImage/pkg/drawing"
)

const padding = 6.0

var cyan = drawing.CreateColor(161, 211, 225)

type Header struct {
	SurveyName, FormType, RuRef, SubmittedAt string
}

func createHeading(header Header, canvas *drawing.Canvas, parent *drawing.Container) {

	headerArea := canvas.AddContainer(1, 0, parent)
	headerArea.SetPadding(0, 0, 0, 60)

	headerBox := canvas.AddContainer(1, 0, headerArea)
	headerBox.SetBorder(drawing.BLACK, 2)

	titleBox := canvas.AddContainer(1, 0, headerBox)
	titleBox.SetPaddingAll(padding).SetBackgroundColor(cyan).SetBorder(drawing.BLACK, 1)

	titleText := canvas.AddBoldText(header.SurveyName, 36, titleBox)
	titleText.SetTextAlign(drawing.TextCenter)

	detailsBox := canvas.AddContainer(1, 0, headerBox)
	detailsBox.SetLayout(drawing.LayoutRow, drawing.JustifyStart, drawing.AlignStart)
	detailsLeft := canvas.AddContainer(0.5, 0, detailsBox)
	detailsRight := canvas.AddContainer(0.5, 0, detailsBox)

	for _, key := range []string{"Form Type", "Respondent", "Submitted At"} {
		c := canvas.AddContainer(1, 0, detailsLeft)
		c.SetBorder(drawing.BLACK, 1)
		c.SetPaddingAll(padding)
		t := canvas.AddText(key, 24, c)
		t.SetTextAlign(drawing.TextLeft)
	}

	for _, value := range []string{header.FormType, header.RuRef, header.SubmittedAt} {
		c := canvas.AddContainer(1, 0, detailsRight)
		c.SetBorder(drawing.BLACK, 1)
		c.SetPaddingAll(padding)
		t := canvas.AddText(value, 24, c)
		t.SetTextAlign(drawing.TextLeft)
	}
}
