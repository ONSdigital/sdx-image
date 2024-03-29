package submission

import (
	"encoding/json"
	"fmt"
	"sdxImage/internal/interfaces"
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
	Data        Data        `json:"data"`
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
	return "NA"
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

func (submission *V1Submission) GetEmploymentDate() string {
	return "the date of employment"
}

func (submission *V1Submission) GetResponses(code string) []interfaces.Response {
	return submission.Data[code]
}

func (submission *V1Submission) GetLocalUnits() []interfaces.LocalUnit {
	//v1 submissions do not have supplementary data
	return []interfaces.LocalUnit{}
}

func (submission *V1Submission) String() string {
	b, err := json.MarshalIndent(submission, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
