package drawing

import (
	"github.com/fogleman/gg"
	"image/color"
)

var WHITE = color.White
var BLACK = color.Black
var RED = color.RGBA{R: 255, G: 1, B: 1}
var CYAN = color.RGBA{R: 161, G: 211, B: 225}

func Draw() {
	width := 1241
	height := 1754
	context := gg.NewContext(width, height)
	context.SetColor(WHITE)
	context.Clear()

	outer := newContainer(700, 1000, context)
	outer.padding = 20
	outer.layout = Column
	outer.backgroundColor = RED
	outer.borderColor = BLACK
	outer.justifyContent = JustifyStart

	middle := newContainer(500, 500, context)
	middle.layout = Row
	middle.backgroundColor = CYAN
	middle.borderColor = BLACK
	middle.borderWeight = 20
	middle.justifyContent = JustifySpaced
	middle.alignItems = AlignCenter

	createChildren(3, middle)

	bottom := newContainer(500, 500, context)
	bottom.layout = Row
	bottom.justifyContent = JustifySpaced
	text := newText("hello, a long line of text that needs wrapping", 48, 0.5, context)
	bottom.children = []Displayable{text}

	outer.children = []Displayable{middle, bottom}

	Display(outer, 10, 10)

	err := gg.SaveJPG("images/test.jpg", context.Image(), 100)
	if err != nil {
		return
	}
}

func createChildren(n int, container *Container) {
	var children = make([]Displayable, n)
	for i := 0; i < n; i++ {
		c := newContainer(0.3, 0.3, container.context)
		c.backgroundColor = WHITE
		c.borderColor = BLACK
		c.borderWeight = 1
		children[i] = c
	}
	container.children = children
}
