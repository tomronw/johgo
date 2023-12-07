package models

type ShopifyProducts struct {
	Products []Product `json:"products"`
}

type Product struct {
	Title    string    `json:"title"`
	Handle   string    `json:"handle"`
	Variants []Variant `json:"variants"`
	Images   []Images  `json:"images"`
}

type Variant struct {
	Available bool   `json:"available"`
	Price     string `json:"price"`
}

type Images struct {
	Src string `json:"src"`
}
