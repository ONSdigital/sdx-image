package schema

type Question struct {
	Id           string
	Title        string
	QuestionType string
	Answers      []*Answer
}

func convertToQuestion(json map[string]any) (*Question, bool) {
	question := &Question{
		Id:           getString(json, "id"),
		Title:        extractTitle(json),
		QuestionType: getString(json, "type"),
		Answers:      convertList(json, "answers", convertToAnswer),
	}
	return question, true
}

type Questions struct {
	idList      []string
	questionMap map[string]*Question
}

func newQuestions() *Questions {
	return &Questions{
		idList:      []string{},
		questionMap: make(map[string]*Question),
	}
}

func (questions *Questions) addQuestion(question *Question) {
	questions.idList = append(questions.idList, question.Id)
	questions.questionMap[question.Id] = question
}

func (questions *Questions) GetQuestionIds() []string {
	return questions.idList
}

func (questions *Questions) GetQuestionTitle(questionId string) string {
	return questions.questionMap[questionId].Title
}

func (questions *Questions) GetQuestionAnswers(questionId string) []string {
	answers := make([]string, len(questions.questionMap[questionId].Answers))
	for i, a := range questions.questionMap[questionId].Answers {
		answers[i] = a.Id
	}
	return answers
}
