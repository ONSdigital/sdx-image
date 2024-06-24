package substitutions

import "testing"

func TestHtml(t *testing.T) {
	text := "the business&#39;s turnover"
	result := html(text)
	if result != "the business's turnover" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}

func TestHtmlAmpersand(t *testing.T) {
	text := "In-house R&amp;D"
	result := html(text)
	if result != "In-house R&D" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}

func TestHtmlStrong(t *testing.T) {
	text := "<strong>Hello World"
	result := html(text)
	if result != "Hello World" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}

func TestHtmlStrongEndtag(t *testing.T) {
	text := "Hello World</strong>"
	result := html(text)
	if result != "Hello World" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}
