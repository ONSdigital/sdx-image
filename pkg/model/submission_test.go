package model

import (
	"testing"
)

func TestNonMissing(t *testing.T) {
	submission := &Submission{
		TxId:        "c2ca5e03-ff68-4d5d-b5a3-d96d3be45cfb",
		SchemaName:  "mbs_0106",
		RuRef:       "12346789012A",
		RuName:      "MegaCorp",
		SubmittedAt: "2023-02-11T16:57:11+00:00",
		StartDate:   "2016-05-01",
		EndDate:     "2016-05-31",
		DataVersion: "0.0.1",
		Responses:   []*Response{{"40", "56000", 0}},
	}

	missing := MissingFields(submission)
	if len(missing) > 0 {
		t.Errorf("Should be no missing fields: %v", missing)
	}
}

func TestMissingRuName(t *testing.T) {
	submission := &Submission{
		TxId:        "c2ca5e03-ff68-4d5d-b5a3-d96d3be45cfb",
		SchemaName:  "mbs_0106",
		RuRef:       "12346789012A",
		SubmittedAt: "2023-02-11T16:57:11+00:00",
		StartDate:   "2016-05-01",
		EndDate:     "2016-05-31",
		DataVersion: "0.0.1",
		Responses:   []*Response{{"40", "56000", 0}},
	}

	missing := MissingFields(submission)
	if len(missing) != 1 {
		t.Errorf("Should have found 1 missing value: %v", missing)
	}
}

func TestGetSingleResponse(t *testing.T) {
	respList := []*Response{
		{"001", "Yes", 1},
		{"002", "Yes", 1},
		{"002", "Yes", 2},
		{"002", "Yes", 3},
	}

	submission := &Submission{
		TxId:        "c2ca5e03-ff68-4d5d-b5a3-d96d3be45cfb",
		SchemaName:  "mbs_0106",
		RuRef:       "12346789012A",
		SubmittedAt: "2023-02-11T16:57:11+00:00",
		StartDate:   "2016-05-01",
		EndDate:     "2016-05-31",
		DataVersion: "0.0.3",
		Responses:   respList,
	}

	result := submission.GetResponses("001")
	expected := []Response{{"001", "Yes", 1}}

	if result[0] != expected[0] {
		t.Errorf("failed to return resp %q, instead got: %q", expected, result)
	}
}

func TestGetMultipleResponses(t *testing.T) {
	respList := []*Response{
		{"001", "Yes", 1},
		{"002", "Yes", 1},
		{"002", "Yes", 2},
		{"002", "Yes", 3},
	}

	submission := &Submission{
		TxId:        "c2ca5e03-ff68-4d5d-b5a3-d96d3be45cfb",
		SchemaName:  "mbs_0106",
		RuRef:       "12346789012A",
		SubmittedAt: "2023-02-11T16:57:11+00:00",
		StartDate:   "2016-05-01",
		EndDate:     "2016-05-31",
		DataVersion: "0.0.3",
		Responses:   respList,
	}

	result := submission.GetResponses("002")
	expected := []Response{{"002", "Yes", 1}, {"002", "Yes", 2}, {"002", "Yes", 3}}

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
