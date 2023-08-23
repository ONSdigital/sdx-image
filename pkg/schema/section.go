package schema

type Section struct {
	Id     string
	Title  string
	Groups []*Group
}

func convertToSection(json map[string]any) (*Section, bool) {
	section := &Section{
		Id:     getString(json, "id"),
		Title:  extractTitle(json),
		Groups: convertList(json, "groups", convertToGroup),
	}
	return section, true
}

type Sections struct {
	titleList  []string
	sectionMap map[string]*Section
}

func newSections() *Sections {
	return &Sections{
		titleList:  []string{},
		sectionMap: make(map[string]*Section),
	}
}

func (sections *Sections) addSection(section *Section) {
	previousSection, exists := sections.sectionMap[section.Title]
	if exists {
		for _, group := range section.Groups {
			previousSection.Groups = append(previousSection.Groups, group)
		}
	} else {
		sections.titleList = append(sections.titleList, section.Title)
		sections.sectionMap[section.Title] = section
	}
}

func (sections *Sections) GetSectionTitles() []string {
	return sections.titleList
}

func (sections *Sections) GetSectionQuestions(sectionId string) []string {
	section := sections.sectionMap[sectionId]
	var questions []string
	for _, group := range section.Groups {
		for _, block := range group.Blocks {
			questions = append(questions, block.Question.Id)
		}
	}
	return questions
}
