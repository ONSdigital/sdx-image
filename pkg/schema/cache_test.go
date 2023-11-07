package schema

import (
	"sdxImage/pkg/test"
	"testing"
)

func TestCacheIntegration(t *testing.T) {
	test.SetCwdToRoot()
	schemaCache := NewCache()
	schemaName := "test_1802"
	schema, err := schemaCache.GetSchema(schemaName)
	if err != nil {
		t.Errorf("failed to get schema: %q with error: %q", schemaName, err.Error())
	}

	titles := schema.ListTitles()
	test.Equal(t, "Income", titles[2])

	schema, err = schemaCache.GetSchema(schemaName)
	if err != nil {
		t.Errorf("failed to get schema: %q with error: %q", schemaName, err.Error())
	}

	qIds := schema.ListQuestionIds("Income")
	test.Equal(t, "question5e9943ec-5896-48dd-8427-12c14d80baca", qIds[0])

	aIds := schema.ListAnswerIds(qIds[0])
	test.Equal(t, "answer06b7045b-f9cb-4a36-8463-5ed4a74f5a67", aIds[0])

	answers := schema.GetAnswers(aIds[0])
	expected := answers[0]
	test.Equal(t, "Currency", expected.GetType())
	test.Equal(t, "399", expected.GetCode())
	test.Equal(t, "Total turnover", expected.GetLabel())
}
