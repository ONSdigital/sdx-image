package schema

import (
	"sdxImage/internal/test"
	"testing"
	"time"
)

func fakeSchemaCreator(schemaName string) (*Schema, error) {
	return &Schema{
		Title:       "",
		SurveyId:    schemaName,
		FormType:    "",
		DataVersion: "",
		Sections:    nil,
		AnswerCodes: nil,
	}, nil
}

func fakeCiCreator(guid string) (*Schema, error) {
	return &Schema{
		Title:       "",
		SurveyId:    guid,
		FormType:    "",
		DataVersion: "",
		Sections:    nil,
		AnswerCodes: nil,
	}, nil
}

func TestCacheSchema(t *testing.T) {
	// Test the schema cache with schemas from file
	cache := NewCache(3, fakeSchemaCreator, fakeCiCreator)
	var s *Schema

	s, _ = cache.GetSchema("001", "")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "001", s.GetSurveyId())

	s, _ = cache.GetSchema("002", "")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "002", s.GetSurveyId())

	s, _ = cache.GetSchema("003", "")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "003", s.GetSurveyId())

	s, _ = cache.GetSchema("001", "")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "001", s.GetSurveyId())
	//ensure the cache hasn't grown because the schema already exists
	test.Equal(t, 3, len(cache.instruments))
	test.Equal(t, 3, len(cache.lastUsed))

	s, _ = cache.GetSchema("004", "")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "004", s.GetSurveyId())
	//ensure the cache hasn't grown
	test.Equal(t, 3, len(cache.instruments))
	test.Equal(t, 3, len(cache.lastUsed))
	//002 should have been removed as the oldest schema
	test.MapContains(t, cache.instruments, "001", "003", "004")
}

func TestCacheCir(t *testing.T) {
	// Test the schema cache with schemas from cir guids
	cache := NewCache(3, fakeSchemaCreator, fakeCiCreator)
	var s *Schema

	s, _ = cache.GetSchema("", "c7b0cfa8-147d-4345-8948-813dcbc1539f")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "c7b0cfa8-147d-4345-8948-813dcbc1539f", s.GetSurveyId())

	s, _ = cache.GetSchema("", "d2e1f4b2-1c4e-4f3a-9f7b-2e5d5c6a7b8c")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "d2e1f4b2-1c4e-4f3a-9f7b-2e5d5c6a7b8c", s.GetSurveyId())

	s, _ = cache.GetSchema("", "e3f2a5c4-2d5e-4f6b-8c9d-3f4e5d6c7b8a")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "e3f2a5c4-2d5e-4f6b-8c9d-3f4e5d6c7b8a", s.GetSurveyId())

	s, _ = cache.GetSchema("", "c7b0cfa8-147d-4345-8948-813dcbc1539f")
	time.Sleep(10 * time.Millisecond)
	test.Equal(t, "c7b0cfa8-147d-4345-8948-813dcbc1539f", s.GetSurveyId())
	//ensure the cache hasn't grown because the schema already exists
	test.Equal(t, 3, len(cache.instruments))
	test.Equal(t, 3, len(cache.lastUsed))
}
