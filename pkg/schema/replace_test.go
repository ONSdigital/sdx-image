package schema

import "testing"

func TestGetParameter(t *testing.T) {
	text := "hello {world}"
	result, _ := getParameter(text, "{", "}")
	if result != "{world}" {
		t.Errorf("failed to extract parameter {world} got %v instead", result)
	}
}

func TestGetHtmlParameter(t *testing.T) {
	text := "the business&#39;s turnover"
	result, _ := getParameter(text, "&", ";")
	if result != "&#39;" {
		t.Errorf("failed to extract parameter &#39; got %v instead", result)
	}
}

func TestReplaceParam(t *testing.T) {
	text := "hello {ref_p_start_date}"
	result := replaceParam(text)
	if result != "hello start date" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}

func TestReplaceTwoParams(t *testing.T) {
	text := "Can you report from {ref_p_start_date} to {ref_p_end_date}?"
	result := replaceParam(text)
	if result != "Can you report from start date to end date?" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}

func TestReplaceHtml(t *testing.T) {
	text := "the business&#39;s turnover"
	result := replaceHtml(text)
	if result != "the business's turnover" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}
