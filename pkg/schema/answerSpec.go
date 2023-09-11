package schema

type AnswerSpec struct {
	AnswerType string
	QCode      string
	Label      string
}

func NewAnswerSpec(answerType, qCode, label string) *AnswerSpec {
	return &AnswerSpec{
		AnswerType: answerType,
		QCode:      qCode,
		Label:      label,
	}
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
