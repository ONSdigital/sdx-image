package survey

import "testing"

const FakeSurveyId = "666"

func TestGetQCode(t *testing.T) {
	result := getQCode("200", FakeSurveyId)
	expected := "200"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestRemoveLetters(t *testing.T) {
	result := removeLetters("c200")
	expected := "200"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestRemoveLettersAndNumbers(t *testing.T) {
	result := removeLetters("42c200")
	expected := "200"
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
