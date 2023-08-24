package read

import (
	"encoding/json"
	"fmt"
	"sdxImage/pkg/interfaces"
)

type SurveyMetaData struct {
	RuRef     string `json:"ru_ref"`
	RuName    string `json:"ru_name"`
	StartDate string `json:"ref_p_start_date"`
	EndDate   string `json:"ref_p_end_date"`
}

type V2Submission struct {
	TxId           string          `json:"tx_id"`
	SchemaName     string          `json:"schema_name"`
	SurveyMetadata *SurveyMetaData `json:"survey_metadata"`
	SubmittedAt    string          `json:"submitted_at"`
	DataVersion    string          `json:"data_version"`
	Data           *Data           `json:"data"`
}

func (submission *V2Submission) GetTxId() string {
	return submission.TxId
}

func (submission *V2Submission) GetSchemaName() string {
	return submission.SchemaName
}

func (submission *V2Submission) GetRuRef() string {
	return submission.SurveyMetadata.RuRef
}

func (submission *V2Submission) GetRuName() string {
	return submission.SurveyMetadata.RuName
}
func (submission *V2Submission) GetSubmittedAt() string {
	return submission.SubmittedAt
}

func (submission *V2Submission) GetStartDate() string {
	return submission.SurveyMetadata.StartDate
}

func (submission *V2Submission) GetEndDate() string {
	return submission.SurveyMetadata.EndDate
}

func (submission *V2Submission) GetDataVersion() string {
	return submission.DataVersion
}

func (submission *V2Submission) GetData() []interfaces.Data {
	result := make([]interfaces.Data, len(submission.Data.Responses))
	for i, r := range submission.Data.Responses {
		result[i] = r
	}
	return result
}

func (submission *V2Submission) String() string {
	b, err := json.MarshalIndent(submission, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
