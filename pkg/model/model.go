package model

type Answer struct {
	Type  string
	QCode string
	Label string
	Value string
}

type Question struct {
	Title   string
	Answers []Answer
}

type Survey struct {
	Title     string
	SurveyId  string
	FormType  string
	Questions []Question
}
