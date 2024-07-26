package submission

import (
	"encoding/json"
	"fmt"
)

type SurveyMetaData struct {
	RuRef          string `json:"ru_ref"`
	RuName         string `json:"ru_name"`
	StartDate      string `json:"ref_p_start_date"`
	EndDate        string `json:"ref_p_end_date"`
	EmploymentDate string `json:"employment_date"`
}

type Submission struct {
	TxId           string          `json:"tx_id"`
	SchemaName     string          `json:"schema_name"`
	SubmittedAt    string          `json:"submitted_at"`
	SurveyMetadata *SurveyMetaData `json:"survey_metadata"`
	Data           Data            `json:"data"`
}

func (submission *Submission) String() string {
	b, err := json.MarshalIndent(submission, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
