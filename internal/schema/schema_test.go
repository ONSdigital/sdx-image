package schema

import (
	"encoding/json"
	"sdxImage/internal/test"
	"testing"
)

func TestUnmarshalVariableString(t *testing.T) {

	type person struct {
		FirstName VariableString `json:"first_name"`
		Surname   VariableString `json:"surname"`
	}

	p := &person{}
	byteValue := []byte("{\"first_name\":\"John\",\"surname\":{\"text\":\"Doe\",\"double_barrelled\":false}}")
	err := json.Unmarshal(byteValue, p)
	if err != nil {
		t.Errorf("Unable to decode json: %v", err)
	}

	test.Equal(t, "John", p.FirstName)
	test.Equal(t, "Doe", p.Surname)
}
