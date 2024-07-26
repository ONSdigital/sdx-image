package submission

func (submission *Submission) GetTxId() string {
	return submission.TxId
}

func (submission *Submission) GetSchemaName() string {
	return submission.SchemaName
}

func (submission *Submission) GetRuRef() string {
	return submission.SurveyMetadata.RuRef
}

func (submission *Submission) GetRuName() string {
	return submission.SurveyMetadata.RuName
}

func (submission *Submission) GetSubmittedAt() string {
	return submission.SubmittedAt
}

func (submission *Submission) GetStartDate() string {
	return submission.SurveyMetadata.StartDate
}

func (submission *Submission) GetEndDate() string {
	return submission.SurveyMetadata.EndDate
}

func (submission *Submission) GetEmploymentDate() string {
	return submission.SurveyMetadata.EmploymentDate
}

func (submission *Submission) GetDataType() DataType {
	return submission.Data.DataType
}

func (submission *Submission) GetListItemNames() []string {
	if submission.GetDataType() == ListDataType {
		return submission.Data.getListItemNames()
	}
	return nil
}

func (submission *Submission) GetListItemIds(name string) []string {
	if submission.GetDataType() == ListDataType {
		return submission.Data.getListItemIds(name)
	}
	return nil
}

func (submission *Submission) GetResponses(listItemId string) map[string]string {
	if submission.GetDataType() == MapDataType {
		return submission.Data.MapData
	}
	return submission.Data.ListData.getResponses(listItemId)
}

func (submission *Submission) GetLocalUnit(listItemId string) *LocalUnit {
	if submission.GetDataType() == ListDataType {
		return submission.Data.ListData.getLocalUnit(listItemId)
	}
	return nil
}

func (submission *Submission) GetLocalUnits() []*LocalUnit {
	if submission.GetDataType() == ListDataType {
		return submission.Data.ListData.Supplementary.Items.LocalUnits
	}
	return nil
}
