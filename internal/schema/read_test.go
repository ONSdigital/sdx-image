package schema

import (
	"fmt"
	"sdxImage/internal/test"
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

func TestReadTiles(t *testing.T) {
	testRead("qrt_0001", t)
}

func TestReadPrices(t *testing.T) {
	testRead("ppi_0001", t)
}

func TestReadIPI(t *testing.T) {
	testRead("ipi_0001", t)
	testRead("ipi_0002", t)
}

func TestReadEPI(t *testing.T) {
	testRead("epi_0001", t)
	testRead("epi_0002", t)
}
