package substitutions

import (
	"sdxImage/pkg/model"
)

func Replace(survey *model.Survey, submission *model.Submission) *model.Survey {

	lookup := parameterLookup{
		"ref_p_start_date": submission.SurveyMetaData.StartDate,
		"ref_p_end_date":   submission.SurveyMetaData.EndDate,
		"ru_name":          submission.SurveyMetaData.RuName,
	}

	for _, section := range survey.Sections {
		for _, question := range section.Questions {
			title := replaceParameters(question.Title, lookup)
			question.Title = html(title)
		}
	}
	return survey
}

type parameterLookup map[string]string

func (pLookup parameterLookup) get(str string) string {
	result, found := pLookup[str]
	if !found {
		return ""
	}
	return result
}
