package page

import "sdxImage/pkg/drawing"

const titleSize = 30
const titlePadding = 20
const divHeight = 6

type Section struct {
	title     string
	canvas    *drawing.Canvas
	container *drawing.Container
}

var grey = drawing.CreateColor(200, 200, 200)

func newSection(title string, canvas *drawing.Canvas, parent *drawing.Container) *Section {

	container := canvas.AddContainer(1, 0, parent).SetLayout(
		drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)

	canvas.AddDiv(1, divHeight, container).SetBackgroundColor(grey)

	titleArea := canvas.AddContainer(1, 0, container).SetPadding(0, titlePadding, 0, titlePadding)
	canvas.AddBoldText(title, titleSize, titleArea).SetTextAlign(drawing.TextLeft)

	return &Section{
		title:     title,
		canvas:    canvas,
		container: container,
	}
}

func (section *Section) addAnswer(qCode, question, answer string) {
	createAnswer(qCode, question, answer, section.canvas, section.container)
}
