package models

import (
	"time"
)

type ShopifyProducts struct {
	Products []Products `json:"products"`
}
type Variants struct {
	ID               int64       `json:"id"`
	Title            string      `json:"title"`
	Option1          string      `json:"option1"`
	Option2          interface{} `json:"option2"`
	Option3          interface{} `json:"option3"`
	Sku              string      `json:"sku"`
	RequiresShipping bool        `json:"requires_shipping"`
	Taxable          bool        `json:"taxable"`
	FeaturedImage    interface{} `json:"featured_image"`
	Available        bool        `json:"available"`
	Price            string      `json:"price"`
	Grams            int         `json:"grams"`
	CompareAtPrice   interface{} `json:"compare_at_price"`
	Position         int         `json:"position"`
	ProductID        int64       `json:"product_id"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}
type Images struct {
	ID         int64         `json:"id"`
	CreatedAt  time.Time     `json:"created_at"`
	Position   int           `json:"position"`
	UpdatedAt  time.Time     `json:"updated_at"`
	ProductID  int64         `json:"product_id"`
	VariantIds []interface{} `json:"variant_ids"`
	Src        string        `json:"src"`
	Width      int           `json:"width"`
	Height     int           `json:"height"`
}
type Options struct {
	Name     string   `json:"name"`
	Position int      `json:"position"`
	Values   []string `json:"values"`
}
type Products struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Handle      string     `json:"handle"`
	BodyHTML    string     `json:"body_html"`
	PublishedAt time.Time  `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Vendor      string     `json:"vendor"`
	ProductType string     `json:"product_type"`
	Tags        []string   `json:"tags"`
	Variants    []Variants `json:"variants"`
	Images      []Images   `json:"images"`
	Options     []Options  `json:"options"`
}
