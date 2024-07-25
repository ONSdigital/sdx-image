// Package survey defines the data types used to capture survey information
package survey

import (
	"encoding/json"
	"fmt"
)

type Instance struct {
	Id      int
	Answers []*Answer
}

type Section struct {
	Title     string
	Instances []*Instance
}

type Survey struct {
	Title       string
	SurveyId    string
	FormType    string
	Respondent  string
	RuName      string
	SubmittedAt string
	Sections    []*Section
	LocalUnits  []*LocalUnit
}

func (survey *Survey) String() string {
	b, err := json.MarshalIndent(survey, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
