package models

type SDResModel struct {
	Products []SDProducts `json:"products"`
}
type SDProducts struct {
	URL              *string  `json:"url,omitempty"`
	Image            *string  `json:"image,omitempty"`
	Name             *string  `json:"name,omitempty"`
	PriceUnFormatted *float64 `json:"priceUnFormatted,omitempty"`
}
