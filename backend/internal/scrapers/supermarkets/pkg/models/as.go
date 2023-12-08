package models

type AsProducts struct {
	Data Data `json:"data"`
}

type Data struct {
	TempoCmsContent TempoCmsContent `json:"tempo_cms_content"`
}

type TempoCmsContent struct {
	Zones []Zones `json:"zones"`
}

type Zones struct {
	Configs Configs `json:"configs,omitempty"`
}

type Configs struct {
	Products Products `json:"products"`
}

type Products struct {
	Items []Items `json:"items"`
}

type Items struct {
	Item      Item      `json:"item"`
	Inventory Inventory `json:"inventory"`
	Price     Price     `json:"price"`
}

type Item struct {
	Name   string `json:"name"`
	Images Images `json:"images"`
	SkuID  string `json:"sku_id"`
}

type Images struct {
	Scene7ID string `json:"scene7_id"`
}

type Inventory struct {
	AvailabilityInfo AvailabilityInfo `json:"availability_info"`
}

type AvailabilityInfo struct {
	Availability string `json:"availability"`
}

type Price struct {
	PriceInfo PriceInfo `json:"price_info"`
}

type PriceInfo struct {
	Price string `json:"price"`
}
