package survey

import "sdxImage/pkg/interfaces"

type Section struct {
	Title     string
	Instances []interfaces.Instance
}

func (section *Section) GetTitle() string {
	return section.Title
}

func (section *Section) GetInstances() []interfaces.Instance {
	return section.Instances
}
