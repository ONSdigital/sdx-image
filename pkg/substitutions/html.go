package substitutions

import "strings"

var htmlLookupMap = map[string]string{
	"&#39;": "'",
	"<em>":  "",
	"</em>": "",
}

func html(text string) string {
	result := text
	for k, v := range htmlLookupMap {
		result = strings.Replace(result, k, v, 1)
	}
	return result
}
