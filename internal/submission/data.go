package submission

import (
	"encoding/json"
)

type Answer struct {
	Id         string          `json:"answer_id"`
	Value      json.RawMessage `json:"value"`
	ListItemId string          `json:"list_item_id"`
}

type Answers []Answer

type SdMapping struct {
	Id         string `json:"identifier"`
	ListItemId string `json:"list_item_id"`
}

type List struct {
	Items      []string    `json:"items"`
	Name       string      `json:"name"`
	SdMappings []SdMapping `json:"supplementary_data_mappings"`
}

type Lists []List

type AnswerCode struct {
	Id   string `json:"answer_id"`
	Code string `json:"code"`
}

type AnswerCodes []AnswerCode

type ListData struct {
	Answers       Answers       `json:"answers"`
	Lists         Lists         `json:"lists"`
	AnswerCodes   AnswerCodes   `json:"answer_codes"`
	Supplementary Supplementary `json:"supplementary_data"`
}

type MapData map[string]string

type DataType string

const (
	ListDataType DataType = "list"
	MapDataType  DataType = "map"
)

type Data struct {
	DataType
	ListData
	MapData
}

func (data *Data) UnmarshalJSON(bytes []byte) error {
	m := map[string]string{}
	err := json.Unmarshal(bytes, &m)
	if err == nil {
		*data = Data{
			DataType: MapDataType,
			ListData: ListData{},
			MapData:  m,
		}
	} else {
		var listData ListData
		err = json.Unmarshal(bytes, &listData)
		if err == nil {
			*data = Data{
				DataType: ListDataType,
				ListData: listData,
				MapData:  MapData{},
			}
		} else {
			return err
		}
	}
	return nil
}

func (a *Answer) getValue() string {
	// If raw value is a JSON string, remove quotes
	if len(a.Value) > 0 && a.Value[0] == '"' && a.Value[len(a.Value)-1] == '"' {
		return string(a.Value[1 : len(a.Value)-1])
	} else if len(a.Value) > 0 && a.Value[0] == '[' && a.Value[len(a.Value)-1] == ']' {
		return string(a.Value[1 : len(a.Value)-1])
	} else {
		return string(a.Value)
	}
}

func (listData *ListData) getCode(answerId string) string {
	for _, answerCode := range listData.AnswerCodes {
		if answerCode.Id == answerId {
			return answerCode.Code
		}
	}
	return ""
}

func (listData *ListData) getListItemName(listItemId string) string {
	for _, list := range listData.Lists {
		for _, listItem := range list.Items {
			if listItem == listItemId {
				return list.Name
			}
		}
	}
	return ""
}

func (listData *ListData) getListItemIds(name string) []string {
	for _, list := range listData.Lists {
		if list.Name == name {
			return list.Items
		}
	}
	return nil
}

func (listData *ListData) getAllListItemIds() []string {
	var listItems []string
	for _, list := range listData.Lists {
		listItems = append(listItems, list.Items...)
	}
	return listItems
}

func (listData *ListData) getResponses(listItemId string) (map[string]string, []string) {
	responses := make(map[string]string)
	var order []string
	for _, answer := range listData.Answers {
		if answer.ListItemId == listItemId {
			code := listData.getCode(answer.Id)
			responses[code] = answer.getValue()
			order = append(order, code)
		}
	}
	return responses, order
}

func (listData *ListData) getLocalUnit(listItemId string) *LocalUnit {
	var sdMapping string
	for _, list := range listData.Lists {
		for _, mapping := range list.SdMappings {
			if mapping.ListItemId == listItemId {
				sdMapping = mapping.Id
			}
		}
	}

	for _, localUnit := range listData.Supplementary.Items.LocalUnits {
		if localUnit.Identifier == sdMapping {
			return localUnit
		}
	}

	return nil
}

func (listData *ListData) getPricesItem(listItemId string) *PricesItem {
	var sdMapping string
	for _, list := range listData.Lists {
		for _, mapping := range list.SdMappings {
			if mapping.ListItemId == listItemId {
				sdMapping = mapping.Id
			}
		}
	}

	for _, ppiItem := range listData.Supplementary.Items.PpiItemList {
		if ppiItem.Identifier == sdMapping {
			return ppiItem
		}
	}

	for _, ppiItem := range listData.Supplementary.Items.SppiItemList {
		if ppiItem.Identifier == sdMapping {
			return ppiItem
		}
	}

	return nil
}
