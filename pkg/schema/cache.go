package schema

import "sdxImage/pkg/interfaces"

type SchemaCache struct {
	instruments map[string]interfaces.Schema
}

var Cache = SchemaCache{instruments: make(map[string]interfaces.Schema)}

func (schemaCache *SchemaCache) GetSchema(schemaName string) (interfaces.Schema, error) {
	guid := schemaName
	instrument, exists := schemaCache.instruments[guid]
	if exists {
		return instrument, nil
	}

	schema, err := Read(schemaName)
	if err != nil {
		return nil, err
	}

	instrument = convert(schema)
	schemaCache.instruments[guid] = instrument
	return instrument, nil
}
