package submission

import "strings"

type LocalUnit struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	Address    []string `json:"address"`
}

// Section title wants item_number and item_specification_1
type Item struct {
	Identifier        string `json:"identifier"`
	ItemNumber        string `json:"item_number"`
	SupplierNumber    string `json:"supplier_number"`
	ItemSpecification string `json:"item_specification_1"`
}

type Items struct {
	LocalUnits []*LocalUnit `json:"local-units"`
	ItemList   []*Item      `json:"item"`
}

type Supplementary struct {
	Items Items `json:"items"`
}

func (su *LocalUnit) GetAddress() string {
	return strings.Join(su.Address, "\n")
}
