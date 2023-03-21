package page

import "sdxImage/pkg/drawing"

const titleSize = 30
const titlePadding = 20
const divHeight = 6
const sectionGap = 40

type Section struct {
	title         string
	canvas        *drawing.Canvas
	container     *drawing.Container
	instanceCount int
}

var grey = drawing.CreateColor(200, 200, 200)

func newSection(title string, canvas *drawing.Canvas, parent *drawing.Container) *Section {

	container := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), parent).SetLayout(
		drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)

	canvas.AddDiv(drawing.MatchParent(), drawing.PX(divHeight), container).SetBackgroundColor(grey)

	titleArea := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), container).SetPadding(0, titlePadding, 0, titlePadding)
	canvas.AddBoldText(title, titleSize, titleArea).SetTextAlign(drawing.TextLeft)

	return &Section{
		title:         title,
		canvas:        canvas,
		container:     container,
		instanceCount: 0,
	}
}

func (section *Section) addInstance(id int) *Instance {
	section.instanceCount++
	return newInstance(id, section.canvas, section.container)
}

func (section *Section) complete() {
	// add some whitespace at the end of a section if contains multiple instances
	if section.instanceCount > 1 {
		section.canvas.AddDiv(drawing.MatchParent(), drawing.PX(sectionGap), section.container)
	}
}
