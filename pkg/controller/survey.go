package controller

import (
	"sdxImage/pkg/model"
	"sdxImage/pkg/substitutions"
	"strconv"
	"strings"
)

func fromSubmission(schema *model.Schema, submission *model.Submission) *model.Survey {

	lookup := substitutions.GetLookup(submission.StartDate, submission.EndDate, submission.RuName)

	survey := &model.Survey{
		Title:       schema.Title,
		SurveyId:    schema.SurveyId,
		FormType:    schema.FormType,
		Respondent:  submission.RuRef,
		SubmittedAt: substitutions.DateFormat(submission.SubmittedAt),
		Sections:    []*model.Section{},
	}

	var sections []*model.Section
	for _, sect := range schema.Sections {
		hasAnswerValue := false
		section := &model.Section{
			Title:     sect.Title,
			Instances: []*model.Instance{},
		}
		instanceMap := map[string]*model.Instance{}
		for _, quest := range sect.Questions {

			title := quest.Title

			for _, ans := range quest.Answers {
				responseList := submission.GetResponses(ans.QCode)
				for _, resp := range responseList {

					instanceKey := strconv.Itoa(resp.Instance)
					instance, found := instanceMap[instanceKey]
					if !found {
						instance = &model.Instance{
							Id:      resp.Instance,
							Answers: []*model.Answer{},
						}
						instanceMap[instanceKey] = instance
						section.Instances = append(section.Instances, instance)
					}

					value := resp.Value
					if value != "" {
						qCode := getQCode(ans.QCode)
						text := getAnswerText(title, ans.Label, ans.Type, len(quest.Answers) > 1, lookup)

						answer := &model.Answer{
							QCode: qCode,
							Text:  text,
							Value: value,
						}

						instance.Answers = append(instance.Answers, answer)
						hasAnswerValue = true
					}
				}
			}
		}
		if hasAnswerValue {
			sections = append(sections, section)
		}
	}
	survey.Sections = sections
	return survey
}

func getAnswerText(title, label, qType string, multiple bool, lookup substitutions.ParameterLookup) string {
	text := title
	if qType == "Date" {
		text += " " + label
	} else if qType == "Number" {
		text += " " + label + ":"
	} else if qType == "Currency" {
		// only include text and label if the question has multiple answers
		// and the label has 5 words or fewer
		if multiple && len(strings.Split(label, " ")) <= 5 {
			text += " " + label + ":"
		} else {
			text = label + "?"
		}

	} else if qType == "Unit" {
		text += " " + label + ":"
	}
	return substitutions.Replace(text, lookup)
}

func getQCode(code string) string {
	_, err := strconv.Atoi(code)
	if err != nil {
		return getQCode(code[1:])
	}
	return code
}
