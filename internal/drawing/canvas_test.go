package drawing

import (
	"sdxImage/internal/test"
	"testing"
)

func TestCanvas(t *testing.T) {
	test.SetCwdToRoot()

	var filename = "drawing-test"
	var width = 1241
	var red = CreateColor(255, 0, 0)
	var blue = CreateColor(0, 0, 255)
	var green = CreateColor(0, 255, 0)
	var yellow = CreateColor(255, 255, 0)

	canvas := NewCanvas(width)
	outer := canvas.AddTopLevelContainer(
		MatchParent(),
		FitContent()).SetPaddingAll(10)

	top := canvas.AddContainer(MatchParent(), FitContent(), outer)
	top.SetBorder(blue, 4)
	canvas.AddText("Title", 24, top).SetTextAlign(TextCenter)

	middle := canvas.AddContainer(MatchParent(), FitContent(),
		outer).SetLayout(LayoutRow, JustifySpaced, AlignCenter)
	canvas.AddDiv(ProportionOfParent(0.25), PX(50),
		middle).SetBackgroundColor(red)
	canvas.AddDiv(ProportionOfParent(0.25), PX(150),
		middle).SetBackgroundColor(blue)
	x := canvas.AddContainer(ProportionOfParent(0.25), PX(80),
		middle).SetLayout(LayoutRow, JustifySpaced, AlignCenter)
	x.SetBackgroundColor(green)
	canvas.AddDiv(ProportionOfParent(0.2), ProportionOfParent(0.25),
		x).SetBackgroundColor(red)
	canvas.AddDiv(ProportionOfParent(0.2), ProportionOfParent(0.25),
		x).SetBackgroundColor(red)
	canvas.AddDiv(ProportionOfParent(0.2), ProportionOfParent(0.25),
		x).SetBackgroundColor(red)

	bottom := canvas.AddContainer(MatchParent(), PX(250),
		outer).SetLayout(LayoutColumn, JustifyEnd, AlignStart)
	canvas.AddDiv(ProportionOfParent(0.4), PX(100),
		bottom).SetBackgroundColor(yellow).SetBorder(green, 20)
	canvas.AddDiv(ProportionOfParent(0.6), PX(120),
		bottom).SetBackgroundColor(blue)

	result := canvas.Draw(10)
	err := test.SaveJPG("temp/"+filename+".jpg", result, 100)
	if err != nil {
		t.Errorf("failed to create image for %s with error: %q", filename, err.Error())
	}
}
