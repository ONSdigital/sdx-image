package read

import (
	"encoding/json"
	"sdxImage/pkg/log"
)

func toMap(x any) map[string]any {
	result, _ := x.(map[string]any)
	return result
}

func getFieldFrom[T any](json map[string]any, fieldName ...string) T {
	result, _ := json[fieldName[0]].(T)
	return result
}

func locateStringFrom(json map[string]any, fieldName ...string) string {
	result, ok := json[fieldName[0]].(string)
	if !ok && len(fieldName) > 1 {
		nextLevel := json[fieldName[0]].(map[string]any)
		return locateStringFrom(nextLevel, fieldName[1:]...)
	}
	return result
}

var getStringFrom = getFieldFrom[string]
var getListFrom = getFieldFrom[[]any]
var getMapFrom = getFieldFrom[map[string]any]

func getOptionalStringField(json map[string]any, fieldName string) string {
	result, found := json[fieldName].(string)
	if !found {
		return ""
	}
	return result
}

func toCompleteMap(bytes []byte) (map[string]any, error) {
	m := map[string]any{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		log.Error("Failed to unmarshal read", err)
		return nil, err
	}
	return m, nil
}
