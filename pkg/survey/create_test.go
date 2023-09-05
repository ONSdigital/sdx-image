package survey

import "testing"

func TestGetQCode(t *testing.T) {
	result := getQCode("200")
	expected := "200"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestGetQCodeWithLetters(t *testing.T) {
	result := getQCode("c200")
	expected := "200"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestGetQCodeWithLettersAndNumbers(t *testing.T) {
	result := getQCode("42c200")
	expected := "200"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
