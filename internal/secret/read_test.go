package secret

import (
	"fmt"
	"testing"
)

func TestGetSecret(t *testing.T) {

	// Not a unit test
	secret, err := Get("sdx-testdata-audience")
	if err != nil {
		return
	}
	fmt.Println(secret)
}
