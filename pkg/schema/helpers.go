package schema

func getFieldFrom[T any](json map[string]any, key string) T {
	result, _ := json[key].(T)
	return result
}

var getString = getFieldFrom[string]
var getList = getFieldFrom[[]any]
var getMap = getFieldFrom[map[string]any]

type Converter[T any] func(json map[string]any) (T, bool)

func convertList[T any](json map[string]any, key string, convert Converter[T]) []T {
	list := getList(json, key)
	var result []T
	for _, t := range list {
		if obj, ok := t.(map[string]any); ok {
			x, success := convert(obj)
			if success {
				result = append(result, x)
			}
		}
	}
	return result
}

func convertField[T any](json map[string]any, key string, convert Converter[T]) T {
	x, _ := convert(getMap(json, key))
	return x
}
