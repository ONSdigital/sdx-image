package survey

import (
	s "sdxImage/internal/submission"
)

const ListName = "local-units"

type Unit interface {
	GetIdentifier() string
	GetName() string
	GetAddress() string
	GetAnswers() []*Answer
	UpdateContext(code, title, qType, label string) bool
}

type ExistingUnit struct {
	localUnit *s.LocalUnit
	answers   []*Answer
}

type NewUnit struct {
	answers []*Answer
}

func NewExistingUnit(localUnit *s.LocalUnit, answer []*Answer) *ExistingUnit {
	return &ExistingUnit{localUnit: localUnit, answers: answer}
}

func GetExistingUnits(submission *s.Submission) []*ExistingUnit {
	var units []*ExistingUnit
	for _, listItemId := range submission.GetListItemIds(ListName) {
		localUnit := submission.GetLocalUnit(listItemId)
		if localUnit != nil {
			var answers []*Answer
			for code, value := range submission.GetResponseForListId(listItemId) {
				answers = append(answers, NewAnswer(code, value))
			}
			units = append(units, NewExistingUnit(localUnit, answers))
		}
	}

	return units
}

func GetNewUnits(listName string, submission *s.Submission) []*NewUnit {
	var units []*NewUnit
	for _, listItemId := range submission.GetListItemIds(listName) {
		var answers []*Answer
		for code, value := range submission.GetResponseForListId(listItemId) {
			answers = append(answers, NewAnswer(code, value))
		}
		units = append(units, &NewUnit{answers: answers})
	}

	return units
}

func (unit *ExistingUnit) GetIdentifier() string {
	return unit.localUnit.Identifier
}

func (unit *ExistingUnit) GetName() string {
	return unit.localUnit.Name
}

func (unit *ExistingUnit) GetAddress() string {
	return unit.localUnit.GetAddress()
}

func (unit *ExistingUnit) GetAnswers() []*Answer {
	return unit.answers
}

func (unit *ExistingUnit) UpdateContext(code, title, qType, label string) bool {
	updated := false
	for _, answer := range unit.answers {
		if answer.GetCode() == code {
			updated = true
			answer.SetContext(title, qType, label, false)
		}
	}
	return updated
}

func (unit *NewUnit) GetIdentifier() string {
	return "New Local Unit"
}

func (unit *NewUnit) GetName() string {
	return ""
}

func (unit *NewUnit) GetAddress() string {
	return ""
}

func (unit *NewUnit) GetAnswers() []*Answer {
	return unit.answers
}

func (unit *NewUnit) UpdateContext(code, title, qType, label string) bool {
	updated := false
	for _, answer := range unit.answers {
		if answer.GetCode() == code {
			updated = true
			answer.SetContext(title, qType, label, false)
		}
	}
	return updated
}
