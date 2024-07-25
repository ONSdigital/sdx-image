package schema

import (
	"sdxImage/internal/test"
	"testing"
)

func getTestInstrument() *CollectionInstrument {
	return &CollectionInstrument{
		title:         "schema title",
		surveyId:      "surveyId",
		formType:      "formType",
		sectionTitles: []string{"t1", "t2"},
		titleToQidMap: map[string][]string{
			"t1": {"qId1", "qId2"},
			"t2": {"qId3"},
		},
		qidToQtMap: map[string]string{
			"qId1": "question title 1",
			"qId2": "question title 2",
			"qId3": "question title 3",
		},
		qidToAidMap: map[string][]string{
			"qId1": {"aId1"},
			"qId2": {"aId2"},
			"qId3": {"aId3", "aId4"},
		},
		answerMap: map[string][]*AnswerSpec{
			"aId1": {&AnswerSpec{"Number", "001", "Employee Count"}},
			"aId2": {&AnswerSpec{"Checkbox", "002", "Products"}},
			"aId3": {&AnswerSpec{"Radio", "003", "Gender"}},
			"aId4": {&AnswerSpec{"Currency", "004", "Total"}},
		},
	}
}

func TestInstrumentImmutability(t *testing.T) {

	instrument := getTestInstrument()

	titles := instrument.ListTitles()
	titles[0] = "mutated title!"
	test.Equal(t, "t1", instrument.ListTitles()[0])

	questionIds := instrument.ListQuestionIds("t1")
	questionIds[0] = "mutated question id!"
	test.Equal(t, "qId1", instrument.ListQuestionIds("t1")[0])

	answerIds := instrument.ListAnswerIds("qId1")
	answerIds[0] = "mutated answer id!"
	test.Equal(t, "aId1", instrument.ListAnswerIds("qId1")[0])

	answers := instrument.GetAnswers("aId1")
	answers[0] = &AnswerSpec{"Number", "001", "Mutated Label"}
	test.Equal(t, "Employee Count", instrument.GetAnswers("aId1")[0].GetLabel())
}

func TestMissingLookups(t *testing.T) {

	instrument := getTestInstrument()

	questionIds := instrument.ListQuestionIds("missing title")
	test.Equal(t, 0, len(questionIds))

	questionTitle := instrument.GetQuestionTitle("Missing Qid")
	test.Equal(t, "", questionTitle)

	answerIds := instrument.ListAnswerIds("missing Qid")
	test.Equal(t, 0, len(answerIds))

	answers := instrument.GetAnswers("missing answerId")
	test.Equal(t, 0, len(answers))
}
