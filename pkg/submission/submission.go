package submission

import (
	"encoding/json"
	"fmt"
	"sdxImage/pkg/model"
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

func From(bytes []byte) *Submission {
	var submission Submission
	err := json.Unmarshal(bytes, &submission)
	if err != nil {
		fmt.Println("Error unmarshalling submission json")
		fmt.Println(err)
	}
	return &submission
}

func Add(survey *model.Survey, submission *Submission) *model.Survey {
	survey.Respondent = submission.RuRef
	survey.SubmittedAt = submission.SubmittedAt
	for _, section := range survey.Sections {
		for _, question := range section.Questions {
			for _, a := range question.Answers {
				value, found := submission.Data[a.QCode]
				if found {
					a.Value = value
				}
			}
		}
	}
	return survey
}
