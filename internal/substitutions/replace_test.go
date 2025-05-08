package substitutions

import (
	"testing"
)

func TestReplace(t *testing.T) {
	lookup := GetLookup("2016-01-01", "2016-12-31", "the business", "2016-05-01")
	text := "Of the <em>{answer1530a741_2540_4c9b_9021_34f95288e671}</em> total employees employed on {employment_date}"
	result := Replace(text, lookup)
	expected := "Of the total employees employed on 2016-05-01"
	if result != expected {
		t.Errorf("failed!")
	}
}
