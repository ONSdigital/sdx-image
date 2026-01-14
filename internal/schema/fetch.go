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

// SecretGetter defines what schema needs from a secret provider
type SecretGetter interface {
	Get(key string) (string, error)
}

type CIRClient struct {
	url      string
	audience string
	client   *http.Client
}

func NewClient(url, audience string) *CIRClient {
	return &CIRClient{url: url, audience: audience, client: &http.Client{}}
}

func (c *CIRClient) setAuthorisedClient() error {
	ctx := context.Background()
	client, err := idtoken.NewClient(ctx, c.audience)
	if err != nil {
		return fmt.Errorf("idtoken.NewClient: %w", err)
	}
	c.client = client
	return nil
}

func (c *CIRClient) fetchCirSchema(guid string) (*Schema, error) {
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

// Service uses dependency injection for secrets and client creation
type Service struct {
	Secrets       SecretGetter
	ClientFactory func(url, audience string) *CIRClient
}

func (s *Service) Fetch(guid string) (*Schema, error) {
	log.Info("Fetching schema for guid: " + guid + " from CIR")

	url, err := s.Secrets.Get("cir-url")
	if err != nil {
		return nil, fmt.Errorf("failed to get cir url from secret manager: %w", err)
	}

	audience, err := s.Secrets.Get("sdx-testdata-audience")
	if err != nil {
		return nil, fmt.Errorf("failed to get sdx-testdata-audience from secret manager: %w", err)
	}

	client := s.ClientFactory(url, audience)
	if err := client.setAuthorisedClient(); err != nil {
		return nil, fmt.Errorf("failed to set authorised client: %w", err)
	}

	schema, err := client.fetchCirSchema(guid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cir schema: %w", err)
	}
	return schema, nil
}
