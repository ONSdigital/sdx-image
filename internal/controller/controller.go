// Package controller manages the entire process of generating images.
package controller

import (
	"fmt"
	"image"
	"runtime"
	"sdxImage/internal/log"
	"sdxImage/internal/page"
	sch "sdxImage/internal/schema"
	sub "sdxImage/internal/submission"
	sur "sdxImage/internal/survey"
	"sync"
)

var mu sync.Mutex

var schemaCache = sch.NewCache(20, sch.Read)

// Run orchestrates the steps required to create an image of the given submission.
//
// This is done by creating an instance of 'survey.Survey' and passing
// it to the 'page' package which generates the actual image.
//
// Instances of 'survey.Survey' represent the combination of
// the survey schema (schema.Schema), which  defines the questions,
// and the respondents' submission (submission.Submission),
// which contains their answers.
//
// Schemas are retrieved through the 'schemaCache' which only reads in
// each schema once.
//
// Various parts of the drawing package are not thread safe and so the whole run function
// is synchronised.
func Run(submissionBytes []byte) (image.Image, error) {
	//synchronise the process
	mu.Lock()
	defer mu.Unlock()

	//get the submission
	submission, err := sub.Read(submissionBytes)
	if err != nil {
		log.Error("Unable to read submission", err)
		return nil, err
	}
	log.Info("Processing submission", submission.GetTxId())

	//get the schema
	schema, err := schemaCache.GetSchema(submission.GetSchemaName())
	if err != nil {
		log.Error("Unable to read schema", err, submission.GetTxId())
		return nil, err
	}

	//create the survey
	survey := sur.Create(schema, submission)
	fmt.Println("------------------")
	fmt.Println("Survey:", survey.String())
	fmt.Println("------------------")

	//generate the image
	result := page.Create(survey)
	log.Info("Successfully created image", submission.GetTxId())

	runtime.GC()
	return result, nil
}
