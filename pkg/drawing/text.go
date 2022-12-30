package drawing

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image/color"
)

const LineSpacing = 1.5

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

func (text *Text) GetHeight(parent Dimension) float64 {
	text.context.SetFontFace(FontMap[text.size])
	width := text.GetWidth(parent)
	lines := text.context.WordWrap(text.value, width)
	size := float64(text.size)
	n := float64(len(lines))
	height := n * size * LineSpacing
	return height
}

func (text *Text) Render(area Rectangle) {
	text.context.SetColor(text.Color)
	text.context.SetFontFace(FontMap[text.size])
	lines := text.context.WordWrap(text.value, area.width)
	h := float64(text.size)
	for _, line := range lines {
		text.context.DrawString(line, area.left, area.top+h)
		h += float64(text.size) * LineSpacing
	}
}
