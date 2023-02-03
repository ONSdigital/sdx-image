package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sdxImage/pkg/model"
)

func Read(schemaName string) *model.Survey {
	bytes := readFile(schemaName)
	m := toCompleteMap(bytes)
	s := toSurvey(m)
	return InterpolateParams(s)
}

func readFile(schemaName string) []byte {
	jsonFile, err := os.Open("schemas/" + schemaName + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	bytes, _ := io.ReadAll(jsonFile)
	return bytes
}

func toCompleteMap(bytes []byte) map[string]any {
	m := map[string]any{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}
