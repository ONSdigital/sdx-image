package schema

import (
	"sdxImage/pkg/interfaces"
)

type CollectionInstrument struct {
	title         string
	surveyId      string
	formType      string
	sectionTitles []string
	titleToQidMap map[string][]string
	qidToQtMap    map[string]string
	qidToAidMap   map[string][]string
	answerMap     map[string][]interfaces.AnswerSpec
}

func (ci *CollectionInstrument) GetTitle() string {
	return ci.title
}

func (ci *CollectionInstrument) GetSurveyId() string {
	return ci.surveyId
}

func (ci *CollectionInstrument) GetFormType() string {
	return ci.formType
}

func (ci *CollectionInstrument) ListTitles() []string {
	return ci.sectionTitles
}

func (ci *CollectionInstrument) ListQuestionIds(title string) []string {
	return ci.titleToQidMap[title]
}

func (ci *CollectionInstrument) GetQuestionTitle(questionId string) string {
	return ci.qidToQtMap[questionId]
}

func (ci *CollectionInstrument) ListAnswers(questionId string) []string {
	return ci.qidToAidMap[questionId]
}

func (ci *CollectionInstrument) GetAnswers(answerId string) []interfaces.AnswerSpec {
	return ci.answerMap[answerId]
}
