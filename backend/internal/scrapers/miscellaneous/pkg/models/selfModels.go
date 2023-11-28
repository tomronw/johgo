package models

type SelfProducts struct {
	Products   []Products `json:"Products"`
	Filters    Filters    `json:"Filters"`
	Type       string     `json:"type"`
	CmPagename string     `json:"cm_pagename"`
	DataLayer  DataLayer  `json:"dataLayer"`
}
type Products struct {
	Pe          bool          `json:"pe"`
	ID          string        `json:"id"`
	Sd          string        `json:"sd"`
	Na          string        `json:"na"`
	Sb          string        `json:"sb"`
	Br          string        `json:"br"`
	Ca          string        `json:"ca"`
	CategoryIds []interface{} `json:"categoryIds"`
	Sz          string        `json:"sz"`
	Ns          int           `json:"ns"`
	Im          string        `json:"im"`
	Ai          string        `json:"ai"`
	Cf          string        `json:"cf"`
	Bs          int           `json:"bs"`
	So          int           `json:"so"`
	Sp          float64       `json:"sp"`
	P           string        `json:"p"`
}
type SortFields struct {
	Title    string `json:"Title"`
	ID       string `json:"id"`
	Selected bool   `json:"Selected"`
}
type Information struct {
	TotalResults          int           `json:"totalResults"`
	SortFields            []SortFields  `json:"sortFields"`
	MyCategories          []interface{} `json:"myCategories"`
	MyBrands              []interface{} `json:"myBrands"`
	InititalFilterOptions []interface{} `json:"inititalFilterOptions"`
}
type SelfridgesValues struct {
	FhQuery           string `json:"fhQuery"`
	Title             string `json:"title"`
	Nr                int    `json:"nr"`
	FilterKey         string `json:"filterKey"`
	FilterValue       string `json:"filterValue"`
	HumanReadableTree string `json:"humanReadableTree"`
}
type Rows struct {
	Label            string             `json:"Label"`
	Type             string             `json:"Type"`
	FilterType       string             `json:"FilterType"`
	SelfridgesValues []SelfridgesValues `json:"Values"`
}
type Sections struct {
	Rows             []Rows `json:"Rows"`
	SectionSortOrder int    `json:"SectionSortOrder"`
}
type Filters struct {
	Information Information `json:"information"`
	Sections    []Sections  `json:"Sections"`
}
type DataLayer struct {
	PAGENAME           string `json:"{PAGE_NAME}"`
	PAGECATEGORYNAME   string `json:"{PAGE_CATEGORY_NAME}"`
	PAGECATEGORYID     string `json:"{PAGE_CATEGORY_ID}"`
	PAGETYPE           string `json:"{PAGE_TYPE}"`
	PAGESORT           string `json:"{PAGE_SORT}"`
	PAGENUMBER         string `json:"{PAGE_NUMBER}"`
	PAGEREFINERESULTS  string `json:"{PAGE_REFINE_RESULTS}"`
	PAGEREFINEBRAND    string `json:"{PAGE_REFINE_BRAND}"`
	PAGEREFINECOLOUR   string `json:"{PAGE_REFINE_COLOUR}"`
	PAGEREFINESIZE     string `json:"{PAGE_REFINE_SIZE}"`
	PAGEREFINEPRICE    string `json:"{PAGE_REFINE_PRICE}"`
	SEARCHPASS         string `json:"{SEARCH_PASS}"`
	SEARCHKEYWORD      string `json:"{SEARCH_KEYWORD}"`
	SEARCHCORRECTION   string `json:"{SEARCH_CORRECTION}"`
	SEARCHRESULTS      string `json:"{SEARCH_RESULTS}"`
	PAGEPRODUCTRESULTS string `json:"{PAGE_PRODUCT_RESULTS}"`
}
