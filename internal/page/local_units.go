package page

import (
	"sdxImage/internal/drawing"
)

const luPadding = 30

func createUnit(
	name string,
	address string,
	canvas *drawing.Canvas,
	parent *drawing.Container) {

	outerContainer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), parent)
	outerContainer.SetLayout(drawing.LayoutColumn, drawing.JustifyStart, drawing.AlignStart)
	outerContainer.SetBorder(drawing.BLACK, 2)
	outerContainer.SetPaddingAll(luPadding)

	nameLabel := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	canvas.AddText("Name:", questionSize, nameLabel)

	nameAnswer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	canvas.AddBoldText(name, answerSize, nameAnswer)
	nameAnswer.SetPadding(0, 0, 0, luPadding)

	addressLabel := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	canvas.AddText("Address:", questionSize, addressLabel)

	addressAnswer := canvas.AddContainer(drawing.MatchParent(), drawing.FitContent(), outerContainer)
	canvas.AddBoldText(address, answerSize, addressAnswer)
}
