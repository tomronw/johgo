package models

type ArgProduct struct {
	Meta     Meta          `json:"meta"`
	Data     []Data        `json:"data"`
	Links    Links         `json:"links"`
	Errors   []interface{} `json:"errors"`
	Included []interface{} `json:"included"`
}
type AppliedSorts struct {
	Name  string `json:"name"`
	Order string `json:"order"`
}
type Children struct {
	ID               string     `json:"id"`
	Label            string     `json:"label"`
	Value            int        `json:"value"`
	ApplicableFilter string     `json:"applicableFilter"`
	Children         []Children `json:"children"`
}
type ArgosValues struct {
	ID               string     `json:"id"`
	Label            string     `json:"label"`
	Value            int        `json:"value"`
	ApplicableFilter string     `json:"applicableFilter"`
	Children         []Children `json:"children"`
}
type Aggregations struct {
	ID          string        `json:"id"`
	Label       string        `json:"label"`
	ArgosValues []ArgosValues `json:"values"`
}
type DisplayableSorts struct {
	PriceLowHigh   string `json:"Price: Low - High"`
	CustomerRating string `json:"Customer Rating"`
	PriceHighLow   string `json:"Price: High - Low"`
	MostPopular    string `json:"Most Popular"`
}
type Tracking struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Meta struct {
	ListerType                 string           `json:"listerType"`
	TotalData                  int              `json:"totalData"`
	PageSize                   int              `json:"pageSize"`
	CurrentPage                int              `json:"currentPage"`
	TotalPages                 int              `json:"totalPages"`
	Canned                     bool             `json:"canned"`
	CategoryMatch              bool             `json:"categoryMatch"`
	LeafMatch                  bool             `json:"leafMatch"`
	AppliedSorts               []AppliedSorts   `json:"appliedSorts"`
	AppliedFilters             []interface{}    `json:"appliedFilters"`
	Suggestions                []interface{}    `json:"suggestions"`
	ProductCategorySuggestions []interface{}    `json:"productCategorySuggestions"`
	Aggregations               []Aggregations   `json:"aggregations"`
	DisplayableSorts           DisplayableSorts `json:"displayableSorts"`
	Tracking                   []Tracking       `json:"tracking"`
	HfssRestricted             bool             `json:"hfssRestricted"`
}
type Badge struct {
}
type Variants struct {
}
type Highlights struct {
}
type DefiningAttributes struct {
}
type Attributes struct {
	RelevancyRank      int                `json:"relevancyRank"`
	RelevancyScore     float64            `json:"relevancyScore"`
	ProductID          string             `json:"productId"`
	Name               string             `json:"name"`
	Brand              string             `json:"brand"`
	Price              float64            `json:"price"`
	AvgRating          float64            `json:"avgRating"`
	ReviewsCount       int                `json:"reviewsCount"`
	Popularity         float64            `json:"popularity"`
	Nodes              []interface{}      `json:"nodes"`
	WcsID              string             `json:"wcsId"`
	Buyable            bool               `json:"buyable"`
	ImageURL           string             `json:"imageURL"`
	WasPrice           float64            `json:"wasPrice"`
	WasText            string             `json:"wasText"`
	Deliverable        bool               `json:"deliverable"`
	Reservable         bool               `json:"reservable"`
	FreeDelivery       bool               `json:"freeDelivery"`
	DeliveryCost       float64            `json:"deliveryCost"`
	SpecialOfferText   string             `json:"specialOfferText"`
	SpecialOfferCount  int                `json:"specialOfferCount"`
	Clearance          bool               `json:"clearance"`
	HasVariations      bool               `json:"hasVariations"`
	HasMultiplePrices  bool               `json:"hasMultiplePrices"`
	Badge              Badge              `json:"badge"`
	Variants           Variants           `json:"variants"`
	VariantsData       []interface{}      `json:"variantsData"`
	DetailAttributes   []interface{}      `json:"detailAttributes"`
	MissingTerms       []interface{}      `json:"missingTerms"`
	Highlights         Highlights         `json:"highlights"`
	Sorts              []interface{}      `json:"sorts"`
	DefiningAttributes DefiningAttributes `json:"definingAttributes"`
}
type Data struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}
type Links struct {
	First     string `json:"first"`
	Self      string `json:"self"`
	Next      string `json:"next"`
	Last      string `json:"last"`
	Canonical string `json:"canonical"`
}
