package submission

const NonListItem = "non_list_item"

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

func (submission *Submission) GetSurveyId() string {
	return submission.SurveyMetadata.SurveyID
}

func (submission *Submission) GetDataType() DataType {
	return submission.Data.DataType
}

func (submission *Submission) GetListItemName(listItemId string) string {
	if submission.GetDataType() == ListDataType {
		return submission.Data.getListItemName(listItemId)
	}
	return ""
}

func (submission *Submission) GetListItemIds(name string) []string {
	if submission.GetDataType() == ListDataType {
		return submission.Data.getListItemIds(name)
	}
	return nil
}

func (submission *Submission) GetResponseForListId(listItemId string) map[string]string {
	if submission.GetDataType() == MapDataType {
		return submission.Data.MapData
	}

	return submission.Data.ListData.getResponses(listItemId)
}

// ResponseMap A mapping of list item id to a map of qcode to answer value
type ResponseMap map[string]map[string]string

func (submission *Submission) GetResponses() ResponseMap {
	if submission.GetDataType() == MapDataType {
		return ResponseMap{NonListItem: submission.Data.MapData}
	}
	listItemIds := submission.Data.getAllListItemIds()
	responses := ResponseMap{}
	for _, listItemId := range listItemIds {
		responses[listItemId] = submission.Data.ListData.getResponses(listItemId)
	}
	responses[NonListItem] = submission.Data.ListData.getResponses("")
	return responses
}

func (submission *Submission) GetNonUnitResponses() ResponseMap {
	if submission.GetDataType() == MapDataType {
		return ResponseMap{NonListItem: submission.Data.MapData}
	}
	listItemIds := submission.Data.getAllListItemIds()
	responses := ResponseMap{}
	for _, listItemId := range listItemIds {
		responses[listItemId] = submission.Data.ListData.getResponses(listItemId)
	}
	responses[NonListItem] = submission.Data.ListData.getResponses("")
	return responses
}

func (submission *Submission) GetLocalUnit(listItemId string) *LocalUnit {
	if submission.GetDataType() == ListDataType {
		return submission.Data.ListData.getLocalUnit(listItemId)
	}
	return nil
}

func (submission *Submission) GetPpiItem(listItemId string) *PpiItem {
	if submission.GetDataType() == ListDataType {
		return submission.Data.ListData.getPpiItem(listItemId)
	}
	return nil
}

func (submission *Submission) HasLocalUnits() bool {
	if submission.Data.ListData.Supplementary.Items.LocalUnits == nil {
		return false
	}
	return true
}

func (submission *Submission) HasPpiItems() bool {
	if submission.Data.ListData.Supplementary.Items.PpiItemList == nil {
		return false
	}
	return true
}
