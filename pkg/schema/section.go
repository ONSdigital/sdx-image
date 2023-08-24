package schema

type Section struct {
	id          string
	title       string
	QuestionIds []string
	groups      []*Group
}

func convertToSection(json map[string]any) (*Section, bool) {
	groups := convertList(json, "groups", convertToGroup)

	var questionIds []string
	for _, group := range groups {
		for _, block := range group.Blocks {
			questionIds = append(questionIds, block.Question.id)
		}
	}

	section := &Section{
		id:          getString(json, "id"),
		title:       extractTitle(json),
		QuestionIds: questionIds,
		groups:      groups,
	}
	return section, true
}

type Sections struct {
	titleList  []string
	SectionMap map[string]*Section
}

func newSections() *Sections {
	return &Sections{
		titleList:  []string{},
		SectionMap: make(map[string]*Section),
	}
}

func (sections *Sections) addSection(section *Section) {
	previousSection, exists := sections.SectionMap[section.title]
	if exists {
		for _, group := range section.groups {
			previousSection.groups = append(previousSection.groups, group)
		}
	} else {
		sections.titleList = append(sections.titleList, section.title)
		sections.SectionMap[section.title] = section
	}
}

func (sections *Sections) ListTitles() []string {
	return sections.titleList
}

func (sections *Sections) ListQuestions(sectionId string) []string {
	return sections.SectionMap[sectionId].QuestionIds
}
