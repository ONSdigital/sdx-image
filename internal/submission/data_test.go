package submission

import (
	"encoding/json"
	"sdxImage/internal/test"
	"testing"
)

func TestStringValue(t *testing.T) {
	jsonStr := "{\"answer_id\":\"1\",\"value\":\"24.1\",\"list_item_id\":\"1\"}"
	answer := &Answer{}
	err := json.Unmarshal([]byte(jsonStr), answer)
	if err != nil {
		t.Errorf("failed to parse json: %v", err)
	}
	test.Equal(t, "24.1", answer.getValue())
}

func TestEmptyStringValue(t *testing.T) {
	jsonStr := "{\"answer_id\":\"1\",\"value\":\"\",\"list_item_id\":\"1\"}"
	answer := &Answer{}
	err := json.Unmarshal([]byte(jsonStr), answer)
	if err != nil {
		t.Errorf("failed to parse json: %v", err)
	}
	test.Equal(t, "", answer.getValue())
}

func TestFloatValue(t *testing.T) {
	jsonStr := "{\"answer_id\":\"1\",\"value\":24.1,\"list_item_id\":\"1\"}"
	answer := &Answer{}
	err := json.Unmarshal([]byte(jsonStr), answer)
	if err != nil {
		t.Errorf("failed to parse json: %v", err)
	}
	test.Equal(t, "24.1", answer.getValue())
}

func TestFloatValueWithTrailingZeros(t *testing.T) {
	jsonStr := "{\"answer_id\":\"1\",\"value\":24.50,\"list_item_id\":\"1\"}"
	answer := &Answer{}
	err := json.Unmarshal([]byte(jsonStr), answer)
	if err != nil {
		t.Errorf("failed to parse json!")
	}
	test.Equal(t, "24.50", answer.getValue())
}

func TestArrayValue(t *testing.T) {
	jsonStr := "{\"answer_id\":\"1\",\"value\":[\"a\",\"b\",\"c\"],\"list_item_id\":\"1\"}"
	answer := &Answer{}
	err := json.Unmarshal([]byte(jsonStr), answer)
	if err != nil {
		t.Errorf("failed to parse json!")
	}
	test.Equal(t, "\"a\",\"b\",\"c\"", answer.getValue())
}
