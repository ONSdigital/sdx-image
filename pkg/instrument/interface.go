package instrument

import (
	"sdxImage/pkg/interfaces"
)

func (ci *CollectionInstrument) GetTitle() string {
	return ci.Title
}

func (ci *CollectionInstrument) GetSurveyId() string {
	return ci.SurveyId
}

func (ci *CollectionInstrument) GetFormType() string {
	return ci.FormType
}

func (ci *CollectionInstrument) ListTitles() []string {
	return ci.SectionTitles
}

func (ci *CollectionInstrument) ListQuestionIds(title string) []string {
	return ci.TitleToQidMap[title]
}

func (ci *CollectionInstrument) GetQuestionTitle(questionId string) string {
	return ci.QidToQtMap[questionId]
}

func (ci *CollectionInstrument) ListAnswers(questionId string) []string {
	return ci.QidToAidMap[questionId]
}

func (ci *CollectionInstrument) GetAnswers(answerId string) []interfaces.AnswerSpec {
	return ci.AnswerMap[answerId]
}
