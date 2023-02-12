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

func Add(survey *Survey, submission *Submission) *Survey {
	survey.Respondent = submission.RuRef
	survey.SubmittedAt = submission.SubmittedAt
	for _, section := range survey.Sections {
		for _, question := range section.Questions {
			for _, a := range question.Answers {
				value, found := submission.Data[a.QCode]
				if found {
					a.Value = value
				}
			}
		}
	}
	return survey
}
