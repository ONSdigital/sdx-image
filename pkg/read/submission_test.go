package read

import (
	"fmt"
	"io"
	"os"
	"sdxImage/pkg/test"
	"testing"
)

func readFile(filename string) ([]byte, error) {
	jsonFile, err := os.Open("examples/" + filename + ".json")
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
	test.SetCwdToRoot()
	filename := "v1/abs_1802"
	fileBytes, e := readFile(filename)
	if e != nil {
		t.Errorf("failed to read file %v with error: %q", filename, e.Error())
	}
	result, err := Submission(fileBytes)
	if err != nil {
		t.Errorf("failed to convert file %v with error: %q", filename, err.Error())
	}
	fmt.Println(result)
}
