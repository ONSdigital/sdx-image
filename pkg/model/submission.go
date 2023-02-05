package model

import (
	"encoding/json"
	"fmt"
)

type SurveyMetaData struct {
	SurveyId  string `json:"survey_id"`
	Period    string `json:"period_id"`
	RuRef     string `json:"ru_ref"`
	RuName    string `json:"ru_name"`
	FormType  string `json:"form_type"`
	StartDate string `json:"ref_p_start_date"`
	EndDate   string `json:"ref_p_end_date"`
}

type Submission struct {
	TxId           string `json:"tx_id"`
	SubmittedAt    string `json:"submitted_at"`
	SchemaName     string `json:"schema_name"`
	SurveyMetaData `json:"survey_metadata"`
	Data           map[string]string `json:"data"`
}

func Convert(bytes []byte) *Submission {
	var submission Submission
	err := json.Unmarshal(bytes, &submission)
	if err != nil {
		fmt.Println("Error unmarshalling submission json")
		fmt.Println(err)
	}
	return &submission
}
