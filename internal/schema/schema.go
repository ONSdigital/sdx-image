package schema

import (
	"encoding/json"
	"fmt"
)

// VariableString represents a value that should be considered a string
// but may appear in the schema file as either a string or an object with a "text" field.
type VariableString string

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
	Id         string         `json:"id"`
	Qcode      string         `json:"q_code"`
	AnswerType string         `json:"type"`
	Label      VariableString `json:"label"`
	Options    []*Option      `json:"options"`
}

type Question struct {
	Id      string         `json:"id"`
	Title   VariableString `json:"title"`
	Answers []*Answer      `json:"answers"`
}

type Block struct {
	Id              string    `json:"id"`
	BlockType       string    `json:"type"`
	Question        *Question `json:"question"`
	AddBlock        *Block    `json:"add_block"`
	RepeatingBlocks []*Block  `json:"repeating_blocks"`
}

type Group struct {
	Id     string         `json:"id"`
	Title  VariableString `json:"title"`
	Blocks []*Block       `json:"blocks"`
}

type Section struct {
	Id     string         `json:"id"`
	Title  VariableString `json:"title"`
	Groups []*Group       `json:"groups"`
}

type Schema struct {
	Title       string        `json:"title"`
	SurveyId    string        `json:"survey_id"`
	FormType    string        `json:"form_type"`
	DataVersion string        `json:"data_version"`
	Sections    []*Section    `json:"sections"`
	AnswerCodes []*AnswerCode `json:"answer_codes"`
}

func (vString *VariableString) UnmarshalJSON(bytes []byte) error {
	var t string
	err := json.Unmarshal(bytes, &t)
	if err == nil {
		*vString = VariableString(t)
		return nil
	}

	var objectString struct {
		Text string `json:"text"`
	}
	err = json.Unmarshal(bytes, &objectString)
	if err != nil {
		return err
	}

	*vString = VariableString(objectString.Text)
	return nil
}

func (group *Group) getBlocks() []*Block {
	var blocks []*Block
	for _, block := range group.Blocks {
		if block.BlockType == "Question" {
			blocks = append(blocks, block)
		}
		if block.AddBlock != nil {
			blocks = append(blocks, block.AddBlock)
		}
		if block.RepeatingBlocks != nil {
			blocks = append(blocks, block.RepeatingBlocks...)
		}
	}
	return blocks
}

func (section *Section) GetTitle() string {
	return string(section.Title)
}

func (section *Section) GetQuestions() []*Question {
	var questions []*Question
	for _, group := range section.Groups {
		for _, block := range group.getBlocks() {
			questions = append(questions, block.Question)
		}
	}
	return questions
}

func (question *Question) GetTitle() string {
	return string(question.Title)
}

func (schema *Schema) String() string {
	b, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}

func (schema *Schema) GetTitle() string {
	return schema.Title
}

func (schema *Schema) GetSurveyId() string {
	return schema.SurveyId
}

func (schema *Schema) GetFormType() string {
	return schema.FormType
}

func (schema *Schema) GetSections() []*Section {
	return schema.Sections
}

func (schema *Schema) GetAnswerSpecs(question *Question) []*AnswerSpec {
	var answers []*AnswerSpec
	multiple := false
	if len(question.Answers) > 1 {
		multiple = true
	}
	for _, answer := range question.Answers {
		answers = append(answers, getAnswerSpecs(answer, schema, multiple)...)
	}

	return answers
}
