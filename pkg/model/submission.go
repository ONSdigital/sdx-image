package model

type SurveyMetaData struct {
	SurveyId string `json:"survey_id"`
	Period   string `json:"period_id"`
	RuRef    string `json:"ru_ref"`
	RuName   string `json:"ru_name"`
	FormType string `json:"form_type"`
}

type Submission struct {
	TxId           string `json:"tx_id"`
	SubmittedAt    string `json:"submitted_at"`
	SchemaName     string `json:"schema_name"`
	SurveyMetaData `json:"survey_metadata"`
	Data           map[string]string `json:"data"`
}
