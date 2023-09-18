package submission

import (
	"encoding/json"
	"sdxImage/pkg/interfaces"
)

type Response struct {
	QCode    string `json:"questioncode"`
	Value    string `json:"response"`
	Instance int    `json:"instance"`
}

type Data map[string][]interfaces.Response

func (data *Data) UnmarshalJSON(bytes []byte) error {
	m := map[string]string{}
	err := json.Unmarshal(bytes, &m)
	if err == nil {
		*data = make(map[string][]interfaces.Response, len(m))
		for k, v := range m {
			(*data)[k] = []interfaces.Response{&Response{
				QCode:    k,
				Value:    v,
				Instance: 0,
			}}
		}
	} else {
		var respList []*Response
		err2 := json.Unmarshal(bytes, &respList)
		if err2 == nil {
			*data = make(map[string][]interfaces.Response)
			for _, v := range respList {
				if instList, found := (*data)[v.QCode]; found {
					instList = append(instList, v)
					(*data)[v.QCode] = instList
				} else {
					(*data)[v.QCode] = []interfaces.Response{v}
				}
			}
		} else {
			return err2
		}
	}
	return nil
}

func (response *Response) GetCode() string {
	return response.QCode
}

func (response *Response) GetValue() string {
	return response.Value
}

func (response *Response) GetInstance() int {
	return response.Instance
}
