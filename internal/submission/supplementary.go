package submission

import (
	"sdxImage/internal/interfaces"
	"strings"
)

type LocalUnit struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	Address    []string `json:"address"`
}

type Items struct {
	LocalUnits []*LocalUnit `json:"local-units"`
}

type Supplementary struct {
	Items Items `json:"items"`
}

type SupplementaryUnit struct {
	LocalUnit
	Responses []interfaces.Response
}

func NewUnit(lu LocalUnit) *SupplementaryUnit {
	return &SupplementaryUnit{
		LocalUnit: lu,
		Responses: make([]interfaces.Response, 0),
	}
}

func (su *SupplementaryUnit) GetIdentifier() string {
	return su.Identifier
}

func (su *SupplementaryUnit) GetName() string {
	return su.Name
}

func (su *SupplementaryUnit) GetAddress() string {
	return strings.Join(su.Address, "\n")
}

func (su *SupplementaryUnit) GetResponses() []interfaces.Response {
	return su.Responses
}

func (su *SupplementaryUnit) AddResponses(response interfaces.Response) {
	su.Responses = append(su.Responses, response)
}
