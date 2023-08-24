package interfaces

type Answers interface {
	ListIds() []string
	GetType(answerId string) string
	GetCode(answerId string) string
	GetLabel(answerId string) string
}

type Questions interface {
	ListIds() []string
	GetTitle(questionId string) string
	ListAnswers(questionId string) []string
}

type Sections interface {
	ListTitles() []string
	ListQuestions(sectionTitle string) []string
}

type Instrument interface {
	GetTitle() string
	GetSurveyId() string
	GetFormType() string
	GetSections() Sections
	GetQuestions() Questions
	GetAnswers() Answers
}
