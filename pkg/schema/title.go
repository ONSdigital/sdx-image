package schema

func extractTitle(json map[string]any) string {
	if result, ok := json["title"].(string); ok {
		return result
	}

	if titleMap, ok := json["title"].(map[string]any); ok {
		return titleMap["text"].(string)
	}

	return ""
}
