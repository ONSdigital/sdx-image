package instrument

import (
	"sdxImage/pkg/interfaces"
	"sdxImage/pkg/schema"
)

func FromSchema(schema *schema.Schema) *CollectionInstrument {

	sectionTitles := getSectionTitles(schema)
	titleToQidMap := getTitleToQidMap(sectionTitles, schema)
	qidToQtMap := GetQidToQtMap(titleToQidMap, schema)
	qidToAidMap := GetQidToAidMap(titleToQidMap, schema)
	answerMap := getAnswerMap(qidToAidMap, schema)

	return &CollectionInstrument{
		Title:         schema.GetTitle(),
		SurveyId:      schema.GetSurveyId(),
		FormType:      schema.GetFormType(),
		SectionTitles: sectionTitles,
		TitleToQidMap: titleToQidMap,
		QidToQtMap:    qidToQtMap,
		QidToAidMap:   qidToAidMap,
		AnswerMap:     answerMap,
	}
}

func getSectionTitles(schema *schema.Schema) []string {
	return schema.ListTitles()
}

func getTitleToQidMap(titles []string, schema *schema.Schema) map[string][]string {
	result := make(map[string][]string)
	for _, title := range titles {
		result[title] = schema.ListQuestionIds(title)
	}
	return result
}

func GetQidToQtMap(titleToQidMap map[string][]string, schema *schema.Schema) map[string]string {
	result := make(map[string]string)
	for _, idList := range titleToQidMap {
		for _, qid := range idList {
			result[qid] = schema.GetQuestionTitle(qid)
		}
	}
	return result
}

func GetQidToAidMap(titleToQidMap map[string][]string, schema *schema.Schema) map[string][]string {
	result := make(map[string][]string)
	for _, idList := range titleToQidMap {
		for _, qid := range idList {
			result[qid] = schema.ListAnswers(qid)
		}
	}
	return result
}

func getAnswerMap(qidToAidMap map[string][]string, schema *schema.Schema) map[string][]interfaces.AnswerSpec {
	result := make(map[string][]interfaces.AnswerSpec)
	for _, aidList := range qidToAidMap {
		for _, aid := range aidList {
			result[aid] = schema.GetAnswers(aid)
		}
	}
	return result
}
