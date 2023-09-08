package schema

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
		for _, group := range section.Groups {
			titles = append(titles, string(group.Title))
		}
	}
	return titles
}

func (schema *Schema) ListQuestionIds(title string) []string {
	var ids []string
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			if string(group.Title) == title {
				for _, block := range group.Blocks {
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
			for _, block := range group.Blocks {
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
			for _, block := range group.Blocks {
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

func (schema *Schema) GetAnswerType(answerId string) string {
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.Blocks {
				for _, answer := range block.Question.Answers {
					if answer.Id == answerId {
						return answer.AnswerType
					}
				}
			}
		}
	}
	return ""
}
func (schema *Schema) GetAnswerCode(answerId string) string {
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.Blocks {
				for _, answer := range block.Question.Answers {
					if answer.Id == answerId {
						return answer.Qcode
					}
				}
			}
		}
	}
	return ""
}

func (schema *Schema) GetAnswerLabel(answerId string) string {
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.Blocks {
				for _, answer := range block.Question.Answers {
					if answer.Id == answerId {
						return answer.Label
					}
				}
			}
		}
	}
	return ""
}
