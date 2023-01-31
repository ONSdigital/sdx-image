package schema

import "testing"

func TestGetParameter(t *testing.T) {
	text := "hello {world}"
	result, _ := getParameter(text)
	if result != "{world}" {
		t.Errorf("failed to extract parameter {world} got %v instead", result)
	}
}

func TestReplace(t *testing.T) {
	text := "hello {world}"
	result := replace(text, map[string]string{"world": "everybody"})
	if result != "hello everybody" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}

func TestReplaceTwoParams(t *testing.T) {
	text := "Can you report from {ref_p_start_date} to {ref_p_end_date}?"
	result := replace(text, map[string]string{"ref_p_start_date": "start_date", "ref_p_end_date": "end_date"})
	if result != "Can you report from start_date to end_date?" {
		t.Errorf("failed to replce: (%q) instead got (%q)", text, result)
	}
}
