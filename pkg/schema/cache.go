package schema

import (
	"sdxImage/pkg/interfaces"
	"sdxImage/pkg/log"
	"sync"
)

var mu sync.Mutex

type Cache struct {
	instruments map[string]interfaces.Schema
}

func NewCache() *Cache {
	return &Cache{instruments: make(map[string]interfaces.Schema)}
}

func (schemaCache *Cache) GetSchema(schemaName string) (interfaces.Schema, error) {
	mu.Lock()
	defer mu.Unlock()

	guid := schemaName
	instrument, exists := schemaCache.instruments[guid]
	if exists {
		log.Info("Found schema: " + schemaName + " in cache")
		return instrument, nil
	}

	log.Info("Schema: " + schemaName + " not in cache")
	schema, err := Read(schemaName)
	if err != nil {
		return nil, err
	}

	instrument = convert(schema)
	schemaCache.instruments[guid] = instrument
	return instrument, nil
}
