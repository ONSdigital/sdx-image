// Package controller manages the entire process of generating images.
package controller

import (
	"image"
	"sdxImage/pkg/log"
	"sdxImage/pkg/model"
	"sdxImage/pkg/page"
	"sdxImage/pkg/read"
	"sdxImage/pkg/substitutions"
)

// Run orchestrates the steps required to create an image of the given submission
// This is done by generating a "model.Survey" populated with data from the submission,
// and information from the corresponding author read.
// The "page" package is then utilised to generate the actual image.
func Run(submissionBytes []byte) (image.Image, error) {
	submission, err := read.Submission(submissionBytes)
	if err != nil {
		log.Error("Unable to read submission", err)
		return nil, err
	}

	log.Info("Processing submission", submission.TxId)
	survey, err := read.Schema(submission.SchemaName)
	if err != nil {
		log.Error("Unable to read read", err, submission.TxId)
		return nil, err
	}
	survey = substitutions.Replace(survey, submission)
	survey = model.Add(survey, submission)
	return page.Create(survey), nil
}
