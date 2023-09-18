// Package controller manages the entire process of generating images.
package controller

import (
	"image"
	"runtime"
	"sdxImage/pkg/log"
	"sdxImage/pkg/page"
	sch "sdxImage/pkg/schema"
	sub "sdxImage/pkg/submission"
	sur "sdxImage/pkg/survey"
	"sync"
)

var mu sync.Mutex

// Run orchestrates the steps required to create an image of the given submission.
//
// This is done by creating an instance of 'survey.Survey' and passing
// it to the "page" package which generates the actual image.
//
// Instances of 'interfaces.Survey' conform to the interfaces.Survey interface
// and represent the combination of the survey schema (interfaces.Schema), which
// defines the questions, and the respondents' submission (interfaces.Submission),
// which contains their answers.
//
// Various parts of the drawing package are not thread safe and so the whole run function
// is synchronised.
func Run(submissionBytes []byte) (image.Image, error) {
	mu.Lock()
	defer mu.Unlock()
	submission, err := sub.Read(submissionBytes)
	if err != nil {
		log.Error("Unable to read submission", err)
		return nil, err
	}

	log.Info("Processing submission", submission.GetTxId())
	schema, err := sch.Read(submission.GetSchemaName())
	if err != nil {
		log.Error("Unable to read schema", err, submission.GetTxId())
		return nil, err
	}
	survey := sur.Create(schema, submission)
	result := page.Create(survey)
	log.Info("Successfully created image", submission.GetTxId())
	runtime.GC()
	return result, nil
}
