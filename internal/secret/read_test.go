package secret

import (
	"fmt"
	"testing"
)

func TestGetSecret(t *testing.T) {

	// Not a unit test
	secret, err := Get("iap-secret")
	if err != nil {
		return
	}
	fmt.Println(secret)
}
