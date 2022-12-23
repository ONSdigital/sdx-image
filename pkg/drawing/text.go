package drawing

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image/color"
)

var FontMap map[int]font.Face = make(map[int]font.Face)

type Text struct {
	*Div
	value string
	size  int
	Color color.Color
}

func newText(value string, size int, width, height float64) *Text {
	_, exists := FontMap[size]
	if !exists {
		fontFace, err := gg.LoadFontFace("fonts/luxisr.ttf", float64(size))
		if err != nil {
			panic(err)
		}
		FontMap[size] = fontFace
	}
	return &Text{value: value, size: size, Color: color.Black, Div: newDiv(width, height)}
}

func (text *Text) GetWidth() float64 {
	return text.width
}

func (text *Text) GetHeight() float64 {
	return text.height
}

func (text *Text) Render(location Rectangle, context *gg.Context) Rectangle {
	rect := text.Div.Render(location, context)
	context.SetColor(text.Color)
	context.SetFontFace(FontMap[text.size])
	context.DrawStringWrapped(text.value, rect.left, rect.top, 0, 0, rect.width, 1.5, gg.AlignLeft)
	return location
}
