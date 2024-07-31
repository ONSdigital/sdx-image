package survey

import (
	"sdxImage/internal/schema"
	s "sdxImage/internal/submission"
	"sdxImage/internal/substitutions"
)

const AdditionalSites = "additional_sites_name"

func Create(schema *schema.Schema, submission *s.Submission) *Survey {

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

	for _, schemaSection := range schema.GetSections() {
		hasAnswerValue := false
		preExistingSection := false

		var section *Section
		//check if the section already exists
		for _, sect := range sections {
			title := substitutions.Replace(schemaSection.GetTitle(), lookup)
			if sect.Title == title {
				hasAnswerValue = true
				preExistingSection = true
				section = sect
				break
			}
		}

		//if the section does not exist, create a new one
		if !preExistingSection {
			section = &Section{
				Title:     substitutions.Replace(schemaSection.GetTitle(), lookup),
				Instances: map[string]*Instance{},
			}
		}
		questions := schemaSection.GetQuestions()

		instances := section.Instances

		instanceCount := 0
		instanceVal := 0
		for listItemId, data := range responseMap {
			// check if the instance already exists
			instance, found := instances[listItemId]
			if !found {
				if listItemId == s.NonListItem {
					instanceVal = 0
				} else {
					instanceCount++
					instanceVal = instanceCount
				}

				instance = &Instance{
					Id:      listItemId,
					Value:   instanceVal,
					Answers: []*Answer{},
				}
			}

			for _, q := range questions {
				title := substitutions.Replace(q.GetTitle(), lookup)

				for _, spec := range schema.GetAnswerSpecs(q) {
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
							Multiple: spec.PartOfGroup(),
						}

						instance.Answers = append(instance.Answers, answer)
						_, exists := section.Instances[listItemId]
						if !exists {
							section.Instances[listItemId] = instance
						}
						hasAnswerValue = true
					}
				}
			}
		}
		if hasAnswerValue {
			if !preExistingSection {
				sections = append(sections, section)
			}
		}
	}
	survey.Sections = sections
	return survey
}
