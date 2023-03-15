package model

import (
	"encoding/json"
	"fmt"
)

type Resp struct {
	QuestionCode string
	Response     string
	Instance     string
}

type Spp struct {
	FormType    string
	Reference   string
	Period      string
	Survey      string
	SubmittedAt string
	Responses   []*Resp
}

func (spp *Spp) String() string {
	b, err := json.MarshalIndent(spp, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}

func (spp *Spp) GetResp(qCode string) []Resp {
	var respList []Resp
	for _, resp := range spp.Responses {
		if resp.QuestionCode == qCode {
			respList = append(respList, *resp)
		}
	}
	return respList
}
