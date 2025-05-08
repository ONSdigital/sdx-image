package substitutions

import "strings"

var htmlLookupMap = map[string]string{
	"&#39;":     "'",
	"&#x2019;":  "'",
	"&amp;":     "&",
	"<em>":      "",
	"</em>":     "",
	"<strong>":  "",
	"</strong>": "",
}

var spaceLookupMap = map[string]string{
	" , ": " ",
	"  ":  " ",
	" ?":  "?",
}

// html replaces common "html" codes within a string
// and tries to fix any unusual gaps.
func html(text string) string {
	result := text
	for k, v := range htmlLookupMap {
		result = strings.Replace(result, k, v, -1)
	}
	for k, v := range spaceLookupMap {
		result = strings.Replace(result, k, v, -1)
	}
	return result
}
