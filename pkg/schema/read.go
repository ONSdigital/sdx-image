package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sdxImage/pkg/model"
)

func Read(schemaName string) (*model.Survey, error) {
	bytes, err := readFile(schemaName)
	if err != nil {
		return nil, err
	}
	m, e := toCompleteMap(bytes)
	if e != nil {
		return nil, e
	}
	s := toSurvey(m)
	return s, nil
}

func readFile(schemaName string) ([]byte, error) {
	jsonFile, err := os.Open("schemas/" + schemaName + ".json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	bytes, _ := io.ReadAll(jsonFile)
	return bytes, nil
}

func toCompleteMap(bytes []byte) (map[string]any, error) {
	m := map[string]any{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return m, nil
}
