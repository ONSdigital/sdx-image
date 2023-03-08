package controller

import (
	"fmt"
	"io"
	"os"
	"sdxImage/pkg/test"
	"testing"
)

func getSubmission(filename string) ([]byte, error) {
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

func runFromFile(filename string, t *testing.T) {
	test.SetCwdToRoot()
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

func TestAbs(t *testing.T) {
	runFromFile("abs_1802", t)
}

func TestMbs(t *testing.T) {
	runFromFile("mbs_0106", t)
}

func TestMbsV1(t *testing.T) {
	runFromFile("v1/mbs_0106", t)
}

func TestAbsV1(t *testing.T) {
	runFromFile("v1/abs_1802", t)
}

func TestAbsV2(t *testing.T) {
	runFromFile("v1/abs_1824", t)
}

func TestAbsV3(t *testing.T) {
	runFromFile("v1/abs_1814", t)
}

func TestBricksV1(t *testing.T) {
	runFromFile("v1/bricks", t)
}
