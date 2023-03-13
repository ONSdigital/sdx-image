// Package model defines the data types used to capture survey information
package model

import (
	"encoding/json"
	"fmt"
)

type Answer struct {
	Type  string
	QCode string
	Label string
	Value string
}

type Question struct {
	Title   string
	Answers []*Answer
}

type Instance struct {
	Id        int
	Questions []*Question
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

func From(schema *Schema, submission *Submission) *Survey {
	survey := &Survey{
		Title:       schema.Title,
		SurveyId:    schema.SurveyId,
		FormType:    schema.FormType,
		Respondent:  submission.RuRef,
		SubmittedAt: submission.SubmittedAt,
		Sections:    []*Section{},
	}
	for _, sect := range schema.Sections {
		hasAnswerValue := false
		instance := &Instance{
			Id:        0,
			Questions: []*Question{},
		}
		section := &Section{
			Title:     sect.Title,
			Instances: []*Instance{instance},
		}

		for _, quest := range sect.Questions {
			question := &Question{
				Title:   quest.Title,
				Answers: []*Answer{},
			}
			for _, ans := range quest.Answers {
				answer := &Answer{
					Type:  ans.Type,
					QCode: ans.QCode,
					Label: ans.Label,
					Value: "",
				}
				value, found := submission.Data[ans.QCode]
				if found {
					answer.Value = value
					hasAnswerValue = true
				}
				question.Answers = append(question.Answers, answer)
			}
			instance.Questions = append(instance.Questions, question)
		}
		if hasAnswerValue {
			survey.Sections = append(survey.Sections, section)
		}
	}
	return survey
}
