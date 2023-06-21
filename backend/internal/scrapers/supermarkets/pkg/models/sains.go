package models

type SainsProduct struct {
	Productsz []Productss `json:"products"`
	Controls  Controls    `json:"controls"`
}
type UnitPrice struct {
	Price         float64 `json:"price"`
	Measure       string  `json:"measure"`
	MeasureAmount int     `json:"measure_amount"`
}
type RetailPrice struct {
	Price   float64 `json:"price"`
	Measure string  `json:"measure"`
}
type Reviews struct {
	IsEnabled     bool   `json:"is_enabled"`
	ProductUID    string `json:"product_uid"`
	Total         int    `json:"total"`
	AverageRating int    `json:"average_rating"`
}
type Assets struct {
	PlpImage string        `json:"plp_image"`
	Images   []interface{} `json:"images"`
	Video    []interface{} `json:"video"`
}
type Categories struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Attributes struct {
	Brand []string `json:"brand"`
}
type Productss struct {
	ProductUID           string        `json:"product_uid"`
	FavouriteUID         interface{}   `json:"favourite_uid"`
	Eans                 []string      `json:"eans"`
	ProductType          string        `json:"product_type"`
	Name                 string        `json:"name"`
	Image                string        `json:"image"`
	ImageZoom            string        `json:"image_zoom"`
	ImageThumbnail       string        `json:"image_thumbnail"`
	ImageThumbnailSmall  string        `json:"image_thumbnail_small"`
	FullURL              string        `json:"full_url"`
	UnitPrice            UnitPrice     `json:"unit_price"`
	RetailPrice          RetailPrice   `json:"retail_price"`
	IsAvailable          bool          `json:"is_available"`
	Promotions           []interface{} `json:"promotions"`
	Associations         []interface{} `json:"associations"`
	IsAlcoholic          bool          `json:"is_alcoholic"`
	IsSpotlight          bool          `json:"is_spotlight"`
	IsIntolerant         bool          `json:"is_intolerant"`
	IsMhra               bool          `json:"is_mhra"`
	Badges               []interface{} `json:"badges"`
	Labels               []interface{} `json:"labels"`
	Zone                 interface{}   `json:"zone"`
	Department           interface{}   `json:"department"`
	Reviews              Reviews       `json:"reviews"`
	Breadcrumbs          []interface{} `json:"breadcrumbs"`
	Assets               Assets        `json:"assets"`
	Description          []interface{} `json:"description"`
	ImportantInformation []string      `json:"important_information"`
	Attachments          []interface{} `json:"attachments"`
	Categories           []Categories  `json:"categories"`
	DisplayIcons         []interface{} `json:"display_icons"`
	PdpDeepLink          string        `json:"pdp_deep_link"`
	Attributes           Attributes    `json:"attributes,omitempty"`
}
type Options struct {
	Display string `json:"display"`
	Value   string `json:"value"`
}
type Sort struct {
	Active  string    `json:"active"`
	Options []Options `json:"options"`
}
type Page struct {
	Active      int   `json:"active"`
	First       int   `json:"first"`
	Last        int   `json:"last"`
	Size        int   `json:"size"`
	SizeOptions []int `json:"size_options"`
}
type Values struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Value    string `json:"value"`
	Selected bool   `json:"selected"`
	Enabled  bool   `json:"enabled"`
}
type Filters struct {
	Key    string   `json:"key"`
	Label  string   `json:"label"`
	Type   string   `json:"type"`
	Values []Values `json:"values"`
}
type Controls struct {
	Sort                Sort      `json:"sort"`
	TotalRecordCount    int       `json:"total_record_count"`
	ReturnedRecordCount int       `json:"returned_record_count"`
	Page                Page      `json:"page"`
	Filters             []Filters `json:"filters"`
}
