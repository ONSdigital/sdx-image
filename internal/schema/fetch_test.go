package schema

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sdxImage/internal/test"
	"testing"

	"google.golang.org/api/idtoken"
)

// Helper to load test schema JSON
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

// FakeSecretGetter implements SecretGetter for testing
type FakeSecretGetter struct {
	URL string
}

func FakeClient(url, audience string) *CirClient {
	return &CirClient{url: url, audience: audience, client: &http.Client{}}
}

// Get returns fake secrets based on key
func (f *FakeSecretGetter) Get(key string) (string, error) {
	switch key {
	case "cir-url":
		return f.URL, nil
	case "sdx-testdata-audience":
		return "fake-audience", nil
	default:
		return "", fmt.Errorf("unknown key: %s", key)
	}
}

func TestFetch(t *testing.T) {
	expectedGuid := "428ae4d1-8e7f-4a9d-8bef-05a266bf81e7"
	expectedEndpoint := "/v2/retrieve_collection_instrument"

	// Start a test HTTP server to simulate CIR
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			t.Errorf("Failed: Invalid request. %v", err)
		}

		guid := query.Get("guid")
		test.Equal(t, expectedGuid, guid)
		test.Equal(t, expectedEndpoint, r.URL.Path)

		// Load the example schema JSON from file
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

	// Create service with fake secrets and client factory
	svc := &Service{
		url:       svr.URL,
		audience:  "fake-audience",
		CirClient: FakeClient(svr.URL, svr.URL),
	}

	// Act
	schema, err := svc.Fetch(expectedGuid)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	// Assert matches 1_0005.json title
	test.Equal(t, schema.Title, "Monthly Wages and Salaries Survey")
}

func TestReal(t *testing.T) {

	fakeAudience := "abc-fake.apps.googleusercontent.com"

	ctx := context.Background()
	ts, err := idtoken.NewTokenSource(ctx, fakeAudience)
	if err != nil {
		t.Fatalf("Failed to create token source: %v", err)
	}

	// Get the token
	token, err := ts.Token()
	if err != nil {
		t.Fatalf("Failed to get token: %v", err)
	}

	// Print the ID token
	fmt.Println("ID Token:", token.AccessToken)

}
