package survey

import "strings"

type Answer struct {
	Title    string
	QType    string
	QCode    string
	Label    string
	Value    string
	Multiple bool
}

func (answer *Answer) GetCode() string {
	return answer.QCode
}

func (answer *Answer) GetValue() string {
	return answer.Value
}

func (answer *Answer) GetText() string {
	text := answer.Title
	if answer.QType == "Date" {
		text += " " + answer.Label
	} else if answer.QType == "Number" {
		text += " " + answer.Label + ":"
	} else if answer.QType == "Currency" {
		// only include text and label if the question has multiple answers
		// and the label has 5 words or fewer
		if answer.Multiple && len(strings.Split(answer.Label, " ")) <= 5 {
			text += " " + answer.Label + ":"
		} else {
			text = answer.Label + "?"
		}

	} else if answer.QType == "Unit" {
		text += " " + answer.Label + ":"
	}
	return text
}
