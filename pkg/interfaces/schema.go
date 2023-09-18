package interfaces

type AnswerSpec interface {
	GetType() string
	GetCode() string
	GetLabel() string
}

type Schema interface {
	GetTitle() string
	GetSurveyId() string
	GetFormType() string

	ListTitles() []string
	ListQuestionIds(title string) []string

	GetQuestionTitle(questionId string) string
	ListAnswers(questionId string) []string

	GetAnswers(answerId string) []AnswerSpec
}
