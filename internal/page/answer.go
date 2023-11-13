package page

import "sdxImage/internal/drawing"

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

	//set the width proportion of qcode to question based on qcode size
	qCodeProportion := 0.07
	questionProportion := 0.93
	if len(qCode) > 3 {
		qCodeProportion = 0.1
		questionProportion = 0.9
	}

	qCodeBox := canvas.AddContainer(drawing.ProportionOfParent(qCodeProportion), drawing.FitContent(), questionContainer)
	canvas.AddBoldText(qCode+".", qCodeSize, qCodeBox)

	questionBox := canvas.AddContainer(drawing.ProportionOfParent(questionProportion), drawing.FitContent(), questionContainer)
	canvas.AddText(question, questionSize, questionBox)

	answerContainer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	answerContainer.SetLayout(drawing.LayoutRow, drawing.JustifyStart, drawing.AlignStart)
	answerContainer.SetPadding(0, answerPadding, 0, answerPadding)
	answerText := canvas.AddText(answer, answerSize, answerContainer)
	answerText.SetColor(red)
}
