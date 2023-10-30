package instrument

import (
	"sdxImage/pkg/interfaces"
)

type CollectionInstrument struct {
	Title         string
	SurveyId      string
	FormType      string
	SectionTitles []string
	TitleToQidMap map[string][]string
	QidToQtMap    map[string]string
	QidToAidMap   map[string][]string
	AnswerMap     map[string][]interfaces.AnswerSpec
}
