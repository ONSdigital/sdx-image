package schema

import (
	"fmt"
	"sdxImage/pkg/interfaces"
)

type Schema struct {
	Title       string
	SurveyId    string
	FormType    string
	DataVersion string
	Sections    interfaces.Sections
	Questions   interfaces.Questions
	Answers     interfaces.Answers
}

func convert(json map[string]any) *Schema {
	tree := convertList(json, "sections", convertToSection)

	sections := newSections()
	questions := newQuestions()
	answers := newAnswers()

	for _, section := range tree {
		sections.addSection(section)
		for _, group := range section.Groups {
			for _, block := range group.Blocks {
				questions.addQuestion(block.Question)
				for _, answer := range block.Question.Answers {
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

		for id, answer := range answers.answerMap {
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

func (schema *Schema) Print() {
	fmt.Print("SurveyId: ")
	fmt.Println(schema.SurveyId)
	fmt.Print("FormType: ")
	fmt.Println(schema.FormType)
	fmt.Print("DataVersion: ")
	fmt.Println(schema.DataVersion)

	fmt.Print("Sections: ")
	fmt.Println(fmt.Sprint(schema.Sections.GetSectionIds()))

	fmt.Print("Questions: ")
	fmt.Println(fmt.Sprint(schema.Questions.GetQuestionIds()))

	fmt.Print("Answers: ")
	fmt.Println(fmt.Sprint(schema.Answers.GetAnswerIds()))
}
