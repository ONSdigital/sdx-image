package schema

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sdxImage/internal/test"
	"testing"
)

func loadCI() ([]byte, error) {
	test.SetCwdToRoot()
	jsonFile, err := os.Open("schemas/1_0005.json")
	defer func(jsonFile *os.File) {
		e := jsonFile.Close()
		if e != nil {
			fmt.Println("Failed to close schema file", e)
		}
	}(jsonFile)
	if err != nil {
		fmt.Println("Failed to open schema file", err)
		return nil, err
	}
	bytes, _ := io.ReadAll(jsonFile)
	return bytes, nil
}

func TestFetch(t *testing.T) {

	expectedGuid := "428ae4d1-8e7f-4a9d-8bef-05a266bf81e7"
	expectedEndpoint := "/v2/retrieve_collection_instrument"

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			t.Errorf("Failed: Invalid request. %v", err)
		}

		guid := query.Get("guid")

		test.Equal(t, expectedGuid, guid)
		test.Equal(t, expectedEndpoint, r.URL.Path)

		ci, err := loadCI()
		if err != nil {
			t.Errorf("failed to load test CI: %v", err)
		}
		_, err = w.Write(ci)
		if err != nil {
			t.Errorf("Failed to write response. %v", err)
		}
	}))

	defer svr.Close()

	err := os.Setenv("CIR_URL", svr.URL)
	if err != nil {
		t.Errorf("Failed to set environment variables")
	}
	err = os.Setenv("CIR_AUDIENCE", "bob.googleusercontent.com")
	if err != nil {
		t.Errorf("Failed to set environment variables")
	}

	schema, err := Fetch(expectedGuid)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	test.Equal(t, schema.Title, "Monthly Wages and Salaries Survey")
}
