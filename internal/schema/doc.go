// Package schema manages the reading of collection instrument schemas.
//
// Collection instrument schemas are json files that define how a survey will
// function in eQ. It describes the questions, routing, validation and metadata,
// as well as any other text that will be displayed to the user.
// As survey responses from eQ only contain answers, to generate an image of
// the survey the data describing the questions need to be taken from the schema.
//
// The schema package provides the capability to read a schema file and return a
// Schema object - a representation of the schema file with useful getters to
// access the required information. The package contains the following files:
//
//   - schema.go
//     The types that make up a schema object and the getter functions
//     to access the data.
//
//   - read.go
//     Functions to read a schema file from disk based on name.
//
//   - cache.go
//     A cache to store schema objects in memory to avoid repeatedly reading
//     the schema file from disk.
//
//   - answerSpec.go
//     Defines a bespoke type to represent the different kinds of answer
//     types that can be used in a survey.
package schema
