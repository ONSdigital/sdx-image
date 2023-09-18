package survey

import "sdxImage/pkg/interfaces"

type Instance struct {
	Id      int
	Answers []interfaces.Answer
}

func (instance *Instance) GetId() int {
	return instance.Id
}

func (instance *Instance) GetAnswers() []interfaces.Answer {
	return instance.Answers
}
