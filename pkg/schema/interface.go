package schema

import "sdxImage/pkg/interfaces"

func (schema *Schema) GetTitle() string {
	return schema.Title
}

func (schema *Schema) GetSurveyId() string {
	return schema.SurveyId
}

func (schema *Schema) GetFormType() string {
	return schema.FormType
}

func (schema *Schema) ListTitles() []string {
	var titles []string
	for _, section := range schema.Sections {
		titles = append(titles, string(section.Title))
	}
	return titles
}

func (schema *Schema) ListQuestionIds(title string) []string {
	var ids []string
	for _, section := range schema.Sections {
		if string(section.Title) == title {
			for _, group := range section.Groups {
				for _, block := range group.getBlocks() {
					ids = append(ids, block.Question.Id)
				}
			}
		}
	}
	return ids
}

func (schema *Schema) GetQuestionTitle(questionId string) string {
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.getBlocks() {
				if block.Question.Id == questionId {
					return string(block.Question.Title)
				}
			}
		}
	}
	return ""
}
func (schema *Schema) ListAnswers(questionId string) []string {
	var answerIds []string
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.getBlocks() {
				if block.Question.Id == questionId {
					for _, answer := range block.Question.Answers {
						answerIds = append(answerIds, answer.Id)
					}
				}
			}
		}
	}
	return answerIds
}

const LoopingDataVersion = "0.0.3"

func (schema *Schema) GetAnswers(answerId string) []interfaces.AnswerSpec {
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.getBlocks() {
				for _, answer := range block.Question.Answers {
					if answer.Id == answerId {
						qCode := answer.Qcode
						if schema.DataVersion == LoopingDataVersion {
							qCode = schema.lookupQCode(answerId)
						}
						answer.Qcode = qCode

						return getAnswerSpecs(answer)
					}
				}
			}
		}
	}
	return nil
}

func (schema *Schema) lookupQCode(answerId string) string {
	for _, answerCode := range schema.AnswerCodes {
		if answerCode.AnswerId == answerId {
			return answerCode.Code
		}
	}
	return ""
}
