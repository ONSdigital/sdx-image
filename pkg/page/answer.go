package page

import "sdxImage/pkg/drawing"

const qCodeSize = 24
const questionSize = 24
const answerSize = 30
const answerPadding = 30

var red = drawing.CreateColor(255, 0, 0)

func createAnswer(
	qCode,
	question,
	answer string,
	canvas *drawing.Canvas,
	parent *drawing.Container) {

	outerContainer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), parent)
	outerContainer.SetLayout(drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)

	questionContainer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	questionContainer.SetLayout(drawing.LayoutRow, drawing.JustifyStart, drawing.AlignStart)

	qCodeBox := canvas.AddContainer(drawing.ProportionOfParent(0.07), drawing.FitContent(), questionContainer)
	canvas.AddBoldText(qCode+".", qCodeSize, qCodeBox)

	questionBox := canvas.AddContainer(drawing.ProportionOfParent(0.93), drawing.FitContent(), questionContainer)
	canvas.AddText(question, questionSize, questionBox)

	answerContainer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	answerContainer.SetLayout(drawing.LayoutRow, drawing.JustifyStart, drawing.AlignStart)
	answerContainer.SetPadding(0, answerPadding, 0, answerPadding)
	answerText := canvas.AddText(answer, answerSize, answerContainer)
	answerText.SetColor(red)
}
