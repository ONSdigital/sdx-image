package model

import "testing"

func TestGetSingleResp(t *testing.T) {
	respList := []*Resp{
		{"001", "Yes", "1"},
		{"002", "Yes", "1"},
		{"002", "Yes", "2"},
		{"002", "Yes", "3"},
	}

	spp := Spp{
		FormType:  "",
		Reference: "",
		Period:    "",
		Survey:    "",
		Responses: respList,
	}

	result := spp.GetResp("001")
	expected := []Resp{{"001", "Yes", "1"}}

	if result[0] != expected[0] {
		t.Errorf("failed to return resp %q, instead got: %q", expected, result)
	}
}

func TestGetMultipleResp(t *testing.T) {
	respList := []*Resp{
		{"001", "Yes", "1"},
		{"002", "Yes", "1"},
		{"002", "Yes", "2"},
		{"002", "Yes", "3"},
	}

	spp := Spp{
		FormType:  "",
		Reference: "",
		Period:    "",
		Survey:    "",
		Responses: respList,
	}

	result := spp.GetResp("002")
	expected := []Resp{{"002", "Yes", "1"}, {"002", "Yes", "2"}, {"002", "Yes", "3"}}

	if result[0] != expected[0] {
		t.Errorf("failed to return resp %q, instead got: %q", expected, result)
	}
	if result[1] != expected[1] {
		t.Errorf("failed to return resp %q, instead got: %q", expected, result)
	}
	if result[2] != expected[2] {
		t.Errorf("failed to return resp %q, instead got: %q", expected, result)
	}
}
