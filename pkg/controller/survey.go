package controller

import (
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
	for _, sect := range schema.Sections {
		hasAnswerValue := false
		instance := &model.Instance{
			Id:      0,
			Answers: []*model.Answer{},
		}
		section := &model.Section{
			Title:     sect.Title,
			Instances: []*model.Instance{instance},
		}

		for _, quest := range sect.Questions {

			title := quest.Title

			for _, ans := range quest.Answers {

				qCode := ans.QCode
				value, found := submission.Data[qCode]

				if found && value != "" {
					text := getAnswerText(title, ans.Label, ans.Type, lookup)

					answer := &model.Answer{
						QCode: qCode,
						Text:  text,
						Value: value,
					}
					hasAnswerValue = true
					instance.Answers = append(instance.Answers, answer)
				}
			}
		}
		if hasAnswerValue {
			survey.Sections = append(survey.Sections, section)
		}
	}
	return survey
}

func fromSpp(schema *model.Schema, spp *model.Spp) *model.Survey {

	lookup := substitutions.GetLookup("start date", "end date", spp.Reference)

	survey := &model.Survey{
		Title:       schema.Title,
		SurveyId:    schema.SurveyId,
		FormType:    schema.FormType,
		Respondent:  spp.Reference,
		SubmittedAt: substitutions.DateFormat(spp.SubmittedAt),
		Sections:    []*model.Section{},
	}

	var sections []*model.Section
	for _, sect := range schema.Sections {
		section := &model.Section{
			Title:     sect.Title,
			Instances: []*model.Instance{},
		}
		sections = append(sections, section)
		instanceMap := map[string]*model.Instance{}
		for _, quest := range sect.Questions {

			title := quest.Title

			for _, ans := range quest.Answers {
				responseList := spp.GetResp(ans.QCode)
				for _, resp := range responseList {
					instance, found := instanceMap[resp.Instance]
					if !found {
						id, _ := strconv.Atoi(resp.Instance)
						instance = &model.Instance{
							Id:      id,
							Answers: []*model.Answer{},
						}
						section.Instances = append(section.Instances, instance)
					}

					value := resp.Response
					if value != "" {
						qCode := ans.QCode
						text := getAnswerText(title, ans.Label, ans.Type, lookup)

						answer := &model.Answer{
							QCode: qCode,
							Text:  text,
							Value: value,
						}

						instance.Answers = append(instance.Answers, answer)
					}
				}
			}
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
