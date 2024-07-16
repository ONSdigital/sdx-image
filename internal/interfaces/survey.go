package interfaces

type Answer interface {
	GetCode() string
	GetText() string
	GetValue() string
}

type SupplementaryUnit interface {
	GetIdentifier() string
	GetName() string
	GetAddress() string
	GetAnswers() []Answer
}

type Instance interface {
	GetId() int
	GetAnswers() []Answer
}

type Section interface {
	GetTitle() string
	GetInstances() []Instance
}

type Survey interface {
	GetTitle() string
	GetSurveyId() string
	GetFormType() string
	GetRespondent() string
	GetRuName() string
	GetSubmittedAt() string
	GetSections() []Section
	GetLocalUnits() []SupplementaryUnit
}
