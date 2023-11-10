package schema

import (
	"sdxImage/internal/interfaces"
	"slices"
)

func convert(schema *Schema) *CollectionInstrument {

	sectionTitles := getSectionTitles(schema)
	titleToQidMap := getTitleToQidMap(sectionTitles, schema)
	qidToQtMap := getQidToQtMap(titleToQidMap, schema)
	qidToAidMap := getQidToAidMap(titleToQidMap, schema)
	answerMap := getAnswerMap(qidToAidMap, schema)

	return &CollectionInstrument{
		title:         schema.Title,
		surveyId:      schema.SurveyId,
		formType:      schema.FormType,
		sectionTitles: sectionTitles,
		titleToQidMap: titleToQidMap,
		qidToQtMap:    qidToQtMap,
		qidToAidMap:   qidToAidMap,
		answerMap:     answerMap,
	}
}

func getSectionTitles(schema *Schema) []string {
	return listTitles(schema)
}

func getTitleToQidMap(titles []string, schema *Schema) map[string][]string {
	result := make(map[string][]string)
	for _, title := range titles {
		result[title] = listQuestionIds(title, schema)
	}
	return result
}

func getQidToQtMap(titleToQidMap map[string][]string, schema *Schema) map[string]string {
	result := make(map[string]string)
	for _, idList := range titleToQidMap {
		for _, qid := range idList {
			result[qid] = getQuestionTitle(qid, schema)
		}
	}
	return result
}

func getQidToAidMap(titleToQidMap map[string][]string, schema *Schema) map[string][]string {
	result := make(map[string][]string)
	for _, idList := range titleToQidMap {
		for _, qid := range idList {
			result[qid] = listAnswers(qid, schema)
		}
	}
	return result
}

func getAnswerMap(qidToAidMap map[string][]string, schema *Schema) map[string][]interfaces.AnswerSpec {
	result := make(map[string][]interfaces.AnswerSpec)
	for _, aidList := range qidToAidMap {
		for _, aid := range aidList {
			result[aid] = getAnswers(aid, schema)
		}
	}
	return result
}

func listTitles(schema *Schema) []string {
	var titles []string
	for _, section := range schema.Sections {
		title := string(section.Title)
		if !slices.Contains(titles, title) {
			titles = append(titles, string(section.Title))
		}
	}
	return titles
}

func listQuestionIds(title string, schema *Schema) []string {
	var ids []string
	for _, section := range schema.Sections {
		if string(section.Title) == title {
			for _, group := range section.Groups {
				for _, block := range group.getBlocks() {
					ids = append(ids, block.Question.Id)
				}
			}
		}
	}
	return ids
}

func getQuestionTitle(questionId string, schema *Schema) string {
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.getBlocks() {
				if block.Question.Id == questionId {
					return string(block.Question.Title)
				}
			}
		}
	}
	return ""
}

func listAnswers(questionId string, schema *Schema) []string {
	var answerIds []string
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.getBlocks() {
				if block.Question.Id == questionId {
					for _, answer := range block.Question.Answers {
						answerIds = append(answerIds, answer.Id)
					}
				}
			}
		}
	}
	return answerIds
}

const LoopingDataVersion = "0.0.3"

func getAnswers(answerId string, schema *Schema) []interfaces.AnswerSpec {
	for _, section := range schema.Sections {
		for _, group := range section.Groups {
			for _, block := range group.getBlocks() {
				for _, answer := range block.Question.Answers {
					if answer.Id == answerId {
						qCode := answer.Qcode
						if schema.DataVersion == LoopingDataVersion {
							qCode = lookupQCode(answerId, schema)
						}
						answer.Qcode = qCode

						return getAnswerSpecs(answer)
					}
				}
			}
		}
	}
	return nil
}

func lookupQCode(answerId string, schema *Schema) string {
	for _, answerCode := range schema.AnswerCodes {
		if answerCode.AnswerId == answerId {
			return answerCode.Code
		}
	}
	return ""
}
