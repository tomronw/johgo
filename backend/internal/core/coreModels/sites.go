package coreModels

type SiteList struct {
	Sites []Site `json:"sites"`
}
type Site struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	ISO     string `json:"ISO"`
	Keyword string `json:"keywords"`
}
