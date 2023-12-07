package models

type JLProductss struct {
	JLProducts []JLProducts `json:"products"`
}

type JLProducts struct {
	OutOfStock bool    `json:"outOfStock"`
	Image      *string `json:"image"`
	Price      Price   `json:"price"`
	Title      *string `json:"title"`
	ProductID  *string `json:"productId"`
}

type Price struct {
	Now *string `json:"now"`
}
