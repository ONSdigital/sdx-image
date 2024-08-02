package schema

import (
	"sdxImage/internal/test"
	"testing"
	"time"
)

func fakeSchemaCreator(guid string) (*Schema, error) {
	return &Schema{
		Title:       "",
		SurveyId:    guid,
		FormType:    "",
		DataVersion: "",
		Sections:    nil,
		AnswerCodes: nil,
	}, nil
}

func TestCache(t *testing.T) {
	cache := NewCache(3, fakeSchemaCreator)
	var s *Schema

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
