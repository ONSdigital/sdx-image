package drawing

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image/color"
)

var FontMap = make(map[int]font.Face)

type Text struct {
	*Base
	value string
	size  int
	Color color.Color
}

func newText(value string, size int, width float64, context *gg.Context) *Text {
	_, exists := FontMap[size]
	if !exists {
		fontFace, err := gg.LoadFontFace("fonts/luxisr.ttf", float64(size))
		if err != nil {
			panic(err)
		}
		FontMap[size] = fontFace
	}
	return &Text{value: value, size: size, Color: color.Black, Base: newBase(width, 0, context)}
}

func (text *Text) GetHeight(location Rectangle) float64 {
	text.context.SetFontFace(FontMap[text.size])
	lines := text.context.WordWrap(text.value, text.GetWidth(location))
	size := float64(text.size)
	n := float64(len(lines))
	height := n*size + (n-1)*size*0.5
	return height
}

func (text *Text) Render(location Rectangle) Rectangle {
	rect := text.Base.Render(location)
	text.context.SetColor(text.Color)
	text.context.SetFontFace(FontMap[text.size])
	lines := text.context.WordWrap(text.value, location.width)
	//text.context.DrawStringWrapped(text.value, rect.left, rect.top, 0, 0, rect.width, 1.5, gg.AlignLeft)
	h := float64(text.size)
	for _, line := range lines {
		text.context.DrawString(line, rect.left, rect.top+h)
		h += float64(text.size) * 1.5
	}
	return location
}
