package schema

const LoopingDataVersion = "0.0.3"

type AnswerSpec struct {
	AnswerType string
	QCode      string
	Label      string
	labelMap   map[string]string
	multiple   bool
}

func getAnswerSpecs(answer *Answer, schema *Schema, multiple bool) []*AnswerSpec {
	labelMap := make(map[string]string)

	if answer.AnswerType == "Checkbox" {
		result := make([]*AnswerSpec, len(answer.Options))

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
				multiple:   true,
			}
		}
		return result
	}

	if answer.AnswerType == "Radio" {
		for _, option := range answer.Options {
			labelMap[option.Value] = option.Label
		}
	}

	qCode := answer.Qcode
	if schema.DataVersion == LoopingDataVersion {
		qCode = lookupQCode(answer.Id, schema)
	}

	return []*AnswerSpec{{
		AnswerType: answer.AnswerType,
		QCode:      qCode,
		Label:      string(answer.Label),
		labelMap:   labelMap,
		multiple:   multiple,
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

func (a *AnswerSpec) PartOfGroup() bool {
	return a.multiple
}

func (a *AnswerSpec) GetValue(respondentValue string) string {
	label, found := a.labelMap[respondentValue]
	if found {
		return label
	}
	return respondentValue
}

func lookupQCode(answerId string, schema *Schema) string {
	for _, answerCode := range schema.AnswerCodes {
		if answerCode.AnswerId == answerId {
			return answerCode.Code
		}
	}
	return ""
}
