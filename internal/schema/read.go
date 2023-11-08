package schema

import (
	"encoding/json"
	"io"
	"os"
	"sdxImage/internal/log"
)

func Read(schemaName string) (*Schema, error) {
	byteValue, err := loadFile(schemaName)
	if err != nil {
		return nil, err
	}
	schema := &Schema{}
	err = json.Unmarshal(byteValue, schema)
	if err != nil {
		return nil, err
	}
	return schema, nil
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
