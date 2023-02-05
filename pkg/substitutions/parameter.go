package substitutions

import (
	"strings"
)

type Lookup interface {
	get(str string) string
}

const pStart = "{"
const pEnd = "}"

func replaceParameters(text string, lookup Lookup) string {
	p, exists := findParameter(text, pStart, pEnd)
	if !exists {
		return text
	}
	x := lookup.get(p[1 : len(p)-1])
	result := strings.Replace(text, p, x, 1)
	return replaceParameters(result, lookup)
}

func findParameter(text, startChar, endChar string) (string, bool) {
	start := -1
	end := -1
	for pos, char := range text {
		if string(char) == startChar && start == -1 {
			start = pos
		} else if string(char) == endChar {
			end = pos
			break
		}
	}
	if start > -1 && end > -1 {
		return text[start : end+1], true
	}
	return text, false
}
