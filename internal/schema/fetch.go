package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sdxImage/internal/log"
	"sdxImage/internal/secret"

	"google.golang.org/api/idtoken"
)

const CirResourcePath = "/v2/retrieve_collection_instrument"
const CirUrlSecret = "cir-url"
const CirAudienceSecret = "sdx-testdata-audience"

// CirClient handles communication with the CIR service
type CirClient struct {
	url    string
	client *http.Client
}

// NewClient creates a new CirClient instance with authorised HTTP client
func NewClient() *CirClient {

	// Retrieve CIR URL and audience from environment or secret manager
	secretMgr := secret.NewManager()
	url, err := secretMgr.Get(CirUrlSecret)

	if err != nil {
		fmt.Println("Error retrieving CIR URL from secrets:", err)
		return nil
	}
	log.Info("DEBUG URL: " + url)

	audience, err := secretMgr.Get(CirAudienceSecret)
	if err != nil {
		fmt.Println("Error retrieving CIR audience from secrets:", err)
		return nil
	}
	log.Info("DEBUG audience: " + audience)

	// Set up authorised client
	ctx := context.Background()
	client, err := idtoken.NewClient(ctx, audience)
	if err != nil {
		fmt.Println("Error creating authorised client:", err)
		return nil
	}
	return &CirClient{url: url, client: client}
}

// fetchCirSchema retrieves the schema from CIR by guid and returns it
func (c *CirClient) fetchCirSchema(guid string) (*Schema, error) {

	log.Info("DEBUG -> URL_STRING: " + c.url + CirResourcePath + "?guid=" + guid)
	log.Info("fetch stage 1")

	resp, err := c.client.Get(c.url + CirResourcePath + "?guid=" + guid)
	if err != nil {
		return nil, err
	}
	log.Info("fetch stage 2")

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Info("fetch stage 3")

	var schema Schema
	if err := json.Unmarshal(body, &schema); err != nil {
		return nil, err
	}

	log.Info("fetch stage 4")

	return &schema, nil
}

// Fetch is the main method to get schema from CIR by guid, requires a Service to be set up
func (c *CirClient) Fetch(guid string) (*Schema, error) {
	log.Info("Fetching schema for guid: " + guid + " from CIR")

	schema, err := c.fetchCirSchema(guid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cir schema: %w", err)
	}
	log.Info("Successfully fetched schema for guid: " + guid)
	return schema, nil
}
