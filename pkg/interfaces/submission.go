package interfaces

type Response interface {
	GetCode() string
	GetValue() string
	GetInstance() int
}

type Submission interface {
	GetTxId() string
	GetSchemaName() string
	GetRuRef() string
	GetRuName() string
	GetSubmittedAt() string
	GetStartDate() string
	GetEndDate() string
	GetDataVersion() string
	GetEmploymentDate() string
	GetResponses(code string) []Response
}
