package page

import (
	"sdxImage/internal/drawing"
	"sdxImage/internal/survey"
)

const luPadding = 30

func createUnit(
	lu survey.Unit,
	canvas *drawing.Canvas,
	parent *drawing.Container) {

	outerContainer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), parent)
	outerContainer.SetLayout(drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)
	outerContainer.SetBorder(drawing.BLACK, 2)
	outerContainer.SetPaddingAll(luPadding)

	refLabel := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	canvas.AddText("Reference:", questionSize, refLabel)

	ref := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	canvas.AddBoldText(lu.GetIdentifier(), answerSize, ref)
	ref.SetPadding(0, 0, 0, luPadding)

	primaryDescription := lu.GetPrimaryDesc()

	if primaryDescription.Value != "" {
		nameLabel := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
		canvas.AddText(primaryDescription.Key+":", questionSize, nameLabel)

		nameAnswer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
		canvas.AddBoldText(primaryDescription.Value, answerSize, nameAnswer)
		nameAnswer.SetPadding(0, 0, 0, luPadding)
	}

	secondaryDescription := lu.GetSecondaryDesc()

	if secondaryDescription.Value != "" {
		addressLabel := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
		canvas.AddText(secondaryDescription.Key+":", questionSize, addressLabel)

		addressAnswer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
		canvas.AddBoldText(secondaryDescription.Value, answerSize, addressAnswer)

		if len(lu.GetAnswers()) > 0 {
			addressAnswer.SetPadding(0, 0, 0, luPadding)
		}
	}

	for _, answer := range lu.GetAnswers() {
		text := answer.GetText()
		if text == "" {
			//default value if nothing provided
			text = "Changes"
		}
		createAnswer(answer.GetCode(), answer.GetText(), answer.GetValue(), canvas, outerContainer)
	}
}
