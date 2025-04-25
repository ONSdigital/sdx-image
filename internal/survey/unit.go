package survey

import (
	s "sdxImage/internal/submission"
)

const ListName = "local-units"
const PpiItems = "item"

type Description struct {
	Key   string
	Value string
}

type Unit interface {
	GetIdentifier() string
	GetPrimaryDesc() Description
	GetSecondaryDesc() Description
	GetAnswers() []*Answer
	UpdateContext(code, displayCode, title, qType, label string) bool
}

type ExistingUnit struct {
	localUnit *s.LocalUnit
	answers   []*Answer
}

type NewUnit struct {
	answers []*Answer
}

type ExistingPpiItem struct {
	ppiItem *s.PpiItem
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

func NewExistingPpiItem(ppiItem *s.PpiItem, answer []*Answer) *ExistingPpiItem {
	return &ExistingPpiItem{ppiItem: ppiItem, answers: answer}
}

func GetExistingPpiItems(submission *s.Submission) []*ExistingPpiItem {
	var ppiItems []*ExistingPpiItem
	for _, listItemId := range submission.GetListItemIds(PpiItems) {
		ppiItem := submission.GetPpiItem(listItemId)
		if ppiItem != nil {
			var answers []*Answer
			for code, value := range submission.GetResponseForListId(listItemId) {
				answers = append(answers, NewAnswer(code, value))
			}
			ppiItems = append(ppiItems, NewExistingPpiItem(ppiItem, answers))
		}
	}

	return ppiItems
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

func (unit *ExistingUnit) GetPrimaryDesc() Description {
	return Description{
		Key:   "Name",
		Value: unit.localUnit.Name,
	}
}

func (unit *ExistingUnit) GetSecondaryDesc() Description {
	return Description{
		Key:   "Address",
		Value: unit.localUnit.GetAddress(),
	}
}

func (unit *ExistingUnit) GetAnswers() []*Answer {
	return unit.answers
}

func (unit *ExistingUnit) UpdateContext(code, displayCode, title, qType, label string) bool {
	updated := false
	for _, answer := range unit.answers {
		if answer.GetCode() == code {
			updated = true
			answer.SetContext(title, displayCode, qType, label, false)
		}
	}
	return updated
}

func (unit *ExistingPpiItem) GetIdentifier() string {
	return unit.ppiItem.Identifier
}

func (unit *ExistingPpiItem) GetPrimaryDesc() Description {
	return Description{
		Key:   "Item Number",
		Value: unit.ppiItem.ItemNumber,
	}
}

func (unit *ExistingPpiItem) GetSecondaryDesc() Description {
	return Description{
		Key:   "Item Specification",
		Value: unit.ppiItem.ItemSpecification,
	}
}

func (unit *ExistingPpiItem) GetAnswers() []*Answer {
	return unit.answers
}

func (unit *ExistingPpiItem) UpdateContext(code, displayCode, title, qType, label string) bool {
	updated := false
	for _, answer := range unit.answers {
		if answer.GetCode() == code {
			updated = true
			answer.SetContext(title, displayCode, qType, label, false)
		}
	}
	return updated
}

func (unit *NewUnit) GetIdentifier() string {
	return "New Local Unit"
}

func (unit *NewUnit) GetPrimaryDesc() Description {
	return Description{}
}

func (unit *NewUnit) GetSecondaryDesc() Description {
	return Description{}
}

func (unit *NewUnit) GetAnswers() []*Answer {
	return unit.answers
}

func (unit *NewUnit) UpdateContext(code, displayCode, title, qType, label string) bool {
	updated := false
	for _, answer := range unit.answers {
		if answer.GetCode() == code {
			updated = true
			answer.SetContext(title, displayCode, qType, label, false)
		}
	}
	return updated
}
