package schema

type Group struct {
	Title  string
	Blocks []*Block
}

func convertToGroup(json map[string]any) (*Group, bool) {
	title := extractTitle(json)
	if title == "Introduction" {
		return nil, false
	}
	group := &Group{
		Title:  title,
		Blocks: convertList(json, "blocks", convertToBlock),
	}
	return group, true
}
