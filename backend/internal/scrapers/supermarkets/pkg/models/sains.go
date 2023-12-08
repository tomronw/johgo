package models

type SainsProduct struct {
	Productsz []Productss `json:"products"`
}

type Productss struct {
	ProductUID  string      `json:"product_uid"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	FullURL     string      `json:"full_url"`
	RetailPrice RetailPrice `json:"retail_price"`
	IsAvailable bool        `json:"is_available"`
}

type RetailPrice struct {
	Price float64 `json:"price"`
}
