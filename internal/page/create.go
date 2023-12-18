package page

import (
	"image"
	"sdxImage/internal/interfaces"
)

func Create(survey interfaces.Survey) image.Image {

	header := Header{
		SurveyName:  survey.GetTitle(),
		FormType:    survey.GetFormType(),
		RuRef:       survey.GetRespondent(),
		RuName:      survey.GetRuName(),
		SubmittedAt: survey.GetSubmittedAt()}

	page := createPage(header)

	for _, s := range survey.GetSections() {
		section := page.addSection(s.GetTitle())
		for _, i := range s.GetInstances() {
			instance := section.addInstance(i.GetId())
			for _, a := range i.GetAnswers() {
				instance.addAnswer(a.GetCode(), a.GetText(), a.GetValue())
			}
		}
		section.complete()
	}

	if len(survey.GetLocalUnits()) != 0 {
		section := page.addSection("Local Units")
		instance := section.addInstance(0)
		for _, lu := range survey.GetLocalUnits() {
			instance.addLocalUnit(lu)
		}
		section.complete()
	}

	return page.draw()
}
