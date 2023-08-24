package read

import "encoding/json"

type Response struct {
	QCode    string `json:"questioncode"`
	Value    string `json:"response"`
	Instance int    `json:"instance"`
}

type Data struct {
	Responses []*Response
}

func (data *Data) UnmarshalJSON(bytes []byte) error {
	m := map[string]string{}
	err := json.Unmarshal(bytes, &m)
	if err == nil {
		responses := make([]*Response, len(m))
		i := 0
		for k, v := range m {
			responses[i] = &Response{
				QCode:    k,
				Value:    v,
				Instance: 0,
			}
			i++
		}
		data.Responses = responses
	} else {
		var responses []*Response
		err2 := json.Unmarshal(bytes, &responses)
		if err2 == nil {
			data.Responses = responses
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
