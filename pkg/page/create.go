package page

import (
	"image"
	"sdxImage/pkg/model"
)

func Create(survey *model.Survey) image.Image {

	header := Header{
		SurveyName:  survey.Title,
		FormType:    survey.FormType,
		RuRef:       survey.Respondent,
		SubmittedAt: survey.SubmittedAt}

	page := createPage(header)

	for _, s := range survey.Sections {
		section := page.addSection(s.Title)
		for _, i := range s.Instances {
			instance := section.addInstance(i.Id)
			for _, a := range i.Answers {
				instance.addAnswer(a.QCode, a.Text, a.Value)
			}
		}
		section.complete()
	}

	return page.draw()
}
