package schema

import (
	"fmt"
	"sdxImage/pkg/test"
	"testing"
)

func testRead(schema string, t *testing.T) {
	test.SetCwdToRoot()
	result, err := Read(schema)
	if err != nil {
		t.Errorf("failed to read %q with error: %q", schema, err.Error())
	}
	fmt.Println(result)
}

func TestReadQcas(t *testing.T) {
	testRead("qcas_0018", t)
}

func TestReadAbs(t *testing.T) {
	testRead("abs_1802", t)
}

func TestReadBerd(t *testing.T) {
	testRead("berd_0001", t)
}
