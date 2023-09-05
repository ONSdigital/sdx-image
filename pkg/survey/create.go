package survey

import (
	"sdxImage/pkg/interfaces"
	"sdxImage/pkg/substitutions"
	"strconv"
)

func Create(instrument interfaces.Schema, submission interfaces.Submission) interfaces.Survey {

	survey := &Survey{
		Title:       instrument.GetTitle(),
		SurveyId:    instrument.GetSurveyId(),
		FormType:    instrument.GetFormType(),
		Respondent:  submission.GetRuRef(),
		SubmittedAt: substitutions.DateFormat(submission.GetSubmittedAt()),
		Sections:    []interfaces.Section{},
	}

	var sections []interfaces.Section
	for _, sectionTitle := range instrument.GetSections().ListTitles() {
		hasAnswerValue := false
		section := &Section{
			Title:     sectionTitle,
			Instances: []interfaces.Instance{},
		}
		instanceMap := map[string]*Instance{}
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
						instance = &Instance{
							Id:      resp.GetInstance(),
							Answers: []interfaces.Answer{},
						}
						instanceMap[instanceKey] = instance
						section.Instances = append(section.Instances, instance)
					}

					value := resp.GetValue()
					if value != "" {

						answer := &Answer{
							Title:    title,
							QType:    answerType,
							QCode:    getQCode(answerQcode),
							Label:    answerLabel,
							Value:    value,
							Multiple: len(answers) > 1,
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

func getQCode(code string) string {
	_, err := strconv.Atoi(code)
	if err != nil {
		return getQCode(code[1:])
	}
	return code
}
