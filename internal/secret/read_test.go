package secret

import (
	"fmt"
	"os"
	"testing"
)

func TestGetSecret(t *testing.T) {

	err := os.Setenv("GOOGLE_CLOUD_PROJECT", os.Getenv("PROJECT_ID"))
	if err != nil {
		return
	}
	// Not a unit test
	secret, err := Get("sdx-testdata-audience")
	if err != nil {
		return
	}
	fmt.Println(secret)
}
