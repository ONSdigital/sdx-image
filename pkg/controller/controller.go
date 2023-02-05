package controller

import (
	"fmt"
	"io"
	"os"
	"sdxImage/pkg/model"
	"sdxImage/pkg/schema"
	"sdxImage/pkg/substitutions"
)

func Run(schemaName string) {
	submissionBytes := readFile()
	submission := model.Convert(submissionBytes)
	survey := schema.Read(schemaName)
	survey = substitutions.Replace(survey, submission)
	fmt.Println(survey)
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
