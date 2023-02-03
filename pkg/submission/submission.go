package submission

import (
	"encoding/json"
	"sdxImage/pkg/model"
	"sdxImage/pkg/schema"
)

type SurveyMetaData struct {
	SurveyId string `json:"survey_id"`
	Period   string `json:"period_id"`
	RuRef    string `json:"ru_ref"`
	RuName   string `json:"ru_name"`
	FormType string `json:"form_type"`
}

type Submission struct {
	TxId           string `json:"tx_id"`
	SubmittedAt    string `json:"submitted_at"`
	SchemaName     string `json:"schema_name"`
	SurveyMetaData `json:"survey_metadata"`
	Data           map[string]string `json:"data"`
}

func GetSurvey(submissionBytes []byte) *model.Survey {
	submission, ok := read(submissionBytes)
	if !ok {
		return nil
	}
	survey := schema.Read(submission.SchemaName)
	survey.Respondent = submission.SurveyMetaData.RuRef
	survey.SubmittedAt = submission.SubmittedAt
	return survey
}

func read(bytes []byte) (*Submission, bool) {
	var submission *Submission
	err := json.Unmarshal(bytes, submission)
	if err != nil {
		return nil, false
	}
	return submission, true
}
