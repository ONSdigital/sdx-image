package schema

import (
	"fmt"
	"sdxImage/pkg/test"
	"testing"
)

func TestReadQcas(t *testing.T) {
	test.SetCwdToRoot()
	result, err := Read("qcas_0018")
	if err != nil {
		t.Errorf("failed to read qcas_0018 with error: %q", err.Error())
	}
	fmt.Println(result)
}
