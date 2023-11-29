package survey

import (
	"sdxImage/internal/interfaces"
	"sdxImage/internal/substitutions"
	"strconv"
)

func Create(schema interfaces.Schema, submission interfaces.Submission) interfaces.Survey {

	lookup := substitutions.GetLookup(submission.GetStartDate(), submission.GetEndDate(), submission.GetRuName(), submission.GetEmploymentDate())

	survey := &Survey{
		Title:       schema.GetTitle(),
		SurveyId:    schema.GetSurveyId(),
		FormType:    schema.GetFormType(),
		Respondent:  submission.GetRuRef(),
		RuName:      submission.GetRuName(),
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
			answerIds := schema.ListAnswerIds(questionId)

			for _, answerId := range answerIds {
				answerSpecs := schema.GetAnswers(answerId)

				for _, spec := range answerSpecs {
					answerQcode := spec.GetCode()
					answerLabel := spec.GetLabel()
					answerType := spec.GetType()

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
								QCode:    getQCode(answerQcode, submission.GetDataVersion(), schema.GetSurveyId()),
								Label:    substitutions.Replace(answerLabel, lookup),
								Value:    value,
								Multiple: len(answerIds) > 1 || len(answerSpecs) > 1,
							}

							instance.Answers = append(instance.Answers, answer)
							hasAnswerValue = true
						}
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

const LoopingDataVersion = "0.0.3"

func getQCode(code, dataVersion, surveyId string) string {
	if dataVersion == LoopingDataVersion {
		_, err := strconv.Atoi(code)
		if err != nil {
			return getQCode(code[1:], LoopingDataVersion, surveyId)
		}
	}

	if surveyId == "024" {
		return getFuelsCode(code)
	} else if surveyId == "194" {
		return getRailwaysCode(code)
	}

	return code
}

func getFuelsCode(code string) string {
	mapping := map[string]string{
		"10":  "0a",
		"11":  "0b",
		"12":  "0c",
		"13":  "0d",
		"14":  "0e",
		"110": "1",
		"120": "2a",
		"121": "2b",
		"122": "2c",
		"130": "3",
		"140": "4a",
		"141": "4b",
		"142": "4c",
		"150": "5",
		"160": "6",
		"180": "8",
		"190": "9",
		"200": "11",
		"210": "12",
		"211": "12a",
		"220": "13",
		"230": "15",
		"240": "16",
		"250": "18",
		"260": "19",
		"270": "20",
		"271": "20a",
		"280": "21",
		"290": "23",
		"300": "24",
		"310": "26",
		"320": "27",
		"330": "28",
		"340": "29",
		"350": "31",
		"360": "32",
		"370": "34",
		"146": "146",
		"12a": "17",
		"20a": "25",
		"28":  "33",
	}

	c, found := mapping[code]
	if found {
		return c
	}
	return code
}

func getRailwaysCode(code string) string {
	mapping := map[string]string{
		"2":   "1.1",
		"3":   "1.2",
		"4":   "2.1",
		"5":   "2.2",
		"6":   "3.1",
		"7":   "3.2",
		"8":   "3.3",
		"9":   "3.4",
		"10":  "4.1",
		"13":  "4.2",
		"146": "146",
	}

	c, found := mapping[code]
	if found {
		return c
	}
	return code
}
