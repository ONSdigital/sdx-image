package page

import (
	"image"
	s "sdxImage/internal/survey"
)

func Create(survey *s.Survey) image.Image {

	header := Header{
		SurveyName:  survey.Title,
		FormType:    survey.FormType,
		RuRef:       survey.Respondent,
		RuName:      survey.RuName,
		SubmittedAt: survey.SubmittedAt}

	page := createPage(header)

	for _, sect := range survey.Sections {
		section := page.addSection(sect.Title)
		instances := sect.GetInstances()
		for _, i := range instances {
			instance := section.addInstance(i.Value)
			for _, a := range i.Answers {
				instance.addAnswer(a.GetCode(), a.GetText(), a.GetValue())
			}
		}
		section.complete()
	}

	if len(survey.Units) != 0 {
		var section *Section
		if survey.UnitType == s.LocalUnit {
			section = page.addSection("Local Units")
		} else if survey.UnitType == s.PpiItem {
			section = page.addSection("Items")
		}
		instance := section.addInstance(0)
		for _, lu := range survey.Units {
			instance.addLocalUnit(lu)
		}
		section.complete()
	}

	return page.draw()
}
