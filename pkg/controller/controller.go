package controller

import (
	"fmt"
	"image"
	"io"
	"os"
	"sdxImage/pkg/model"
	"sdxImage/pkg/page"
	"sdxImage/pkg/schema"
	"sdxImage/pkg/substitutions"
)

func Run(submission *model.Submission) (image.Image, error) {
	survey, err := schema.Read(submission.SchemaName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	survey = substitutions.Replace(survey, submission)
	survey = model.Add(survey, submission)
	return page.Create(survey), nil
}

func readFile(schemaName string) []byte {
	jsonFile, err := os.Open("examples/" + schemaName + ".json")
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)
	bytes, _ := io.ReadAll(jsonFile)
	return bytes
}
