package submission

import "strings"

type LocalUnit struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	Address    []string `json:"address"`
}

type PricesItem struct {
	Identifier        string `json:"identifier"`
	ItemNumber        string `json:"item_number"`
	SupplierNumber    string `json:"supplier_number"`
	ItemSpecification string `json:"item_specification_1"`
}

type Items struct {
	LocalUnits   []*LocalUnit  `json:"local-units"`
	PpiItemList  []*PricesItem `json:"item"`
	SppiItemList []*PricesItem `json:"service"`
}

type Supplementary struct {
	Items        Items  `json:"items"`
	CurrentMonth string `json:"current_month"` //required for ppi
}

func (su *LocalUnit) GetAddress() string {
	return strings.Join(su.Address, "\n")
}
