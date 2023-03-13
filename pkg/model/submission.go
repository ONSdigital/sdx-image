package model

type Submission struct {
	TxId        string
	SchemaName  string
	RuRef       string
	RuName      string
	SubmittedAt string
	StartDate   string
	EndDate     string
	DataVersion string
	Data        map[string]string
}

type SubmissionError struct {
	Msg string
}

func (e *SubmissionError) Error() string {
	return e.Msg
}

func MissingFields(s *Submission) []string {
	missing := make([]string, 0)
	if s.TxId == "" {
		missing = append(missing, "TxId")
	}
	if s.SchemaName == "" {
		missing = append(missing, "SchemaName")
	}
	if s.RuRef == "" {
		missing = append(missing, "RuRef")
	}
	if s.RuName == "" {
		missing = append(missing, "RuName")
	}
	if s.SubmittedAt == "" {
		missing = append(missing, "SubmittedAt")
	}
	if s.StartDate == "" {
		missing = append(missing, "StartDate")
	}
	if s.EndDate == "" {
		missing = append(missing, "EndDate")
	}
	if s.DataVersion == "" {
		missing = append(missing, "DataVersion")
	}
	return missing
}

func Add(schema *Schema, submission *Submission) *Survey {
	survey := &Survey{
		Title:       schema.Title,
		SurveyId:    schema.SurveyId,
		FormType:    schema.FormType,
		Respondent:  submission.RuRef,
		SubmittedAt: submission.SubmittedAt,
		Sections:    []*Section{},
	}
	for _, sect := range schema.Sections {
		instance := &Instance{
			Id:        0,
			Questions: []*Question{},
		}
		section := &Section{
			Title:     sect.Title,
			Instances: []*Instance{instance},
		}

		for _, quest := range sect.Questions {
			question := &Question{
				Title:   quest.Title,
				Answers: []*Answer{},
			}
			for _, ans := range quest.Answers {
				answer := &Answer{
					Type:  ans.Type,
					QCode: ans.QCode,
					Label: ans.Label,
					Value: "",
				}
				value, found := submission.Data[ans.QCode]
				if found {
					answer.Value = value
				}
				question.Answers = append(question.Answers, answer)
			}
			instance.Questions = append(instance.Questions, question)
		}
		survey.Sections = append(survey.Sections, section)
	}
	return survey
}
