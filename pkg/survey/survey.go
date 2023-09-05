// Package survey defines the data types used to capture survey information
package survey

import (
	"encoding/json"
	"fmt"
	"sdxImage/pkg/interfaces"
)

type Survey struct {
	Title       string
	SurveyId    string
	FormType    string
	Respondent  string
	SubmittedAt string
	Sections    []interfaces.Section
}

func (survey *Survey) String() string {
	b, err := json.MarshalIndent(survey, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}

func (survey *Survey) GetTitle() string {
	return survey.Title
}

func (survey *Survey) GetSurveyId() string {
	return survey.SurveyId
}

func (survey *Survey) GetFormType() string {
	return survey.FormType
}

func (survey *Survey) GetRespondent() string {
	return survey.Respondent
}

func (survey *Survey) GetSubmittedAt() string {
	return survey.SubmittedAt
}

func (survey *Survey) GetSections() []interfaces.Section {
	return survey.Sections
}
