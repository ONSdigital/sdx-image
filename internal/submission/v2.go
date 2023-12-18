package submission

import (
	"encoding/json"
	"fmt"
	"sdxImage/internal/interfaces"
)

type SurveyMetaData struct {
	RuRef          string `json:"ru_ref"`
	RuName         string `json:"ru_name"`
	StartDate      string `json:"ref_p_start_date"`
	EndDate        string `json:"ref_p_end_date"`
	EmploymentDate string `json:"employment_date"`
}

type V2Submission struct {
	TxId           string          `json:"tx_id"`
	SchemaName     string          `json:"schema_name"`
	SurveyMetadata *SurveyMetaData `json:"survey_metadata"`
	SubmittedAt    string          `json:"submitted_at"`
	Version        string          `json:"version"`
	DataVersion    string          `json:"data_version"`
	Data           Data            `json:"data"`
	Supplementary  Supplementary   `json:"supplementary_data"`
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

func (submission *V2Submission) GetEmploymentDate() string {
	if submission.SurveyMetadata.EmploymentDate == "" {
		return "the date of employment"
	}
	return submission.SurveyMetadata.EmploymentDate
}

func (submission *V2Submission) GetResponses(code string) []interfaces.Response {
	return submission.Data[code]
}

func (submission *V2Submission) GetLocalUnits() []interfaces.LocalUnit {
	luList := make([]interfaces.LocalUnit, len(submission.Supplementary.Items.LocalUnits))
	for i, lu := range submission.Supplementary.Items.LocalUnits {
		unit := NewUnit(*lu)
		for _, responseList := range submission.Data {
			for _, response := range responseList {
				if response.GetSdIdentifier() == unit.GetIdentifier() {
					unit.AddResponses(response)
				}
			}
		}
		luList[i] = unit
	}
	return luList
}

func (submission *V2Submission) String() string {
	b, err := json.MarshalIndent(submission, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
