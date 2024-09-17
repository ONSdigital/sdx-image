// Package substitutions manages the interpolation of parameterised strings in a survey.
//
// Questions in a survey often contain parameterised strings that need to be replaced
// with values from the actual submission to read correctly:
// E.G: "What is the turnover of {ru_name} from {ref_p_start_date} to {ref_p_end_date}?"
// needs to be interpolated with the respondent unit name, and start and end dates.
//
// The questions may also include html tags, such as "<em>" that need to be replaced,
// and Date formats are also handled.
//
// Currently, we only support substitutions where the required data exists
// in the metadata of the survey, such as ru_name etc. We do not substitute
// values that can only be obtained from the respondents answers.
// In some cases we provide generic alternatives for presentation:
// E.G: "{total_turnover}" will be replaced with "the total turnover".
//
// The substitutions package contains the following files:
//
//   - replace.go
//     Defines the lookup types and contains the public functions to invoke the interpolation.
//
//   - parameter.go
//     Logic for handling parameters.
//
//   - date.go
//     Logic for handling dates.
//
//   - html.go
//     Logic for replacing html within strings.
package substitutions
