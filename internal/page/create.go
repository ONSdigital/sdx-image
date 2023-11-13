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

	return page.draw()
}
