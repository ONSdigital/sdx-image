package survey

import (
	"sdxImage/internal/schema"
	s "sdxImage/internal/submission"
	"sdxImage/internal/substitutions"
	"strconv"
)

const AdditionalSites = "additional_sites_name"

func Create(schema *schema.CollectionInstrument, submission *s.Submission) *Survey {

	lookup := substitutions.GetLookup(
		submission.GetStartDate(),
		submission.GetEndDate(),
		submission.GetRuName(),
		submission.GetEmploymentDate())

	survey := &Survey{
		Title:       schema.GetTitle(),
		SurveyId:    schema.GetSurveyId(),
		FormType:    schema.GetFormType(),
		Respondent:  submission.GetRuRef(),
		RuName:      submission.GetRuName(),
		SubmittedAt: substitutions.DateFormat(submission.GetSubmittedAt()),
		Sections:    []*Section{},
		Units:       []Unit{},
	}

	for _, lu := range GetExistingUnits(submission) {
		survey.Units = append(survey.Units, lu)
	}

	for _, lu := range GetNewUnits(AdditionalSites, submission) {
		survey.Units = append(survey.Units, lu)
	}

	responseMap := submission.GetResponses()
	for listItemId := range responseMap {
		name := submission.GetListItemName(listItemId)
		if name == AdditionalSites || name == ListName {
			delete(responseMap, listItemId)
		}
	}

	var sections []*Section

	for _, sectionTitle := range schema.ListTitles() {
		hasAnswerValue := false
		section := &Section{
			Title:     substitutions.Replace(sectionTitle, lookup),
			Instances: map[string]*Instance{},
		}
		questions := schema.ListQuestionIds(sectionTitle)

		instanceCount := 0
		instanceId := 0
		for listItemId, data := range responseMap {
			if listItemId == s.NonListItem {
				instanceId = 0
			} else {
				instanceId = instanceCount
				instanceCount++
			}

			instance := &Instance{
				Id:      instanceId,
				Answers: []*Answer{},
			}

			for _, questionId := range questions {
				title := substitutions.Replace(schema.GetQuestionTitle(questionId), lookup)
				answerIds := schema.ListAnswerIds(questionId)

				for _, answerId := range answerIds {
					answerSpecs := schema.GetAnswers(answerId)

					for _, spec := range answerSpecs {
						answerQcode := spec.GetCode()
						answerLabel := substitutions.Replace(spec.GetLabel(), lookup)
						answerType := spec.GetType()

						for _, unit := range survey.Units {
							//add question context to local localUnit
							unit.UpdateContext(answerQcode, title, answerType, answerLabel)
						}

						value := data[answerQcode]
						if value != "" {

							answer := &Answer{
								Title:    title,
								QType:    answerType,
								QCode:    getQCode(answerQcode, schema.GetSurveyId()),
								Label:    answerLabel,
								Value:    value,
								Multiple: len(answerIds) > 1 || len(answerSpecs) > 1,
							}

							instance.Answers = append(instance.Answers, answer)
							id := strconv.Itoa(instanceId)
							_, exists := section.Instances[id]
							if !exists {
								section.Instances[id] = instance
							}
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
