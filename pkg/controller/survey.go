package controller

import (
	"fmt"
	"sdxImage/pkg/model"
	"sdxImage/pkg/substitutions"
	"strconv"
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

					//**********
					if resp.QuestionCode == "c202" {
						fmt.Print(resp.Instance)
					}

					//***********

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
						qCode := ans.QCode
						text := getAnswerText(title, ans.Label, ans.Type, lookup)

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

func getAnswerText(title, label, qType string, lookup substitutions.ParameterLookup) string {
	text := title
	if qType == "Date" {
		text += " " + label
	} else if qType == "Number" {
		text += " " + label + ":"
	} else if qType == "Currency" {
		text = label + "?"
	}
	return substitutions.Replace(text, lookup)
}
