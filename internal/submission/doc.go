// Package submission manages the reading of respondent submission json.
//
// The submission json describes the respondents answers to a survey.
// This information is held in the data field of the submission and can
// take one of two formats:
//
//  1. A simple mapping between qcode and answer:
//     E.G: "data": {"123": "Yes", "456": "No"}
//
//  2. A more complex structure that facilitates multiple
//     answers for a single qcode, based on list item groupings:
//     E.G: "data": {
//     "answers": [
//     {
//     "answer_id": "12345",
//     "value": "Yes",
//     "list_item_id": "QjjNIP"
//     },
//     {
//     "answer_id": "67890",
//     "value": "No",
//     "list_item_id": "QjjNIP"
//     }
//     ],
//     "lists": [
//     {
//     "items": [
//     "QjjNIP"
//     ]
//     }
//     ],
//     "answer_codes": [
//     {
//     "answer_id": "12345",
//     "code": "123"
//     },
//     {
//     "answer_id": "67890",
//     "code": "123"
//     }
//     ]
//     }
//
// The data may also include "supplementary" data, which is previously collected
// data used to pre-populate the survey. In the "lists" field above in example 2
// a list of supplementary data mappings can exist which link a unit of supplementary
// data to a list item id.
//
// The submission will also contain metadata about the survey and the respondent.
//
// The submission package handles reading of the submission json into a Submission struct
// with useful getters to access the required information. The package contains the following files:
//
//   - submission.go
//     The types that make up the top level submission object.
//
//   - data.go
//     Handles unmarshalling and retrieving the data sections of the submission.
//
//   - read.go
//     Handles the actual reading of the submission json in bytes.
//
//   - supplementary.go
//     Defines types to handle the supplementary data and associated mappings.
//
//   - getters.go
//     Defines the getters for the submission type.
package submission
