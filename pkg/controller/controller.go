package controller

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"sdxImage/pkg/model"
	"sdxImage/pkg/page"
	"sdxImage/pkg/schema"
	"sdxImage/pkg/substitutions"
)

func Run(submission *model.Submission) image.Image {
	survey := schema.Read(submission.SchemaName)
	survey = substitutions.Replace(survey, submission)
	survey = model.Add(survey, submission)
	//fmt.Println(survey)
	return page.Create(survey)
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

func saveJPG(path string, im image.Image, quality int) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	var opt jpeg.Options
	opt.Quality = quality

	return jpeg.Encode(file, im, &opt)
}
