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
	value     string
	size      int
	color     color.Color
	textAlign TextAlign
	context   *gg.Context
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
		Base:      newBase(1, 0),
		value:     value,
		size:      size,
		color:     color.Black,
		textAlign: TextLeft,
		context:   context,
	}
}

func (text *Text) SetColor(color color.Color) {
	text.color = color
}

func (text *Text) SetTextAlign(textAlign TextAlign) {
	text.textAlign = textAlign
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

func (text *Text) Render(area Rectangle, context *gg.Context) {
	context.SetColor(text.color)
	context.SetFontFace(FontMap[text.size])
	lines := text.context.WordWrap(text.value, area.width)

	h := float64(text.size)
	if len(lines) == 1 {
		vw, _ := text.context.MeasureString(text.value)
		if text.textAlign == TextCenter {
			context.DrawString(text.value, area.left+(area.width-vw)/2, area.top+h)
		} else if text.textAlign == TextRight {
			context.DrawString(text.value, area.left+(area.width-vw), area.top+h)
		} else {
			context.DrawString(text.value, area.left, area.top+h)
		}
	} else {
		for _, line := range lines {
			context.DrawString(line, area.left, area.top+h)
			h += float64(text.size) * LineSpacing
		}
	}
}
