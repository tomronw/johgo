package models

type WHSProducts struct {
	Hits []Hit `json:"hits"`
}

type Hit struct {
	CImages     []CImage `json:"c_images"`
	Price       *float64 `json:"price"`
	ProductName *string  `json:"product_name"`
	ProductID   *string  `json:"product_id"`
	CPageURL    *string  `json:"c_pageURL"`
}

type CImage struct {
	URL *string `json:"url"`
}
