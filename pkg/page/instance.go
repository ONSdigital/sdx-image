package page

import "sdxImage/pkg/drawing"

const indent = 10
const topPadding = 20

type Instance struct {
	canvas    *drawing.Canvas
	container *drawing.Container
}

func newInstance(id int, canvas *drawing.Canvas, parent *drawing.Container) *Instance {

	container := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), parent).SetLayout(
		drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)

	if id > 0 {
		container.SetPadding(indent, topPadding, 0, 0).SetBorder(red, 2.0)
	}

	return &Instance{
		canvas:    canvas,
		container: container,
	}
}

func (instance *Instance) addAnswer(qCode, question, answer string) {
	createAnswer(qCode, question, answer, instance.canvas, instance.container)
}