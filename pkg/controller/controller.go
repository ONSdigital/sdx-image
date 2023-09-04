// Package controller manages the entire process of generating images.
package controller

import (
	"image"
	"runtime"
	"sdxImage/pkg/log"
	"sdxImage/pkg/page"
	"sdxImage/pkg/schema"
	"sdxImage/pkg/submission"
	"sync"
)

var mu sync.Mutex

// Run orchestrates the steps required to create an image of the given submission.
// This is done by generating a "model.Survey" populated with data from the submission,
// and information from the corresponding author read.
// The "page" package is then utilised to generate the actual image.
// Various parts of the drawing package are not thread safe and so the whole run function
// is synchronised.
func Run(submissionBytes []byte) (image.Image, error) {
	mu.Lock()
	defer mu.Unlock()
	sub, err := submission.Read(submissionBytes)
	if err != nil {
		log.Error("Unable to read submission", err)
		return nil, err
	}

	log.Info("Processing submission", sub.GetTxId())
	instrument, err := schema.Read(sub.GetSchemaName())
	if err != nil {
		log.Error("Unable to read schema", err, sub.GetTxId())
		return nil, err
	}
	survey := fromSubmission(instrument, sub)
	result := page.Create(survey)
	log.Info("Successfully created image", sub.GetTxId())
	runtime.GC()
	return result, nil
}
