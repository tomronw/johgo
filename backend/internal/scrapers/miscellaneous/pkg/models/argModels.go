package models

type ArgProduct struct {
	Data []Data `json:"data"`
}

type Data struct {
	ID         *string    `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
	WcsID *string  `json:"wcsId"`
}
