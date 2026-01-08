package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"google.golang.org/api/idtoken"
)

const CirCollectionInstrumentResourcePath = "/v2/retrieve_collection_instrument"

type Client struct {
	url      string
	audience string // The Audience (Client ID) of the IAP protected resource
	client   *http.Client
}

// NewClient Create a new client to fetch from IAP protected resource
func NewClient(url string, audience string) Client {
	return Client{url, audience, &http.Client{}}
}

// setAuthorisedClient Add the needed headers for IAP authentication
func (c *Client) setAuthorisedClient() error {
	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, c.audience)
	if err != nil {
		return fmt.Errorf("idtoken.NewClient: %w", err)
	}
	c.client = client
	return nil
}

// fetchCirSchema from the IAP protected resource using the provided guid
func (c *Client) fetchCirSchema(guid string) (*Schema, error) {
	resp, err := c.client.Get(c.url + CirCollectionInstrumentResourcePath + "?guid=" + guid)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var schema Schema
	err = json.Unmarshal(body, &schema)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}

// Fetch a schema from the IAP protected resource using the provided guid
func Fetch(guid string) (*Schema, error) {
	url := os.Getenv("CIR_URL")
	audience := os.Getenv("CIR_AUDIENCE")
	client := NewClient(url, audience)
	schema, err := client.fetchCirSchema(guid)
	if err != nil {
		err = fmt.Errorf("failed to fetch cir schema: %w", err)
		return nil, err
	}
	return schema, nil
}
