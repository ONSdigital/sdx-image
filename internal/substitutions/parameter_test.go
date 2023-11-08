package substitutions

import (
	"testing"
)

type testMap map[string]string

func (tm testMap) get(str string) string {
	return tm[str]
}

var testLookup = testMap{
	"ref_p_start_date": "start date",
	"ref_p_end_date":   "end date",
	"ru_name":          "the business",
}

func TestGetParameter(t *testing.T) {
	text := "hello {world}"
	result, _ := findParameter(text, "{", "}")
	if result != "{world}" {
		t.Errorf("failed to extract replaceParameters {world} got %v instead", result)
	}
}

func TestReplaceParam(t *testing.T) {
	text := "hello {ref_p_start_date}"
	result := replaceParameters(text, testLookup)
	if result != "hello start date" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}

func TestReplaceTwoParams(t *testing.T) {
	text := "Can you report from {ref_p_start_date} to {ref_p_end_date}?"
	result := replaceParameters(text, testLookup)
	if result != "Can you report from start date to end date?" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}
