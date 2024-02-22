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

		used := false
		for i, option := range answer.Options {
			code := option.Qcode
			if code == "" && !used {
				code = answer.Qcode
				used = true
			}
			result[i] = &AnswerSpec{
				AnswerType: answer.AnswerType,
				QCode:      code,
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
