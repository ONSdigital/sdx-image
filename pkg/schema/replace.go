package schema

import (
	"fmt"
	"sdxImage/pkg/model"
	"strings"
)

const unknownIdentifier = "~"

type lookup func(param string) string

func InterpolateParams(m *model.Survey) *model.Survey {
	for _, section := range m.Sections {
		for _, question := range section.Questions {
			fmt.Println(question.Title)
			q := replaceParam(question.Title)
			question.Title = replaceHtml(q)
		}
	}
	return m
}

var paramLookupMap = map[string]string{
	"ref_p_start_date": "start date",
	"ref_p_end_date":   "end date",
	"ru_name":          "the business",
}

func paramLookup(p string) string {
	x, found := paramLookupMap[p[1:len(p)-1]]
	if !found {
		x = unknownIdentifier
	}
	return x
}

var htmlLookupMap = map[string]string{
	"&#39;": "'",
}

func htmlLookup(p string) string {
	x, found := htmlLookupMap[p]
	if !found {
		x = ""
	}
	return x
}

func replaceParam(text string) string {
	result := replace(text, "{", "}", paramLookup)
	return replaceUnknown(result)
}

func replaceUnknown(text string) string {
	result := replace(text,
		unknownIdentifier,
		unknownIdentifier,
		func(p string) string { return "" })
	return strings.Replace(result, unknownIdentifier, "", 1)
}

func replaceHtml(text string) string {
	return replace(text, "&", ";", htmlLookup)
}

func replace(text, startChar, endChar string, f lookup) string {
	p, exists := getParameter(text, startChar, endChar)
	if !exists {
		return text
	}
	x := f(p)
	result := strings.Replace(text, p, x, 1)
	return replace(result, startChar, endChar, f)
}

func getParameter(text, startChar, endChar string) (string, bool) {
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
