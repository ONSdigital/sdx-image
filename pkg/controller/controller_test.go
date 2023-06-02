package controller

import (
	"fmt"
	"io"
	"os"
	"sdxImage/pkg/test"
	"testing"
)

func getSubmission(filename string) ([]byte, error) {
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
	err = test.SaveJPG("examples/images/"+filename+".jpg", result, 100)
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

func TestV1Abs1808(t *testing.T) {
	runFromFile("v1/abs_1808", t)
}

func TestV1Abs1862(t *testing.T) {
	runFromFile("v1/abs_1862", t)
}

func TestV1Abs1874(t *testing.T) {
	runFromFile("v1/abs_1874", t)
}

func TestBricksV1(t *testing.T) {
	runFromFile("v1/bricks", t)
}

func TestBerd0001V1(t *testing.T) {
	runFromFile("v1/berd_0001", t)
}

func TestBerd0006V1(t *testing.T) {
	runFromFile("v1/berd_0006", t)
}

func TestUkis0001V1(t *testing.T) {
	runFromFile("v1/ukis_0001", t)
}

func TestUkisFail(t *testing.T) {
	runFromFile("v1/ukis_fail", t)
}

func TestSandLand(t *testing.T) {
	runFromFile("v1/sand_land", t)
}

func TestSandMarine(t *testing.T) {
	runFromFile("v1/sand_marine", t)
}

func TestCreditGrantors(t *testing.T) {
	runFromFile("v1/mgc_0001", t)
}
