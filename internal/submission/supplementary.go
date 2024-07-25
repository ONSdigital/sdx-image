package submission

import "strings"

type SupplementaryUnit struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	Address    []string `json:"address"`
}

type Items struct {
	LocalUnits []SupplementaryUnit `json:"local-units"`
}

type Supplementary struct {
	Items Items `json:"items"`
}

func (su *SupplementaryUnit) GetAssociatedResponses(submission *Submission) []Response {
	responses := make([]Response, 0)
	for _, responseList := range submission.Data {
		for _, response := range responseList {
			if response.GetSdIdentifier() == su.Identifier {
				responses = append(responses, response)
			}
		}
	}
	return responses
}

func (su *SupplementaryUnit) GetAddress() string {
	return strings.Join(su.Address, "\n")
}
