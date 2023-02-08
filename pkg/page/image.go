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
		for _, q := range s.Questions {
			for _, a := range q.Answers {
				if a.Value != "" {
					text := q.Title
					section.addAnswer(a.QCode, text, a.Value)
				}
			}
		}
	}

	return page.draw()
}