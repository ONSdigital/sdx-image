package secret

import (
	"context"
	"fmt"
	"os"
	"sdxImage/internal/log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func Get(secretId string) (string, error) {
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT")

	// Access the latest version (alias "latest").
	secretName := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectId, secretId)

	// Log the secret name being accessed
	log.Info("Accessing secret: " + secretName)

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Error("secretmanager.NewClient", err)
	}
	defer client.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Error("AccessSecretVersion", err)
	}

	// Secret payload is bytes. Convert to string if it's UTF-8 text.
	secret := string(result.Payload.Data)
	fmt.Println("Secret value:", secret)

	return secret, nil
}
