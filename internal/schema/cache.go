package schema

import (
	"sdxImage/internal/log"
	"sync"
	"time"
)

var mu sync.Mutex

type CreatorFunc func(string) (*Schema, error)

type Cache struct {
	size         int
	createSchema CreatorFunc
	instruments  map[string]*Schema
	lastUsed     map[string]int64
}

func NewCache(size int, createSchema CreatorFunc) *Cache {
	return &Cache{
		size:         size,
		createSchema: createSchema,
		instruments:  make(map[string]*Schema, size),
		lastUsed:     make(map[string]int64, size)}
}

func (cache *Cache) GetSchema(schemaName string) (*Schema, error) {
	mu.Lock()
	defer mu.Unlock()

	guid := schemaName

	if !cache.contains(guid) {

		//not in cache so invoke createSchema to create a new instrument
		instrument, err := cache.createSchema(schemaName)
		if err != nil {
			return nil, err
		}

		//remove oldest from cache if it is full
		if len(cache.instruments) == cache.size {
			oldest := cache.getOldest()
			cache.removeInstrument(oldest)
		}

		//add to cache
		cache.instruments[guid] = instrument
		cache.lastUsed[guid] = 0
	}

	return cache.getInstrument(guid), nil
}

func (cache *Cache) contains(guid string) bool {
	_, exists := cache.instruments[guid]
	if exists {
		log.Info("Schema: " + guid + " found in cache")
	} else {
		log.Info("Schema: " + guid + " not in cache")
	}
	return exists
}

// getInstrument returns the schema associated with the guid
// and records/overwrites the previous last used time.
func (cache *Cache) getInstrument(guid string) *Schema {
	cache.lastUsed[guid] = time.Now().UnixMilli()
	return cache.instruments[guid]
}

func (cache *Cache) removeInstrument(guid string) {
	log.Info("Removing schema: " + guid + " from cache")
	delete(cache.lastUsed, guid)
	delete(cache.instruments, guid)
}

func (cache *Cache) getOldest() string {
	var oldest string
	t := time.Now().UnixMilli()
	for guid, lastUseTime := range cache.lastUsed {
		if lastUseTime < t {
			oldest = guid
			t = lastUseTime
		}
	}
	return oldest
}
