package model

import (
	"fmt"
	"strings"
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

func (s *Schema) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{Title: %s\n", s.Title))
	sb.WriteString(fmt.Sprintf("SurveyId: %s\n", s.SurveyId))
	sb.WriteString(fmt.Sprintf("FormType: %s\n", s.FormType))
	sb.WriteString("Sections: [\n")
	for _, sec := range s.Sections {
		sb.WriteString(fmt.Sprintf("  {Title: %s\n", sec.Title))
		sb.WriteString("  Questions: [\n")
		for _, q := range sec.Questions {
			sb.WriteString(fmt.Sprintf("    {Title: %s\n", q.Title))
			sb.WriteString("    Answers: [\n")
			for _, a := range q.Answers {
				sb.WriteString(fmt.Sprintf("      {Type: %s\n", a.Type))
				sb.WriteString(fmt.Sprintf("      QCode: %s\n", a.QCode))
				sb.WriteString(fmt.Sprintf("      Label: %s}\n", a.Label))
			}
			sb.WriteString("    ]}\n")
		}
		sb.WriteString("  ]}\n")
	}
	sb.WriteString("]}\n")
	return sb.String()
}
