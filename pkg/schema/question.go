package schema

type Question struct {
	id           string
	Title        string
	QuestionType string
	AnswerIds    []string
	answers      []*Answer
}

func convertToQuestion(json map[string]any) (*Question, bool) {
	answers := convertList(json, "answers", convertToAnswer)
	answerIds := make([]string, len(answers))
	for i, a := range answers {
		answerIds[i] = a.Id
	}

	question := &Question{
		id:           getString(json, "id"),
		Title:        extractTitle(json),
		QuestionType: getString(json, "type"),
		AnswerIds:    answerIds,
		answers:      answers,
	}
	return question, true
}

type Questions struct {
	idList      []string
	QuestionMap map[string]*Question
}

func newQuestions() *Questions {
	return &Questions{
		idList:      []string{},
		QuestionMap: make(map[string]*Question),
	}
}

func (questions *Questions) addQuestion(question *Question) {
	questions.idList = append(questions.idList, question.id)
	questions.QuestionMap[question.id] = question
}

func (questions *Questions) ListIds() []string {
	return questions.idList
}

func (questions *Questions) GetTitle(questionId string) string {
	return questions.QuestionMap[questionId].Title
}

func (questions *Questions) ListAnswers(questionId string) []string {
	return questions.QuestionMap[questionId].AnswerIds
}
