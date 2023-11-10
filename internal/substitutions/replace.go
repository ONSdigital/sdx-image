// Package substitutions manages the interpolation of parameterised strings in a survey.
package substitutions

import (
	"strings"
)

// Replace performs the interpolation of parameterised strings of a survey with values from the submission.
// Some substitutions are hard coded to avoid overly complex knowledge of the specific survey.
func Replace(text string, lookup ParameterLookup) string {
	result := replaceParameters(text, lookup)
	result = html(result)
	return result
}

func DateFormat(dateString string) string {
	return convertSubmittedAt(dateString)
}

type ParameterLookup map[string]string

func (pLookup ParameterLookup) get(str string) string {
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

func GetLookup(startDate, endDate, ruName, employmentDate string) ParameterLookup {
	start := convertDate(startDate)
	end := convertDate(endDate)

	return ParameterLookup{
		"ref_p_start_date": start,
		"ref_p_end_date":   end,
		"ru_name":          ruName,
		"employment_date":  employmentDate,
		"total_turnover":   "the total turnover",
		"from":             "start date",
		"to":               "end date",
	}
}
