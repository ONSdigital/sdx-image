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

	container := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), parent).SetLayout(
		drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)

	canvas.AddDiv(drawing.MatchParent(), drawing.PX(divHeight), container).SetBackgroundColor(grey)

	titleArea := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), container).SetPadding(0, titlePadding, 0, titlePadding)
	canvas.AddBoldText(title, titleSize, titleArea).SetTextAlign(drawing.TextLeft)

	return &Section{
		title:     title,
		canvas:    canvas,
		container: container,
	}
}

func (section *Section) addInstance(id int) *Instance {
	return newInstance(id, section.canvas, section.container)
}
