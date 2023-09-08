package survey

import (
	"sdxImage/pkg/interfaces"
	"sdxImage/pkg/substitutions"
	"strconv"
)

func Create(schema interfaces.Schema, submission interfaces.Submission) interfaces.Survey {

	lookup := substitutions.GetLookup(submission.GetStartDate(), submission.GetEndDate(), submission.GetRuName())

	survey := &Survey{
		Title:       schema.GetTitle(),
		SurveyId:    schema.GetSurveyId(),
		FormType:    schema.GetFormType(),
		Respondent:  submission.GetRuRef(),
		SubmittedAt: substitutions.DateFormat(submission.GetSubmittedAt()),
		Sections:    []interfaces.Section{},
	}

	var sections []interfaces.Section
	for _, sectionTitle := range schema.ListTitles() {
		hasAnswerValue := false
		section := &Section{
			Title:     sectionTitle,
			Instances: []interfaces.Instance{},
		}
		instanceMap := map[string]*Instance{}
		questions := schema.ListQuestionIds(sectionTitle)

		for _, questionId := range questions {
			title := schema.GetQuestionTitle(questionId)
			answers := schema.ListAnswers(questionId)

			for _, answerId := range answers {
				answerQcode := schema.GetAnswerCode(answerId)
				answerLabel := schema.GetAnswerLabel(answerId)
				answerType := schema.GetAnswerType(answerId)

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
							Title:    substitutions.Replace(title, lookup),
							QType:    answerType,
							QCode:    getQCode(answerQcode, submission.GetDataVersion()),
							Label:    substitutions.Replace(answerLabel, lookup),
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

func getQCode(code, dataVersion string) string {
	if dataVersion == "0.0.3" {
		_, err := strconv.Atoi(code)
		if err != nil {
			return getQCode(code[1:], "0.0.3")
		}
	}
	return code
}
