package drawing

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

const LineSpacing = 1.5

type TextAlign int

const (
	TextLeft TextAlign = iota
	TextRight
	TextCenter
)

var FontMap = make(map[int]font.Face)
var BoldFontMap = make(map[int]font.Face)

type Text struct {
	*Base
	value     string
	size      int
	bold      bool
	color     Colour
	textAlign TextAlign
	context   *gg.Context
}

func newText(value string, size int, bold bool, context *gg.Context) *Text {
	if bold {
		_, exists := BoldFontMap[size]
		if !exists {
			fontFace, err := gg.LoadFontFace("fonts/luxisb.ttf", float64(size))
			if err != nil {
				panic(err)
			}
			BoldFontMap[size] = fontFace
		}
	} else {
		_, exists := FontMap[size]
		if !exists {
			fontFace, err := gg.LoadFontFace("fonts/luxisr.ttf", float64(size))
			if err != nil {
				panic(err)
			}
			FontMap[size] = fontFace
		}
	}
	return &Text{
		Base:      newBase(1, 0),
		value:     value,
		size:      size,
		bold:      bold,
		color:     Color{0, 0, 0},
		textAlign: TextLeft,
		context:   context,
	}
}

func (text *Text) SetColor(color Colour) *Text {
	text.color = color
	return text
}

func (text *Text) SetTextAlign(textAlign TextAlign) *Text {
	text.textAlign = textAlign
	return text
}

func (text *Text) GetHeight(parent Dimension) float64 {
	if text.bold {
		text.context.SetFontFace(BoldFontMap[text.size])
	} else {
		text.context.SetFontFace(FontMap[text.size])
	}
	width := text.GetWidth(parent)
	lines := text.context.WordWrap(text.value, width)
	size := float64(text.size)
	n := float64(len(lines))
	height := n * size * LineSpacing
	return height
}

func (text *Text) Render(area Rectangle, context *gg.Context) {

	if text.bold {
		context.SetFontFace(BoldFontMap[text.size])
	} else {
		context.SetFontFace(FontMap[text.size])
	}

	setColour(text.color, context)
	lines := context.WordWrap(text.value, area.width)

	h := float64(text.size)
	if len(lines) == 1 {
		vw, _ := context.MeasureString(text.value)
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
