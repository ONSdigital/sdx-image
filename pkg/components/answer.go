package components

import "sdxImage/pkg/drawing"

const qCodeSize = 24
const questionSize = 24
const answerSize = 30

func CreateAnswer(qCode, question, answer string, canvas *drawing.Canvas, container *drawing.Container) {

	outerContainer := canvas.AddContainer(1, 0, container)
	outerContainer.SetLayout(drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)

	questionContainer := canvas.AddContainer(1, 0, outerContainer)
	questionContainer.SetLayout(drawing.LayoutRow, drawing.JustifyStart, drawing.AlignStart)

	qCodeBox := canvas.AddContainer(0.1, 80, questionContainer)
	canvas.AddBoldText(qCode, qCodeSize, qCodeBox)

	questionBox := canvas.AddContainer(0.9, 0, questionContainer)
	canvas.AddText(question, questionSize, questionBox)

	answerContainer := canvas.AddContainer(1, 0, outerContainer)
	answerContainer.SetLayout(drawing.LayoutRow, drawing.JustifyStart, drawing.AlignStart)
	answerText := canvas.AddText(answer, answerSize, answerContainer)
	answerText.SetColor(drawing.RED)
}
