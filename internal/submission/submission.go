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
	SurveyMetadata *SurveyMetaData `json:"survey_metadata"`
	SubmittedAt    string          `json:"submitted_at"`
	Version        string          `json:"version"`
	DataVersion    string          `json:"data_version"`
	Data           Data            `json:"data"`
	Supplementary  Supplementary   `json:"supplementary_data"`
}

func (submission *Submission) GetTxId() string {
	return submission.TxId
}

func (submission *Submission) GetSchemaName() string {
	return submission.SchemaName
}

func (submission *Submission) GetRuRef() string {
	return submission.SurveyMetadata.RuRef
}

func (submission *Submission) GetRuName() string {
	return submission.SurveyMetadata.RuName
}
func (submission *Submission) GetSubmittedAt() string {
	return submission.SubmittedAt
}

func (submission *Submission) GetStartDate() string {
	return submission.SurveyMetadata.StartDate
}

func (submission *Submission) GetEndDate() string {
	return submission.SurveyMetadata.EndDate
}

func (submission *Submission) GetDataVersion() string {
	return submission.DataVersion
}

func (submission *Submission) GetEmploymentDate() string {
	if submission.SurveyMetadata.EmploymentDate == "" {
		return "the date of employment"
	}
	return submission.SurveyMetadata.EmploymentDate
}

func (submission *Submission) GetResponses(code string) []Response {
	return submission.Data[code]
}

func (submission *Submission) String() string {
	b, err := json.MarshalIndent(submission, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
