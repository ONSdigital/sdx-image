package secret

import (
	"context"
	"fmt"
	"os"
	"sdxImage/internal/log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type Manager struct {
	ProjectID string
}

func NewSecretManager() *Manager {
	return &Manager{
		ProjectID: os.Getenv("PROJECT_ID"),
	}
}

func (sm *Manager) Get(key string) (string, error) {
	if sm.ProjectID == "" {
		return "", fmt.Errorf("PROJECT_ID environment variable is not set")
	}

	secretName := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", sm.ProjectID, key)
	log.Info("Accessing secret: " + secretName)

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %w", err)
	}
	defer client.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{Name: secretName}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %w", err)
	}

	return string(result.Payload.Data), nil
}
