package schema

import (
	"encoding/json"
	"io"
	"os"
	"sdxImage/pkg/log"
	"sdxImage/pkg/model"
)

func Read(schemaName string) (*model.Survey, error) {
	bytes, err := readFile(schemaName)
	if err != nil {
		log.Error("Failed to read schema", err)
		return nil, err
	}
	m, e := toCompleteMap(bytes)
	if e != nil {
		log.Error("Failed to convert schema bytes to map", e)
		return nil, e
	}
	s := toSurvey(m)
	return s, nil
}

func readFile(schemaName string) ([]byte, error) {
	jsonFile, err := os.Open("schemas/" + schemaName + ".json")
	defer func(jsonFile *os.File) {
		e := jsonFile.Close()
		if e != nil {
			log.Error("Failed to close schema file", e)
		}
	}(jsonFile)
	if err != nil {
		log.Error("Failed to open schema file", err)
		return nil, err
	}
	bytes, _ := io.ReadAll(jsonFile)
	return bytes, nil
}

func toCompleteMap(bytes []byte) (map[string]any, error) {
	m := map[string]any{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		log.Error("Failed to unmarshal schema", err)
		return nil, err
	}
	return m, nil
}
