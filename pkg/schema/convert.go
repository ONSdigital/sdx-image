package schema

import "sdxImage/pkg/model"

func toMap(x any) map[string]any {
	result, _ := x.(map[string]any)
	return result
}

func getFieldFrom[T any](json map[string]any, fieldName ...string) T {
	result, _ := json[fieldName[0]].(T)
	return result
}

func locateStringFrom(json map[string]any, fieldName ...string) string {
	result, ok := json[fieldName[0]].(string)
	if !ok {
		nextLevel := json[fieldName[0]].(map[string]any)
		return locateStringFrom(nextLevel, fieldName[1:]...)
	}
	return result
}

var getStringFrom = getFieldFrom[string]
var getListFrom = getFieldFrom[[]any]
var getMapFrom = getFieldFrom[map[string]any]

func getOptionalStringField(json map[string]any, fieldName string) string {
	result, found := json[fieldName].(string)
	if !found {
		return ""
	}
	return result
}

func toSurvey(m map[string]any) *model.Survey {
	title := getStringFrom(m, "title")
	surveyId := getStringFrom(m, "survey_id")
	formType := getStringFrom(m, "form_type")
	sections := getListFrom(m, "sections")
	survey := model.Survey{
		Title:    title,
		SurveyId: surveyId,
		FormType: formType,
		Sections: []*model.Section{},
	}

	for _, s := range sections {
		sect := toMap(s)
		section := &model.Section{
			Title:     getOptionalStringField(sect, "title"),
			Questions: []*model.Question{},
		}
		groups := getListFrom(sect, "groups")
		for _, g := range groups {
			group := toMap(g)
			if getStringFrom(group, "title") != "Introduction" {
				blocks := getListFrom(group, "blocks")
				for _, b := range blocks {
					block := toMap(b)
					if getStringFrom(block, "type") == "Question" {
						q := getMapFrom(block, "question")

						question := &model.Question{
							Title:   locateStringFrom(q, "title", "text"),
							Answers: []*model.Answer{},
						}

						answers := getListFrom(q, "answers")
						for _, a := range answers {
							ans := toMap(a)
							label, exists := ans["label"]
							if !exists {
								label = "label"
							}
							answer := &model.Answer{
								Type:  getStringFrom(ans, "type"),
								QCode: getStringFrom(ans, "q_code"),
								Label: label.(string),
							}
							question.Answers = append(question.Answers, answer)
						}

						section.Questions = append(section.Questions, question)
					}
				}
			}
		}
		survey.Sections = append(survey.Sections, section)
	}
	return &survey
}
