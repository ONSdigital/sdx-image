package controller

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"sdxImage/pkg/page"
	"sdxImage/pkg/schema"
	"sdxImage/pkg/submission"
	"sdxImage/pkg/substitutions"
)

func Run(schemaName string) {
	subBytes := readFile()
	sub := submission.From(subBytes)
	survey := schema.Read(schemaName)
	survey = substitutions.Replace(survey, sub)
	survey = submission.Add(survey, sub)
	fmt.Println(survey)

	err := saveJPG("images/"+schemaName+".jpg", page.Create(survey), 100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}

func readFile() []byte {
	jsonFile, err := os.Open("examples/mbs_0106.json")
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}
	defer jsonFile.Close()
	bytes, _ := io.ReadAll(jsonFile)
	return bytes
}

func saveJPG(path string, im image.Image, quality int) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	var opt jpeg.Options
	opt.Quality = quality

	return jpeg.Encode(file, im, &opt)
}
