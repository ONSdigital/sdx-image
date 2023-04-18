package substitutions

import "strings"

var htmlLookupMap = map[string]string{
	"&#39;": "'",
	"&amp;": "&",
	"<em>":  "",
	"</em>": "",
}

// html replaces common "html" codes within a string
func html(text string) string {
	result := text
	for k, v := range htmlLookupMap {
		result = strings.Replace(result, k, v, -1)
	}
	return result
}
