package controller

import (
	"sdxImage/pkg/interfaces"
	"sdxImage/pkg/model"
	"sdxImage/pkg/substitutions"
	"strconv"
	"strings"
)

func fromSubmission(instrument interfaces.Schema, submission interfaces.Submission) *model.Survey {

	lookup := substitutions.GetLookup(submission.GetStartDate(), submission.GetEndDate(), submission.GetRuName())

	survey := &model.Survey{
		Title:       instrument.GetTitle(),
		SurveyId:    instrument.GetSurveyId(),
		FormType:    instrument.GetFormType(),
		Respondent:  submission.GetRuRef(),
		SubmittedAt: substitutions.DateFormat(submission.GetSubmittedAt()),
		Sections:    []*model.Section{},
	}

	var sections []*model.Section
	for _, sectionTitle := range instrument.GetSections().ListTitles() {
		hasAnswerValue := false
		section := &model.Section{
			Title:     sectionTitle,
			Instances: []*model.Instance{},
		}
		instanceMap := map[string]*model.Instance{}
		questions := instrument.GetSections().ListQuestions(sectionTitle)

		for _, questionId := range questions {
			title := instrument.GetQuestions().GetTitle(questionId)
			answers := instrument.GetQuestions().ListAnswers(questionId)

			for _, answerId := range answers {
				answerQcode := instrument.GetAnswers().GetCode(answerId)
				answerLabel := instrument.GetAnswers().GetLabel(answerId)
				answerType := instrument.GetAnswers().GetType(answerId)

				responseList := submission.GetResponses(answerQcode)
				for _, resp := range responseList {

					instanceKey := strconv.Itoa(resp.GetInstance())
					instance, found := instanceMap[instanceKey]
					if !found {
						instance = &model.Instance{
							Id:      resp.GetInstance(),
							Answers: []*model.Answer{},
						}
						instanceMap[instanceKey] = instance
						section.Instances = append(section.Instances, instance)
					}

					value := resp.GetValue()
					if value != "" {
						qCode := getQCode(answerQcode)
						text := getAnswerText(title, answerLabel, answerType, len(answers) > 1, lookup)

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
