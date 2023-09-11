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
				for _, block := range group.Blocks {
					if block.BlockType == "Question" || block.BlockType == "ListCollector" {
						ids = append(ids, block.Question.Id)
					}
				}
			}
		}
	}
	return ids
}

func (schema *Schema) GetQuestionTitle(questionId string) string {
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.Blocks {
				if block.BlockType == "Question" || block.BlockType == "ListCollector" {
					if block.Question.Id == questionId {
						return string(block.Question.Title)
					}
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
			for _, block := range group.Blocks {
				if block.Question.Id == questionId && (block.BlockType == "Question" || block.BlockType == "ListCollector") {
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
			for _, block := range group.Blocks {
				if block.BlockType == "Question" || block.BlockType == "ListCollector" {
					for _, answer := range block.Question.Answers {
						if answer.Id == answerId {
							qCode := answer.Qcode
							if schema.DataVersion == LoopingDataVersion {
								qCode = schema.lookupQCode(answerId)
							}

							if len(answer.Options) > 0 {
								result := make([]interfaces.AnswerSpec, len(answer.Options))
								for i, option := range answer.Options {
									result[i] = NewAnswerSpec(answer.AnswerType, option.Qcode, option.Label)
								}
							} else {
								return []interfaces.AnswerSpec{NewAnswerSpec(answer.AnswerType, qCode, answer.Label)}
							}
						}
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
