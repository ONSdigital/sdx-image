package schema

import (
	"sdxImage/internal/test"
	"testing"
	"time"
)

func fakeSchemaCreator(guid string) (*CollectionInstrument, error) {
	return &CollectionInstrument{
		title:         "",
		surveyId:      guid,
		formType:      "",
		sectionTitles: nil,
		titleToQidMap: nil,
		qidToQtMap:    nil,
		qidToAidMap:   nil,
		answerMap:     nil,
	}, nil
}

func TestCache(t *testing.T) {
	cache := NewCache(3, fakeSchemaCreator)
	var s *CollectionInstrument

	s, _ = cache.GetSchema("001")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "001", s.GetSurveyId())

	s, _ = cache.GetSchema("002")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "002", s.GetSurveyId())

	s, _ = cache.GetSchema("003")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "003", s.GetSurveyId())

	s, _ = cache.GetSchema("001")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "001", s.GetSurveyId())
	//ensure the cache hasn't grown because the schema already exists
	test.Equal(t, 3, len(cache.instruments))
	test.Equal(t, 3, len(cache.lastUsed))

	s, _ = cache.GetSchema("004")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "004", s.GetSurveyId())
	//ensure the cache hasn't grown
	test.Equal(t, 3, len(cache.instruments))
	test.Equal(t, 3, len(cache.lastUsed))
	//002 should have been removed as the oldest schema
	test.MapContains(t, cache.instruments, "001", "003", "004")
}

func TestCacheIntegration(t *testing.T) {
	test.SetCwdToRoot()
	cache := NewCache(20, CreateInstrument)
	schemaName := "test_1802"
	schema, err := cache.GetSchema(schemaName)
	if err != nil {
		t.Errorf("failed to get schema: %q with error: %q", schemaName, err.Error())
	}

	titles := schema.ListTitles()
	test.Equal(t, "Income", titles[2])

	schema, err = cache.GetSchema(schemaName)
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
