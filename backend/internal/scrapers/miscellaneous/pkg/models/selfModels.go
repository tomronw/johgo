package models

type SelfProducts struct {
	Products []Product `json:"Products"`
}

type Product struct {
	ID *string `json:"id"`
	Na *string `json:"na"`
	Im *string `json:"im"`
	P  *string `json:"p"`
}
