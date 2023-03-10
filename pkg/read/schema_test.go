package read

import (
	"fmt"
	"log"
	"os"
	"sdxImage/pkg/test"
	"testing"
)

func TestReadAbs(t *testing.T) {
	test.SetCwdToRoot()
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	result, err := Schema("abs_1802")
	if err != nil {
		t.Errorf("failed to read abs_1802 with error: %q", err.Error())
	}
	fmt.Println(result)
}

func TestReadMbs(t *testing.T) {
	test.SetCwdToRoot()
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	result, err := Schema("mbs_0106")
	if err != nil {
		t.Errorf("failed to read mbs_0106 with error: %q", err.Error())
	}
	fmt.Println(result)
}

func TestReadBerd(t *testing.T) {
	test.SetCwdToRoot()
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	result, err := Schema("berd_0001")
	if err != nil {
		t.Errorf("failed to read berd_0001 with error: %q", err.Error())
	}
	fmt.Println(result)
}
