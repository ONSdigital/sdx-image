// Package schema handles all interactions with a schema.
// A schema defines the questions present within a survey.
package schema

import (
	"encoding/json"
	"fmt"
	"sdxImage/pkg/interfaces"
)

type Schema struct {
	Title       string
	SurveyId    string
	FormType    string
	DataVersion string
	Sections    *Sections
	Questions   *Questions
	Answers     *Answers
}

func convert(json map[string]any) *Schema {
	tree := convertList(json, "sections", convertToSection)

	sections := newSections()
	questions := newQuestions()
	answers := newAnswers()

	for _, section := range tree {
		sections.addSection(section)
		for _, group := range section.groups {
			for _, block := range group.Blocks {
				questions.addQuestion(block.Question)
				for _, answer := range block.Question.answers {
					answers.addAnswer(answer)
				}
			}
		}
	}

	if _, exists := json["answer_codes"]; exists {
		answerCodes := convertList(json, "answer_codes", convertToAnswerCode)
		qCodeMap := map[string]string{}
		for _, answerCode := range answerCodes {
			qCodeMap[answerCode.AnswerId] = answerCode.Code
		}

		for id, answer := range answers.AnswerMap {
			if answer.Qcode == "" {
				answer.Qcode = qCodeMap[id]
			}
		}
	}

	schema := &Schema{
		Title:       extractTitle(json),
		SurveyId:    getString(json, "survey_id"),
		FormType:    getString(json, "form_type"),
		DataVersion: getString(json, "data_version"),
		Sections:    sections,
		Questions:   questions,
		Answers:     answers,
	}

	return schema
}

func (schema *Schema) GetTitle() string {
	return schema.Title
}

func (schema *Schema) GetSurveyId() string {
	return schema.SurveyId
}

func (schema *Schema) GetFormType() string {
	return schema.FormType
}

func (schema *Schema) GetSections() interfaces.Sections {
	return schema.Sections
}

func (schema *Schema) GetQuestions() interfaces.Questions {
	return schema.Questions
}

func (schema *Schema) GetAnswers() interfaces.Answers {
	return schema.Answers
}

func (schema *Schema) String() string {
	b, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
