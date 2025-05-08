package submission

import (
	"encoding/json"
	"sdxImage/internal/test"
	"testing"
)

func TestFloatValueNoTrailingZeros(t *testing.T) {
	jsonStr := "{\"answer_id\":\"1\",\"value\":24.1,\"list_item_id\":\"1\"}"
	answer := &Answer{}
	err := json.Unmarshal([]byte(jsonStr), answer)
	if err != nil {
		t.Errorf("failed to parse json!")
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
	test.Equal(t, "24.5", answer.getValue())
}
