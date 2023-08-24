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
	AnswerMap map[string]*Answer
}

func newAnswers() *Answers {
	return &Answers{
		idList:    []string{},
		AnswerMap: make(map[string]*Answer),
	}
}

func (answers *Answers) addAnswer(answer *Answer) {
	answers.idList = append(answers.idList, answer.Id)
	answers.AnswerMap[answer.Id] = answer
}

func (answers *Answers) ListIds() []string {
	return answers.idList
}

func (answers *Answers) GetType(answerId string) string {
	return answers.AnswerMap[answerId].AnswerType
}

func (answers *Answers) GetCode(answerId string) string {
	return answers.AnswerMap[answerId].Qcode
}

func (answers *Answers) GetLabel(answerId string) string {
	return answers.AnswerMap[answerId].Label
}
