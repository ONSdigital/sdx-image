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
	idList     []string
	sectionMap map[string]*Section
}

func newSections() *Sections {
	return &Sections{
		idList:     []string{},
		sectionMap: make(map[string]*Section),
	}
}

func (sections *Sections) addSection(section *Section) {
	sections.idList = append(sections.idList, section.Id)
	sections.sectionMap[section.Id] = section
}

func (sections *Sections) GetSectionIds() []string {
	return sections.idList
}

func (sections *Sections) GetSectionTitle(sectionId string) string {
	return sections.sectionMap[sectionId].Title
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
