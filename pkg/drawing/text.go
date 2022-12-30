package drawing

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image/color"
)

const LineSpacing = 1.5

type TextAlign int

const (
	TextLeft TextAlign = iota
	TextRight
	TextCenter
)

var FontMap = make(map[int]font.Face)

type Text struct {
	*Base
	value string
	size  int
	Color color.Color
	TextAlign
}

func newText(value string, size int, context *gg.Context) *Text {
	_, exists := FontMap[size]
	if !exists {
		fontFace, err := gg.LoadFontFace("fonts/luxisr.ttf", float64(size))
		if err != nil {
			panic(err)
		}
		FontMap[size] = fontFace
	}
	return &Text{
		Base:      newBase(1, 0, context),
		value:     value,
		size:      size,
		Color:     color.Black,
		TextAlign: TextLeft,
	}
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
	if len(lines) == 1 {
		vw, _ := text.context.MeasureString(text.value)
		if text.TextAlign == TextCenter {
			text.context.DrawString(text.value, area.left+(area.width-vw)/2, area.top+h)
		} else if text.TextAlign == TextRight {
			text.context.DrawString(text.value, area.left+(area.width-vw), area.top+h)
		} else {
			text.context.DrawString(text.value, area.left, area.top+h)
		}
	} else {
		for _, line := range lines {
			text.context.DrawString(line, area.left, area.top+h)
			h += float64(text.size) * LineSpacing
		}
	}
}
