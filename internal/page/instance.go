package page

import (
	"sdxImage/internal/drawing"
	"sdxImage/internal/survey"
)

const indent = 10
const topPadding = 20
const spaceBelow = 10

type Instance struct {
	canvas    *drawing.Canvas
	container *drawing.Container
}

func newInstance(id int, canvas *drawing.Canvas, parent *drawing.Container) *Instance {

	container := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), parent).SetLayout(
		drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)

	if id > 0 {
		container.SetPadding(indent, topPadding, indent, 0).SetBorder(red, 2.0)
		canvas.AddDiv(drawing.MatchParent(), drawing.PX(spaceBelow), parent)
	}

	return &Instance{
		canvas:    canvas,
		container: container,
	}
}

func (instance *Instance) addAnswer(qCode, question, answer string) {
	createAnswer(qCode, question, answer, instance.canvas, instance.container)
}

func (instance *Instance) addLocalUnit(lu *survey.Unit) {
	createUnit(lu, instance.canvas, instance.container)
}
