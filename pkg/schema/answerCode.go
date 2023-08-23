package schema

type AnswerCode struct {
	AnswerId string
	Code     string
}

func convertToAnswerCode(json map[string]any) (*AnswerCode, bool) {
	answerCode := &AnswerCode{
		AnswerId: getString(json, "answer_id"),
		Code:     getString(json, "code"),
	}
	return answerCode, true
}
