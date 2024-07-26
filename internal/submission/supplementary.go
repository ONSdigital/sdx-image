package submission

import "strings"

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

func (su *LocalUnit) GetAddress() string {
	return strings.Join(su.Address, "\n")
}
