package survey

import (
	"sort"
)

type Instance struct {
	Id      string
	Value   int
	Answers []*Answer
}

type Section struct {
	Title     string
	Instances map[string]*Instance
}

type UnitType int

const (
	None UnitType = iota
	LocalUnit
	PpiItem
)

type Survey struct {
	Title       string
	SurveyId    string
	FormType    string
	Respondent  string
	RuName      string
	SubmittedAt string
	Sections    []*Section
	Units       []Unit
	UnitType    UnitType
}

func (survey *Survey) String() string {
	result := ""
	for _, section := range survey.Sections {
		for _, instance := range section.Instances {
			for _, answer := range instance.Answers {
				result += answer.GetCode() + ", "
			}
		}
	}
	for _, unit := range survey.Units {
		for _, answer := range unit.GetAnswers() {
			result += answer.GetCode() + ", "
		}
	}
	return result
}

type InstanceList []*Instance

func (a InstanceList) Len() int           { return len(a) }
func (a InstanceList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a InstanceList) Less(i, j int) bool { return a[i].Value < a[j].Value }

func (section *Section) GetInstances() []*Instance {
	var instances []*Instance
	for _, instance := range section.Instances {
		instances = append(instances, instance)
	}
	sort.Sort(InstanceList(instances))
	return instances
}
