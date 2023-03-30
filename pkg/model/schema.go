package model

import (
	"encoding/json"
	"fmt"
)

type Ans struct {
	Type  string
	QCode string
	Label string
}

type Quest struct {
	Title   string
	Answers []*Ans
}

type Sect struct {
	Title     string
	Questions []*Quest
}

type Schema struct {
	Title    string
	SurveyId string
	FormType string
	Sections []*Sect
}

func (schema *Schema) String() string {
	b, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}

func (sect *Sect) String() string {
	return sect.Title
}
