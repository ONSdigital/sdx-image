package drawing

import (
	"sdxImage/internal/test"
	"testing"
)

func TestSplitLongWordsOnce(t *testing.T) {
	value := "This is a long word aaaaaaaaaax"
	expected := "This is a long word aaaaaaaaaa x"
	test.Equal(t, splitLongWords(value, 10), expected)
}

func TestSplitLongWordsMultiple(t *testing.T) {
	value := "This is a long word aaaaaaaaaabbbbbbbbbbcccccccccc"
	expected := "This is a long word aaaaaaaaaa bbbbbbbbbb cccccccccc"
	test.Equal(t, splitLongWords(value, 10), expected)
}

func TestSplitTwoLongWords(t *testing.T) {
	value := "This is two long words aaaaaaaaaax and bbbbbbbbbbxx"
	expected := "This is two long words aaaaaaaaaa x and bbbbbbbbbb xx"
	test.Equal(t, splitLongWords(value, 10), expected)
}

func TestMaintainOriginalSpacing(t *testing.T) {
	value := "This is a big  space!"
	test.Equal(t, splitLongWords(value, 10), value)
}
