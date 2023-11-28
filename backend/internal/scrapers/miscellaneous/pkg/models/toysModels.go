package models

type ToysResponse struct {
	Status string      `json:"status"`
	Path   interface{} `json:"path"`
	Pages  int         `json:"pages"`
	HTML   string      `json:"html"`
	Header string      `json:"header"`
}
