package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sdxImage/internal/log"

	"google.golang.org/api/idtoken"
)

const CirResourcePath = "/v2/retrieve_collection_instrument"

// SecretGetter defines the interface for retrieving secrets (allows easy mocking)
type SecretGetter interface {
	Get(key string) (string, error)
}

// CirClient handles communication with the CIR service
type CirClient struct {
	url      string
	audience string
	client   *http.Client
}

// NewClient creates a new CirClient instance
func NewClient(url, audience string) *CirClient {
	c := &CirClient{url: url, audience: audience, client: &http.Client{}}
	err := c.setAuthorisedClient()
	if err != nil {
		return nil
	}
	return c
}

// setAuthorisedClient generates a bearer token on the CIR client
func (c *CirClient) setAuthorisedClient() error {
	ctx := context.Background()
	client, err := idtoken.NewClient(ctx, c.audience)
	if err != nil {
		return fmt.Errorf("idtoken.NewClient: %w", err)
	}
	c.client = client
	return nil
}

// fetchCirSchema retrieves the schema from CIR by guid and returns it
func (c *CirClient) fetchCirSchema(guid string) (*Schema, error) {
	resp, err := c.client.Get(c.url + CirResourcePath + "?guid=" + guid)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var schema Schema
	if err := json.Unmarshal(body, &schema); err != nil {
		return nil, err
	}
	return &schema, nil
}

// Service to interact with CIR, needs secrets and a client factory
type Service struct {
	url       string
	audience  string
	CirClient *CirClient
}

// Fetch is the main method to get schema from CIR by guid, requires a Service to be set up
func (s *Service) Fetch(guid string) (*Schema, error) {
	log.Info("Fetching schema for guid: " + guid + " from CIR")
	schema, err := s.CirClient.fetchCirSchema(guid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cir schema: %w", err)
	}
	return schema, nil
}
