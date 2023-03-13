package read

import (
	"testing"
)

func TestGetQCode(t *testing.T) {
	code := "27"
	expected := "27"
	actual := getQcode(code)
	if expected != actual {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestGetQCodeWithLetter(t *testing.T) {
	code := "e27"
	expected := "27"
	actual := getQcode(code)
	if expected != actual {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func TestGetQCodeWithLetterAndNumber(t *testing.T) {
	code := "14e27"
	expected := "27"
	actual := getQcode(code)
	if expected != actual {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
