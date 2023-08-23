package schema

import (
	"encoding/json"
	"io"
	"os"
	"sdxImage/pkg/log"
)

func Read(schemaName string) (*Schema, error) {
	bytes, _ := loadFile(schemaName)
	m := map[string]any{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		log.Error("Failed to unmarshall json", err)
		return nil, err
	}
	return convert(m), nil
}

func loadFile(schemaName string) ([]byte, error) {
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
