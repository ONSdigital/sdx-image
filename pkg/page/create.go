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
			for _, q := range i.Questions {
				for _, a := range q.Answers {
					if a.Value != "" {
						text := q.Title
						if a.Type == "Date" {
							text += " " + a.Label
						} else if a.Type == "Number" {
							text += " " + a.Label + ":"
						} else if a.Type == "Currency" {
							text = a.Label + "?"
						}
						section.addAnswer(a.QCode, text, a.Value)
					}
				}
			}
		}
	}

	return page.draw()
}
