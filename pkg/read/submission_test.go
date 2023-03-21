package read

import (
	"fmt"
	"io"
	"os"
	"sdxImage/pkg/model"
	"sdxImage/pkg/test"
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
	test.SetCwdToRoot()
	filename := "v1/abs_1808"
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

func TestGetResponsesFromList(t *testing.T) {
	var i float64 = 1
	resp := map[string]any{"questioncode": "200", "response": "Yes", "instance": i}
	data := []any{resp}
	m := map[string]any{"data": data}
	result := getResponsesFromList(m)[0]
	expected := model.Response{
		QuestionCode: "200",
		Value:        "Yes",
		Instance:     1,
	}

	if result.Instance != expected.Instance {
		t.Errorf("failed to get correct instance, instead got: %v", result.Instance)
	}

}
