package schema

import (
	"sdxImage/internal/interfaces"
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
	titles := make([]string, len(ci.sectionTitles))
	copy(titles, ci.sectionTitles)
	return titles
}

func (ci *CollectionInstrument) ListQuestionIds(title string) []string {
	ids := ci.titleToQidMap[title]
	result := make([]string, len(ids))
	copy(result, ids)
	return result
}

func (ci *CollectionInstrument) GetQuestionTitle(questionId string) string {
	return ci.qidToQtMap[questionId]
}

func (ci *CollectionInstrument) ListAnswerIds(questionId string) []string {
	answers := ci.qidToAidMap[questionId]
	result := make([]string, len(answers))
	copy(result, answers)
	return result
}

func (ci *CollectionInstrument) GetAnswers(answerId string) []interfaces.AnswerSpec {
	answers := ci.answerMap[answerId]
	result := make([]interfaces.AnswerSpec, len(answers))
	copy(result, answers)
	return result
}
