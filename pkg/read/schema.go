// Package read manages the reading of schemas and submissions into surveys
package read

import (
	"io"
	"os"
	"sdxImage/pkg/log"
	"sdxImage/pkg/model"
)

// Schema loads the requested schema and returns a *model.Schema.
func Schema(schemaName string) (*model.Schema, error) {
	bytes, err := loadFile(schemaName)
	if err != nil {
		log.Error("Failed to read schema", err)
		return nil, err
	}
	m, e := toCompleteMap(bytes)
	if e != nil {
		log.Error("Failed to convert schema bytes to map", e)
		return nil, e
	}
	s := convert(m)
	return s, nil
}

func loadFile(schemaName string) ([]byte, error) {
	jsonFile, err := os.Open("schemas/" + schemaName + ".json")
	defer func(jsonFile *os.File) {
		e := jsonFile.Close()
		if e != nil {
			log.Error("Failed to close schema file", e)
		}
	}(jsonFile)
	if err != nil {
		log.Error("Failed to open schema file", err)
		return nil, err
	}
	bytes, _ := io.ReadAll(jsonFile)
	return bytes, nil
}

func convert(m map[string]any) *model.Schema {
	title := getStringFrom(m, "title")
	surveyId := getStringFrom(m, "survey_id")
	formType := getStringFrom(m, "form_type")
	sections := getListFrom(m, "sections")

	dataVersion := getStringFrom(m, "data_version")
	var qCodeMap map[string]string
	if dataVersion == "0.0.3" {
		qCodeMap = getQcodeMap(m)
	}

	schema := model.Schema{
		Title:    title,
		SurveyId: surveyId,
		FormType: formType,
		Sections: []*model.Sect{},
	}

	createdSections := map[string]*model.Sect{}

	for _, s := range sections {
		sect := toMap(s)

		sectTitle := getOptionalStringField(sect, "title")
		section, exists := createdSections[sectTitle]
		if !exists {
			section = &model.Sect{
				Title:     sectTitle,
				Questions: []*model.Quest{},
			}
			createdSections[sectTitle] = section
		}

		groups := getListFrom(sect, "groups")
		for _, g := range groups {
			group := toMap(g)
			if getStringFrom(group, "title") != "Introduction" {
				blocks := getListFrom(group, "blocks")
				for _, b := range blocks {
					block := toMap(b)
					blockType := getStringFrom(block, "type")

					var q map[string]any
					if blockType == "Question" {
						q = getMapFrom(block, "question")
					} else if blockType == "ListCollector" {
						ab := getMapFrom(block, "add_block")
						q = getMapFrom(ab, "question")
					} else {
						continue
					}

					question := &model.Quest{
						Title:   locateStringFrom(q, "title", "text"),
						Answers: []*model.Ans{},
					}

					answers := getListFrom(q, "answers")
					for _, a := range answers {
						ans := toMap(a)
						label, hasLabel := ans["label"]
						if !hasLabel {
							label = "label"
						}

						qCode, found := ans["q_code"]
						if !found {
							id := getStringFrom(ans, "id")
							qCode = qCodeMap[id]
						}

						answer := &model.Ans{
							Type:  getStringFrom(ans, "type"),
							QCode: qCode.(string),
							Label: label.(string),
						}
						question.Answers = append(question.Answers, answer)
					}
					section.Questions = append(section.Questions, question)
				}
			}
		}
		schema.Sections = append(schema.Sections, section)
	}
	return &schema
}
