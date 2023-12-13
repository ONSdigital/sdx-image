package submission

import (
	"strings"
)

type LocalUnit struct {
	Name    string   `json:"name"`
	Address []string `json:"address"`
}

type Items struct {
	LocalUnits []*LocalUnit `json:"local-units"`
}

type Supplementary struct {
	Items Items `json:"items"`
}

func (lu *LocalUnit) GetName() string {
	return lu.Name
}

func (lu *LocalUnit) GetAddress() string {
	return strings.Join(lu.Address, "\n")
}
