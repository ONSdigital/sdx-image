package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sdxImage/pkg/model"
	"sdxImage/pkg/test"
	"testing"
)

func getSubmission(filename string) (*model.Submission, error) {
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
	var submission model.Submission
	err = json.Unmarshal(bytes, &submission)
	if err != nil {
		return nil, err
	}
	return &submission, nil
}

func TestRun(t *testing.T) {
	test.SetCwdToRoot()
	var filename = "abs_1802"
	submission, err := getSubmission(filename)
	if err != nil {
		t.Errorf("failed to read submission: %q", err)
	}
	result, e := Run(submission)
	if e != nil {
		t.Errorf("failed with error: %q", err)
	}
	err = test.SaveJPG("temp/"+filename+".jpg", result, 100)
	if err != nil {
		t.Errorf("failed to create image for %s with error: %q", filename, err.Error())
	}
}
