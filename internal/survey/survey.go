package survey

import (
	"encoding/json"
	"fmt"
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

type Survey struct {
	Title       string
	SurveyId    string
	FormType    string
	Respondent  string
	RuName      string
	SubmittedAt string
	Sections    []*Section
	Units       []Unit
}

func (survey *Survey) String() string {
	b, err := json.MarshalIndent(survey, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
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
