package substitutions

import "testing"

func TestConvertStartDate(t *testing.T) {
	startDate := "2022-01-01"
	expected := "01/01/2022"
	actual := convertDate(startDate)
	if actual != expected {
		t.Errorf("failed to convert to: (%q) instead got (%q)", expected, actual)
	}
}

func TestConvertEndDate(t *testing.T) {
	startDate := "2022-12-31"
	expected := "31/12/2022"
	actual := convertDate(startDate)
	if actual != expected {
		t.Errorf("failed to convert to: (%q) instead got (%q)", expected, actual)
	}
}

func TestConvertSubmittedAt(t *testing.T) {
	dateTime := "2023-03-03T11:36:23+00:00"
	expected := "03 March 2023 11:36:23"
	actual := convertSubmittedAt(dateTime)
	if actual != expected {
		t.Errorf("failed to convert to: (%q) instead got (%q)", expected, actual)
	}
}
