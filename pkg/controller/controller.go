// Package controller manages the entire process of generating images.
package controller

import (
	"fmt"
	"image"
	"sdxImage/pkg/log"
	"sdxImage/pkg/model"
	"sdxImage/pkg/page"
	"sdxImage/pkg/schema"
	"sdxImage/pkg/substitutions"
)

// Run orchestrates the steps required to create an image of the given submission
// This is done by generating a "model.Survey" populated with data from the submission,
// and information from the corresponding author schema.
// The "page" package is then utilised to generate the actual image.
func Run(submission *model.Submission) (image.Image, error) {
	log.Info("Processing submission", submission.TxId)
	survey, err := schema.Read(submission.SchemaName)
	if err != nil {
		log.Error("Unable to read schema", err, submission.TxId)
		return nil, err
	}
	survey = substitutions.Replace(survey, submission)
	survey = model.Add(survey, submission)
	fmt.Println(survey)
	return page.Create(survey), nil
}
