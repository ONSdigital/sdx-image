package submission

import (
	"fmt"
	"io"
	"os"
	"sdxImage/internal/test"
	"testing"
)

func readFile(filename string) ([]byte, error) {
	jsonFile, err := os.Open("examples/submissions/" + filename + ".json")
	if err != nil {
		return nil, err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)
	bytes, _ := io.ReadAll(jsonFile)
	return bytes, nil
}

func TestAbsSubmission(t *testing.T) {
	//Data version  0.0.1
	test.SetCwdToRoot()
	filename := "abs_1802"
	fileBytes, e := readFile(filename)
	if e != nil {
		t.Errorf("failed to read file %v with error: %q", filename, e.Error())
	}
	result, err := Read(fileBytes)
	if err != nil {
		t.Errorf("failed to convert file %v with error: %q", filename, err.Error())
	}
	fmt.Println(result)
}

func TestBerdSubmission(t *testing.T) {
	//Data version  0.0.3
	test.SetCwdToRoot()
	filename := "berd_0001"
	fileBytes, e := readFile(filename)
	if e != nil {
		t.Errorf("failed to read file %v with error: %q", filename, e.Error())
	}
	result, err := Read(fileBytes)
	if err != nil {
		t.Errorf("failed to convert file %v with error: %q", filename, err.Error())
	}
	fmt.Println(result)
}

func TestBresSubmission(t *testing.T) {
	//Supplementary Data and Data version  0.0.3
	test.SetCwdToRoot()
	filename := "bres_0019"
	fileBytes, e := readFile(filename)
	if e != nil {
		t.Errorf("failed to read file %v with error: %q", filename, e.Error())
	}
	result, err := Read(fileBytes)
	if err != nil {
		t.Errorf("failed to convert file %v with error: %q", filename, err.Error())
	}
	fmt.Println(result)
}
