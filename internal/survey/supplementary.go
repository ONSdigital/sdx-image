package survey

import (
	s "sdxImage/internal/submission"
)

type LocalUnit struct {
	unit    s.SupplementaryUnit
	answers []*Answer
}

func NewLocalUnit(unit s.SupplementaryUnit, submission *s.Submission) *LocalUnit {
	responses := unit.GetAssociatedResponses(submission)
	answers := make([]*Answer, len(responses))
	for i, response := range responses {
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

func (lu *LocalUnit) GetIdentifier() string {
	return lu.unit.Identifier
}

func (lu *LocalUnit) GetName() string {
	return lu.unit.Name
}

func (lu *LocalUnit) GetAddress() string {
	return lu.unit.GetAddress()
}

func (lu *LocalUnit) GetAnswers() []*Answer {
	return lu.answers
}

func (lu *LocalUnit) updateAnswer(code, title, qType, label string) bool {
	updated := false
	for i, answer := range lu.answers {
		if answer.GetCode() == code {
			updated = true
			lu.answers[i] = &Answer{
				Title:    title,
				QType:    qType,
				QCode:    answer.GetCode(),
				Label:    label,
				Value:    answer.GetValue(),
				Multiple: false,
			}
		}
	}
	return updated
}
