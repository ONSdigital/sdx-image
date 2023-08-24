package read

import (
	"encoding/json"
	"fmt"
	"sdxImage/pkg/interfaces"
)

type Collection struct {
	SchemaName string `json:"schema_name"`
}

type V1MetaData struct {
	RuRef     string `json:"ru_ref"`
	StartDate string `json:"ref_period_start_date"`
	EndDate   string `json:"ref_period_end_date"`
}

type V1Submission struct {
	TxId        string      `json:"tx_id"`
	Collection  *Collection `json:"collection"`
	Metadata    *V1MetaData `json:"metadata"`
	SubmittedAt string      `json:"submitted_at"`
	DataVersion string      `json:"version"`
	Data        *Data       `json:"data"`
}

func (submission *V1Submission) GetTxId() string {
	return submission.TxId
}

func (submission *V1Submission) GetSchemaName() string {
	return submission.Collection.SchemaName
}

func (submission *V1Submission) GetRuRef() string {
	return submission.Metadata.RuRef
}

func (submission *V1Submission) GetRuName() string {
	return "the business"
}
func (submission *V1Submission) GetSubmittedAt() string {
	return submission.SubmittedAt
}

func (submission *V1Submission) GetStartDate() string {
	return submission.Metadata.StartDate
}

func (submission *V1Submission) GetEndDate() string {
	return submission.Metadata.EndDate
}

func (submission *V1Submission) GetDataVersion() string {
	return submission.DataVersion
}

func (submission *V1Submission) GetData() []interfaces.Data {
	result := make([]interfaces.Data, len(submission.Data.Responses))
	for i, r := range submission.Data.Responses {
		result[i] = r
	}
	return result
}

func (submission *V1Submission) String() string {
	b, err := json.MarshalIndent(submission, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
