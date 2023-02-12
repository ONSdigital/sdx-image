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
