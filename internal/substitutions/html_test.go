package substitutions

import "testing"

func TestHtml(t *testing.T) {
	text := "the business&#39;s turnover"
	result := html(text)
	if result != "the business's turnover" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}
