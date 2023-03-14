package model

import "strconv"

func FromSubmission(schema *Schema, submission *Submission) *Survey {
	survey := &Survey{
		Title:       schema.Title,
		SurveyId:    schema.SurveyId,
		FormType:    schema.FormType,
		Respondent:  submission.RuRef,
		SubmittedAt: submission.SubmittedAt,
		Sections:    []*Section{},
	}
	for _, sect := range schema.Sections {
		hasAnswerValue := false
		instance := &Instance{
			Id:      0,
			Answers: []*Answer{},
		}
		section := &Section{
			Title:     sect.Title,
			Instances: []*Instance{instance},
		}

		for _, quest := range sect.Questions {

			title := quest.Title

			for _, ans := range quest.Answers {

				qCode := ans.QCode
				value, found := submission.Data[qCode]

				if found && value != "" {
					text := title
					qType := ans.Type
					label := ans.Label
					if qType == "Date" {
						text += " " + label
					} else if qType == "Number" {
						text += " " + label + ":"
					} else if qType == "Currency" {
						text = label + "?"
					}

					answer := &Answer{
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

func FromSpp(schema *Schema, spp *Spp) *Survey {
	survey := &Survey{
		Title:       schema.Title,
		SurveyId:    schema.SurveyId,
		FormType:    schema.FormType,
		Respondent:  spp.Reference,
		SubmittedAt: spp.SubmittedAt,
		Sections:    []*Section{},
	}

	var sections []*Section
	for _, sect := range schema.Sections {
		section := &Section{
			Title:     sect.Title,
			Instances: []*Instance{},
		}
		sections = append(sections, section)
		instanceMap := map[string]*Instance{}
		for _, quest := range sect.Questions {

			title := quest.Title

			for _, ans := range quest.Answers {
				responseList := spp.getResp(ans.QCode)
				for _, resp := range responseList {
					instance, found := instanceMap[resp.Instance]
					if !found {
						id, _ := strconv.Atoi(resp.Instance)
						instance = &Instance{
							Id:      id,
							Answers: []*Answer{},
						}
						section.Instances = append(section.Instances, instance)
					}

					value := resp.Response
					if value != "" {
						qCode := ans.QCode
						text := title
						qType := ans.Type
						label := ans.Label
						if qType == "Date" {
							text += " " + label
						} else if qType == "Number" {
							text += " " + label + ":"
						} else if qType == "Currency" {
							text = label + "?"
						}

						answer := &Answer{
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
	return survey
}
