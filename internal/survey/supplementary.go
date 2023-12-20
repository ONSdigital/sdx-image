package survey

import (
	"sdxImage/internal/interfaces"
)

type LocalUnit struct {
	unit    interfaces.LocalUnit
	answers []interfaces.Answer
}

func NewLocalUnit(unit interfaces.LocalUnit) *LocalUnit {
	answers := make([]interfaces.Answer, len(unit.GetResponses()))
	for i, response := range unit.GetResponses() {
		answers[i] = &Answer{
			Title:    "",
			QType:    "",
			QCode:    response.GetCode(),
			Label:    "",
			Value:    response.GetValue(),
			Multiple: false,
		}
	}
	return &LocalUnit{unit: unit, answers: answers}
}

func (lu *LocalUnit) GetName() string {
	return lu.unit.GetName()
}

func (lu *LocalUnit) GetAddress() string {
	return lu.unit.GetAddress()
}

func (lu *LocalUnit) GetAnswers() []interfaces.Answer {
	return lu.answers
}

func (lu *LocalUnit) AddAnswer() []interfaces.Answer {
	return lu.answers
}
