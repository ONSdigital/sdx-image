// Package substitutions manages the interpolation of parameterised strings in a survey.
package substitutions

import (
	"sdxImage/pkg/model"
	"strings"
)

// Replace performs the interpolation of parameterised strings of a survey with values from the submission.
// Some substitutions are hard coded to avoid overly complex knowledge of the specific survey.
func Replace(survey *model.Survey, submission *model.Submission) *model.Survey {

	submittedAt := convertSubmittedAt(submission.SubmittedAt)
	survey.SubmittedAt = submittedAt

	startDate := convertDate(submission.StartDate)
	endDate := convertDate(submission.EndDate)

	lookup := parameterLookup{
		"ref_p_start_date": startDate,
		"ref_p_end_date":   endDate,
		"ru_name":          submission.RuName,
		"total_turnover":   "the total turnover",
		"from":             "start date",
		"to":               "end date",
	}

	for _, section := range survey.Sections {
		for _, instance := range section.Instances {
			for _, question := range instance.Questions {
				title := replaceParameters(question.Title, lookup)
				question.Title = html(title)
			}
		}
	}
	return survey
}

type parameterLookup map[string]string

func (pLookup parameterLookup) get(str string) string {
	result, found := pLookup[str]
	if !found {
		if strings.HasSuffix(str, "from") {
			return "start date"
		} else if strings.HasSuffix(str, "to") {
			return "end date"
		} else {
			return ""
		}
	}
	return result
}
