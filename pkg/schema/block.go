package schema

type Block struct {
	Type string
	*Question
}

func convertToBlock(json map[string]any) (*Block, bool) {
	blockType := getString(json, "type")
	if blockType == "ListCollector" {
		json = getMap(json, "add_block")
		blockType = "Question"
	}
	if blockType != "Question" {
		return nil, false
	}
	block := &Block{
		Type:     getString(json, "type"),
		Question: convertField(json, "question", convertToQuestion),
	}
	return block, true
}
