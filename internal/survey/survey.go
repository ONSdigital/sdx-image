// Package survey defines the data types used to capture survey information
package survey

import (
	"encoding/json"
	"fmt"
)

type Instance struct {
	Id      string
	Value   int
	Answers []*Answer
}

type Section struct {
	Title     string
	Instances map[string]*Instance
}

type Survey struct {
	Title       string
	SurveyId    string
	FormType    string
	Respondent  string
	RuName      string
	SubmittedAt string
	Sections    []*Section
	Units       []Unit
}

func (survey *Survey) String() string {
	b, err := json.MarshalIndent(survey, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
