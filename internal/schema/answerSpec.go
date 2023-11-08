package schema

import "sdxImage/internal/interfaces"

type AnswerSpec struct {
	AnswerType string
	QCode      string
	Label      string
}

func getAnswerSpecs(answer *Answer) []interfaces.AnswerSpec {
	if answer.AnswerType == "Checkbox" {
		result := make([]interfaces.AnswerSpec, len(answer.Options))
		for i, option := range answer.Options {
			result[i] = &AnswerSpec{
				AnswerType: answer.AnswerType,
				QCode:      option.Qcode,
				Label:      option.Label,
			}
		}
		return result
	}
	return []interfaces.AnswerSpec{&AnswerSpec{
		AnswerType: answer.AnswerType,
		QCode:      answer.Qcode,
		Label:      string(answer.Label),
	}}
}

func (a *AnswerSpec) GetType() string {
	return a.AnswerType
}

func (a *AnswerSpec) GetCode() string {
	return a.QCode
}

func (a *AnswerSpec) GetLabel() string {
	return a.Label
}
