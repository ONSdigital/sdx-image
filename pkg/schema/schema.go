package schema

import (
	"encoding/json"
	"fmt"
)

type Title string

type AnswerCode struct {
	AnswerId string `json:"answer_id"`
	Code     string `json:"code"`
}

type Option struct {
	Qcode string `json:"q_code"`
	Value string `json:"value"`
	Label string `json:"label"`
}

type Answer struct {
	Id         string   `json:"id"`
	Qcode      string   `json:"q_code"`
	AnswerType string   `json:"type"`
	Label      string   `json:"label"`
	Options    []Option `json:"options"`
}

type Question struct {
	Id      string   `json:"id"`
	Title   Title    `json:"title"`
	Answers []Answer `json:"answers"`
}

type Block struct {
	Id        string   `json:"id"`
	BlockType string   `json:"type"`
	Question  Question `json:"question"`
}

type Group struct {
	Id     string  `json:"id"`
	Title  Title   `json:"title"`
	Blocks []Block `json:"blocks"`
}

type Section struct {
	Id     string  `json:"id"`
	Title  Title   `json:"title"`
	Groups []Group `json:"groups"`
}

type Schema struct {
	Title       string       `json:"title"`
	SurveyId    string       `json:"survey_id"`
	FormType    string       `json:"form_type"`
	DataVersion string       `json:"data_version"`
	Sections    []Section    `json:"sections"`
	AnswerCodes []AnswerCode `json:"answer_codes"`
}

func (title *Title) UnmarshalJSON(bytes []byte) error {
	var t string
	err := json.Unmarshal(bytes, &t)
	if err == nil {
		*title = Title(t)
		return nil
	}

	var extendedTitle struct {
		Text string `json:"text"`
	}
	err2 := json.Unmarshal(bytes, &extendedTitle)
	if err2 != nil {
		return err2
	}

	*title = Title(extendedTitle.Text)
	return nil
}

func (schema *Schema) String() string {
	b, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
