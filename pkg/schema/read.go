package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Answer struct {
	Type  string
	QCode string
	Label string
}

type Question struct {
	Title   string
	Answers []Answer
}

type Schema struct {
	Title     string
	SurveyId  string
	FormType  string
	Questions []Question
}

func Read() {

	bytes := readFile()

	m := toMap(bytes)
	s := toSchema(m)
	fmt.Print(s)

}

func readFile() []byte {
	jsonFile, err := os.Open("schemas/mbs_0106.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	bytes, _ := io.ReadAll(jsonFile)
	return bytes
}

func toMap(bytes []byte) map[string]interface{} {
	m := map[string]interface{}{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func toSchema(m map[string]interface{}) Schema {
	title := m["title"].(string)
	surveyId := m["survey_id"].(string)
	formType := m["form_type"].(string)
	sections := m["sections"].([]interface{})
	schema := Schema{
		Title:     title,
		SurveyId:  surveyId,
		FormType:  formType,
		Questions: []Question{},
	}

	for _, s := range sections {
		section := s.(map[string]interface{})
		groups := section["groups"].([]interface{})
		for _, g := range groups {
			group := g.(map[string]interface{})
			if group["title"] != "Introduction" {
				blocks := group["blocks"].([]interface{})
				for _, b := range blocks {
					block := b.(map[string]interface{})
					if block["type"] == "Question" {
						q := block["question"].(map[string]interface{})

						questionTitle, ok := q["title"].(string)
						if !ok {
							questionTitle = q["title"].(map[string]interface{})["text"].(string)
						}

						question := Question{
							Title:   questionTitle,
							Answers: []Answer{},
						}

						answers := q["answers"].([]interface{})
						for _, a := range answers {
							ans := a.(map[string]interface{})
							label, exists := ans["label"]
							if !exists {
								label = "label"
							}
							answer := Answer{
								Type:  ans["type"].(string),
								QCode: ans["q_code"].(string),
								Label: label.(string),
							}
							question.Answers = append(question.Answers, answer)
						}

						schema.Questions = append(schema.Questions, question)
					}
				}
			}
		}

	}
	return schema
}

func Print(schema map[string]interface{}) {
	data, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println("")
	fmt.Println("--------------")
	fmt.Printf("%s\n", data)
}
