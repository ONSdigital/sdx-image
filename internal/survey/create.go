package survey

import (
	"sdxImage/internal/schema"
	s "sdxImage/internal/submission"
	"sdxImage/internal/substitutions"
)

const AdditionalSites = "additional_sites_name"
const BerdSurveyId = "002"
const BerdSectionTitle = "Workplace information"

// Create combines a schema and submission to return a Survey
// This is done by comparing the Submission responses to the Schema questions
// and matching on qcode.
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
		UnitType:    None,
	}

	if submission.HasLocalUnits() {
		survey.UnitType = LocalUnit
		for _, lu := range GetExistingUnits(submission) {
			survey.Units = append(survey.Units, lu)
		}

		for _, lu := range GetNewUnits(AdditionalSites, submission) {
			survey.Units = append(survey.Units, lu)
		}
	} else if submission.HasPpiItems() {
		// do ppi stuff
		survey.UnitType = PpiItem
		for _, ppiItem := range GetExistingPpiItems(submission) {
			survey.Units = append(survey.Units, ppiItem)
		}
	}

	responseMap := submission.GetResponses()
	for listItemId := range responseMap {
		name := submission.GetListItemName(listItemId)
		if name == AdditionalSites || name == ListName || name == PpiItems {
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
						// add question context to local localUnit
						// display code is how the qcode should be displayed e.g. c56 -> 56
						displayCode := getQCode(answerQcode, schema.GetSurveyId())
						unit.UpdateContext(answerQcode, displayCode, title, answerType, answerLabel)
						answer := unit.GetAnswer(answerQcode)
						if answer != nil {
							value := spec.GetValue(answer.Value)
							answer.Value = value
						}
					}

					respondentValue := data[answerQcode]
					value := spec.GetValue(respondentValue)
					if value != "" {

						//BERD hack for postcodes
						if schema.GetSurveyId() == BerdSurveyId {
							if section.Title == BerdSectionTitle {
								instance, instanceCount = berdSpecificInstance(instances, instanceCount, answerQcode)
							}
						}

						answer := &Answer{
							Title:    title,
							QType:    answerType,
							QCode:    getQCode(answerQcode, schema.GetSurveyId()),
							Label:    answerLabel,
							Value:    value,
							Multiple: spec.PartOfGroup(),
						}

						instance.Answers = append(instance.Answers, answer)
						_, exists := section.Instances[instance.Id]
						if !exists {
							section.Instances[instance.Id] = instance
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

func berdSpecificInstance(instances map[string]*Instance, instanceCount int, answerQcode string) (*Instance, int) {
	var instance *Instance
	instanceCount++
	var id string
	if len(answerQcode) > 3 {
		// all qcodes within this section without a letter are 3 characters long
		if len(answerQcode) > 4 {
			// some qcodes have a number and then a letter at the start.
			// If this is the case remove the number
			id = answerQcode[len(answerQcode)-4:]

		} else {
			id = answerQcode
		}
		// reduce the id to just the first letter. This will allow us to group them.
		id = id[0:1]
		instanceVal := instanceCount
		if id == "e" {
			// e identifies the qcode with the address for the current workplace
			// and so should be first in the section
			instanceVal = -1
		}
		ins, exists := instances[id]
		if exists {
			instance = ins
		} else {
			instance = &Instance{
				Id:      id,
				Value:   instanceVal,
				Answers: []*Answer{},
			}
		}
	} else {
		instance = &Instance{
			Id:      answerQcode,
			Value:   0,
			Answers: []*Answer{},
		}
	}
	return instance, instanceCount
}
