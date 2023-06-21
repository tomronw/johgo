package models

type JLProductss struct {
	PageInfo        PageInfo        `json:"pageInfo"`
	Results         int             `json:"results"`
	Facets          []Facets        `json:"facets"`
	PageInformation PageInformation `json:"pageInformation"`
	PagesAvailable  int             `json:"pagesAvailable"`
	JLProducts      []JLProducts    `json:"products"`
	ShowInStockOnly bool            `json:"showInStockOnly"`
	Crumbs          []Crumbs        `json:"crumbs"`
}
type PageInfo struct {
	AllowSearch    bool `json:"allowSearch"`
	AllowFiltering bool `json:"allowFiltering"`
}
type Details struct {
	Label          string `json:"label"`
	TrackingID     string `json:"trackingId"`
	FacetID        string `json:"facetId"`
	ColorSwatchURL string `json:"colorSwatchUrl"`
	Color          string `json:"color"`
	IsSelected     string `json:"isSelected"`
	Qty            string `json:"qty"`
}
type Facets struct {
	DimensionName string    `json:"dimensionName"`
	Tooltip       string    `json:"tooltip"`
	Details       []Details `json:"details"`
	Name          string    `json:"name"`
}
type PageInformation struct {
	Description string `json:"description"`
	Heading     string `json:"heading"`
	Title       string `json:"title"`
}
type PromoMessages struct {
	Offer                    string `json:"offer"`
	CustomPromotionalMessage string `json:"customPromotionalMessage"`
	PriceMatched             string `json:"priceMatched"`
	BundleHeadline           string `json:"bundleHeadline"`
}
type Value struct {
	Max string `json:"max"`
	Min string `json:"min"`
}
type VariantPriceRange struct {
	Value            Value         `json:"value"`
	ReductionHistory []interface{} `json:"reductionHistory"`
	For              string        `json:"for"`
}
type Price struct {
	Then2    string `json:"then2"`
	Was      string `json:"was"`
	Now      string `json:"now"`
	Currency string `json:"currency"`
	Then1    string `json:"then1"`
	Uom      string `json:"uom"`
}
type DynamicAttributes struct {
	Homeproducttype2                string `json:"homeproducttype2"`
	Duvetcoverstyle                 string `json:"duvetcoverstyle"`
	Dimensions                      string `json:"dimensions"`
	Material                        string `json:"material"`
	Tumbledry                       string `json:"tumbledry"`
	Producttype1                    string `json:"producttype1"`
	Careinstructions                string `json:"careinstructions"`
	Homearea                        string `json:"homearea"`
	Iron                            string `json:"iron"`
	Colour                          string `json:"colour"`
	Typethrowsbedspreadsandblankets string `json:"typethrowsbedspreadsandblankets"`
	Homeproducttype1                string `json:"homeproducttype1"`
	Truecolour                      string `json:"truecolour"`
	Brand                           string `json:"brand"`
	Fabric                          string `json:"fabric"`
	Crediteligibilitystatus         string `json:"crediteligibilitystatus"`
	Countryoforigin                 string `json:"countryoforigin"`
	Washinginstructions             string `json:"washinginstructions"`
	Creditofferingids               string `json:"creditofferingids"`
}
type ColorSwatches struct {
	IsAvailable    bool   `json:"isAvailable"`
	SkuID          string `json:"skuId"`
	ColorSwatchURL string `json:"colorSwatchUrl"`
	ImageURL       string `json:"imageUrl"`
	Color          string `json:"color"`
	BasicColor     string `json:"basicColor"`
}
type JLProducts struct {
	Fabric               string            `json:"fabric"`
	SwatchAvailable      bool              `json:"swatchAvailable"`
	PromoMessages        PromoMessages     `json:"promoMessages"`
	NonPromoMessage      string            `json:"nonPromoMessage"`
	IsProductSet         bool              `json:"isProductSet"`
	VariantPriceRange    VariantPriceRange `json:"variantPriceRange"`
	FabricByLength       bool              `json:"fabricByLength"`
	DefaultSkuID         string            `json:"defaultSkuId"`
	Reviews              int               `json:"reviews"`
	IsInStoreOnly        bool              `json:"isInStoreOnly"`
	Features             []interface{}     `json:"features"`
	ColorWheelMessage    string            `json:"colorWheelMessage"`
	ColorSwatchSelected  int               `json:"colorSwatchSelected"`
	Compare              bool              `json:"compare"`
	Brand                string            `json:"brand"`
	AgeRestriction       int               `json:"ageRestriction"`
	DisplaySpecialOffer  string            `json:"displaySpecialOffer"`
	Price                Price             `json:"price"`
	IsMadeToMeasure      bool              `json:"isMadeToMeasure"`
	Image                string            `json:"image"`
	EmailMeWhenAvailable bool              `json:"emailMeWhenAvailable"`
	PromotionalFeatures  []interface{}     `json:"promotionalFeatures"`
	ProductID            string            `json:"productId"`
	OutOfStock           bool              `json:"outOfStock"`
	IsAvailableToOrder   bool              `json:"isAvailableToOrder"`
	DynamicAttributes    DynamicAttributes `json:"dynamicAttributes,omitempty"`
	Title                string            `json:"title"`
	IsBundle             bool              `json:"isBundle"`
	ColorSwatches        []ColorSwatches   `json:"colorSwatches"`
	Code                 string            `json:"code"`
	AvailabilityMessage  string            `json:"availabilityMessage"`
	FutureRelease        bool              `json:"futureRelease"`
	MultiSku             bool              `json:"multiSku"`
	Type                 string            `json:"type"`
}
type Crumbs struct {
	Type        string `json:"type"`
	Clickable   string `json:"clickable"`
	DisplayName string `json:"displayName"`
}
