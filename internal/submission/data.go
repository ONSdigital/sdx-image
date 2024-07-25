package submission

import (
	"encoding/json"
)

type Response struct {
	Code         string `json:"questioncode"`
	Value        string `json:"response"`
	Instance     int    `json:"instance"`
	SdIdentifier string `json:"sd_identifier"`
}

type Data map[string][]Response

func (data *Data) UnmarshalJSON(bytes []byte) error {
	m := map[string]string{}
	err := json.Unmarshal(bytes, &m)
	if err == nil {
		*data = make(map[string][]Response, len(m))
		for k, v := range m {
			(*data)[k] = []Response{{
				Code:         k,
				Value:        v,
				Instance:     0,
				SdIdentifier: "",
			}}
		}
	} else {
		var respList []Response
		err = json.Unmarshal(bytes, &respList)
		if err == nil {
			*data = make(map[string][]Response)
			for _, v := range respList {
				if instList, found := (*data)[v.Code]; found {
					instList = append(instList, v)
					(*data)[v.Code] = instList
				} else {
					(*data)[v.Code] = []Response{v}
				}
			}
		} else {
			return err
		}
	}
	return nil
}

func (response *Response) GetCode() string {
	return response.Code
}

func (response *Response) GetValue() string {
	return response.Value
}

func (response *Response) GetInstance() int {
	return response.Instance
}

func (response *Response) GetSdIdentifier() string {
	return response.SdIdentifier
}
