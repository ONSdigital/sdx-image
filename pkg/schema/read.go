package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Schema struct {
	SurveyId string `json:"survey_id"`
}

var schema Schema

func Read() {

	jsonFile, err := os.Open("schemas/mbs_0106.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(bytes, &schema)
	if err != nil {
		return
	}

	fmt.Println(schema.SurveyId)

}
