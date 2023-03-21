// Package model defines the data types used to capture survey information
package model

import (
	"encoding/json"
	"fmt"
)

type Answer struct {
	QCode string
	Text  string
	Value string
}

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
	SubmittedAt string
	Sections    []*Section
}

func (survey *Survey) String() string {
	b, err := json.MarshalIndent(survey, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
