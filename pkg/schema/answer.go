package schema

type Answer struct {
	Id         string
	Qcode      string
	AnswerType string
	Label      string
}

func convertToAnswer(json map[string]any) (*Answer, bool) {
	answer := &Answer{
		Id:         getString(json, "id"),
		Qcode:      getString(json, "q_code"),
		AnswerType: getString(json, "type"),
		Label:      getString(json, "label"),
	}
	return answer, true
}

type Answers struct {
	idList    []string
	answerMap map[string]*Answer
}

func newAnswers() *Answers {
	return &Answers{
		idList:    []string{},
		answerMap: make(map[string]*Answer),
	}
}

func (answers *Answers) addAnswer(answer *Answer) {
	answers.idList = append(answers.idList, answer.Id)
	answers.answerMap[answer.Id] = answer
}

func (answers *Answers) GetAnswerIds() []string {
	return answers.idList
}

func (answers *Answers) GetAnswerType(answerId string) string {
	return answers.answerMap[answerId].AnswerType
}

func (answers *Answers) GetAnswerCode(answerId string) string {
	return answers.answerMap[answerId].Qcode
}

func (answers *Answers) GetAnswerLabel(answerId string) string {
	return answers.answerMap[answerId].Label
}
