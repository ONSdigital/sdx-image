package interfaces

type Answers interface {
	GetAnswerIds() []string
	GetAnswerType(answerId string) string
	GetAnswerCode(answerId string) string
	GetAnswerLabel(answerId string) string
}

type Questions interface {
	GetQuestionIds() []string
	GetQuestionTitle(questionId string) string
	GetQuestionAnswers(questionId string) []string
}

type Sections interface {
	GetSectionTitles() []string
	GetSectionQuestions(sectionTitle string) []string
}

type Instrument interface {
	GetTitle() string
	GetSurveyId() string
	GetFormType() string
	GetSections() Sections
	GetQuestions() Questions
	GetAnswers() Answers
	Print()
}
