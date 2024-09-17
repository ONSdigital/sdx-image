// Package survey defines the data types used to capture survey information
//
// To generate an image of the survey, both the Schema (the details of the questions)
// and the Submission (the respondents answers to the questions) are required.
// These are amalgamated to create the Survey type; an object that holds all the
// information required to describe the survey.
//
// The survey package contains the following files:
//
//   - survey.go
//     Defines the Survey type.
//
//   - code.go
//     Logic for extracting qcodes.
//
//   - create.go
//     Contains the public function for creating a Survey from a Submission and Schema.
//
//   - unit.go
//     Logic for working with elements defined as local units from the schema.
//
//   - answer.go
//     Defines the Answer type.
package survey
