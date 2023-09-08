package interfaces

type Schema interface {
	GetTitle() string
	GetSurveyId() string
	GetFormType() string

	ListTitles() []string
	ListQuestionIds(title string) []string

	GetQuestionTitle(questionId string) string
	ListAnswers(questionId string) []string

	GetAnswerType(answerId string) string
	GetAnswerCode(answerId string) string
	GetAnswerLabel(answerId string) string
}
