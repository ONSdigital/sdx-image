package schema

import (
	"sdxImage/pkg/model"
	"strings"
)

type lookup func(param string) string

func InterpolateParams(m *model.Survey) *model.Survey {
	for _, question := range m.Questions {
		q := replaceParam(question.Title)
		question.Title = replaceHtml(q)
	}
	return m
}

var paramLookupMap = map[string]string{
	"ref_p_start_date": "start date",
	"ref_p_end_date":   "end date",
}

func paramLookup(p string) string {
	x, found := paramLookupMap[p[1:len(p)-1]]
	if !found {
		x = ""
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
	return replace(text, "{", "}", paramLookup)
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
		if string(char) == startChar {
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
