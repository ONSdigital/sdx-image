package schema

import (
	"fmt"
	"sdxImage/pkg/model"
	"strings"
)

func InterpolateParams(m *model.Survey) *model.Survey {
	for _, question := range m.Questions {
		fmt.Println(question)
	}
	return m
}

func replace(text string, lookup map[string]string) string {
	p, exists := getParameter(text)
	if !exists {
		return text
	}

	x, found := lookup[p[1:len(p)-1]]
	if !found {
		x = ""
	}
	result := strings.Replace(text, p, x, 1)
	return replace(result, lookup)
}

func getParameter(text string) (string, bool) {
	start := -1
	end := -1
	for pos, char := range text {
		if string(char) == "{" {
			start = pos
		} else if string(char) == "}" {
			end = pos
			break
		}
	}
	if start > -1 && end > -1 {
		return text[start : end+1], true
	}
	return text, false
}
