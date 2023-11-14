package page

import (
	"sdxImage/internal/drawing"
)

const padding = 6.0

var cyan = drawing.CreateColor(161, 211, 225)

type Header struct {
	SurveyName, FormType, RuRef, RuName, SubmittedAt string
}

func createHeading(header Header, canvas *drawing.Canvas, parent *drawing.Container) {

	headerArea := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), parent)
	headerArea.SetPadding(0, 0, 0, 60)

	headerBox := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), headerArea)
	headerBox.SetBorder(drawing.BLACK, 2)

	titleBox := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), headerBox)
	titleBox.SetPaddingAll(padding).SetBackgroundColor(cyan).SetBorder(drawing.BLACK, 1)

	titleText := canvas.AddBoldText(header.SurveyName, 36, titleBox)
	titleText.SetTextAlign(drawing.TextCenter)

	detailsBox := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), headerBox)
	detailsBox.SetLayout(drawing.LayoutRow, drawing.JustifyStart, drawing.AlignStart)
	detailsLeft := canvas.AddContainer(drawing.ProportionOfParent(0.5), drawing.FitContent(), detailsBox)
	detailsRight := canvas.AddContainer(drawing.ProportionOfParent(0.5), drawing.FitContent(), detailsBox)

	for _, key := range []string{"Form Type", "Respondent", "Company Name", "Submitted At"} {
		c := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), detailsLeft)
		c.SetBorder(drawing.BLACK, 1)
		c.SetPaddingAll(padding)
		t := canvas.AddText(key, 24, c)
		t.SetTextAlign(drawing.TextLeft)
	}

	companyName := header.RuName
	if len(companyName) > 30 {
		companyName = companyName[:30]
	}
	for _, value := range []string{header.FormType, header.RuRef, companyName, header.SubmittedAt} {
		c := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), detailsRight)
		c.SetBorder(drawing.BLACK, 1)
		c.SetPaddingAll(padding)
		t := canvas.AddText(value, 24, c)
		t.SetTextAlign(drawing.TextLeft)
	}
}
