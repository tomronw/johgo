package models

import "time"

type WHSProducts struct {
	V                       string                  `json:"_v"`
	Type                    string                  `json:"_type"`
	Count                   int                     `json:"count"`
	Hits                    []Hits                  `json:"hits"`
	Query                   string                  `json:"query"`
	Refinements             []Refinements           `json:"refinements"`
	SearchPhraseSuggestions SearchPhraseSuggestions `json:"search_phrase_suggestions"`
	SelectedRefinements     SelectedRefinements     `json:"selected_refinements"`
	SortingOptions          []SortingOptions        `json:"sorting_options"`
	Start                   int                     `json:"start"`
	Total                   int                     `json:"total"`
}
type ProductType struct {
	Type   string `json:"_type"`
	Master bool   `json:"master"`
}
type RepresentedProduct struct {
	Type string `json:"_type"`
	ID   string `json:"id"`
	Link string `json:"link"`
}
type CImages struct {
	URL    string  `json:"url"`
	Alt    *string `json:"alt"`
	Title  string  `json:"title"`
	Srcset string  `json:"srcset"`
	Zoom   string  `json:"zoom"`
}
type Inventory struct {
	StockStatus      string `json:"stock_status"`
	Orderable        bool   `json:"orderable"`
	Ats              int    `json:"ats"`
	StockLevel       int    `json:"stock_level"`
	CAvailabilityMsg string `json:"c_availabilityMsg"`
	Preorderable     bool   `json:"preorderable"`
	Backorderable    bool   `json:"backorderable"`
}
type Prices struct {
	WhsmithGbpSaleprices string  `json:"whsmith-gbp-saleprices"`
	WhsmithGbpListprices float64 `json:"whsmith-gbp-listprices"`
}
type Only2WhenYouBuyAnything struct {
	CCancellationDisallowed string    `json:"c_cancellationDisallowed"`
	CMainCategoryID         string    `json:"c_mainCategoryID"`
	CPromoType              string    `json:"c_promoType"`
	CShowOnDeliveryPage     string    `json:"c_showOnDeliveryPage"`
	CShowOnPDP              string    `json:"c_showOnPDP"`
	CShowOnPLP              string    `json:"c_showOnPLP"`
	CClass                  string    `json:"c_class"`
	CDetails                string    `json:"c_details"`
	CCalloutMsg             string    `json:"c_calloutMsg"`
	LastModified            time.Time `json:"lastModified"`
	CPageURL                string    `json:"c_pageUrl"`
}
type Only1249WhenYouBuyAnything struct {
	CCancellationDisallowed string    `json:"c_cancellationDisallowed"`
	CPromoType              string    `json:"c_promoType"`
	CShowOnDeliveryPage     string    `json:"c_showOnDeliveryPage"`
	CShowOnPDP              string    `json:"c_showOnPDP"`
	CShowOnPLP              string    `json:"c_showOnPLP"`
	CClass                  string    `json:"c_class"`
	CDetails                string    `json:"c_details"`
	CCalloutMsg             string    `json:"c_calloutMsg"`
	LastModified            time.Time `json:"lastModified"`
}
type HPWhenYouBuyAnything struct {
	CCancellationDisallowed string    `json:"c_cancellationDisallowed"`
	CPromoType              string    `json:"c_promoType"`
	CShowOnDeliveryPage     string    `json:"c_showOnDeliveryPage"`
	CShowOnPDP              string    `json:"c_showOnPDP"`
	CShowOnPLP              string    `json:"c_showOnPLP"`
	CClass                  string    `json:"c_class"`
	CDetails                string    `json:"c_details"`
	CCalloutMsg             string    `json:"c_calloutMsg"`
	LastModified            time.Time `json:"lastModified"`
}
type TestShippingPromo struct {
	CAboveThresholdCalloutMsg string    `json:"c_aboveThresholdCalloutMsg"`
	CBasketCallout            string    `json:"c_basketCallout"`
	CCancellationDisallowed   bool      `json:"c_cancellationDisallowed"`
	CDeliveryCallout          string    `json:"c_deliveryCallout"`
	CPdpCallout               string    `json:"c_pdpCallout"`
	CShowOnDeliveryPage       string    `json:"c_showOnDeliveryPage"`
	CShowOnPDP                bool      `json:"c_showOnPDP"`
	CShowOnPLP                string    `json:"c_showOnPLP"`
	CThreshold                int       `json:"c_threshold"`
	CUnderThresholdCalloutMsg string    `json:"c_underThresholdCalloutMsg"`
	CClass                    string    `json:"c_class"`
	CDetails                  string    `json:"c_details"`
	CCalloutMsg               string    `json:"c_calloutMsg"`
	LastModified              time.Time `json:"lastModified"`
}
type PromotionCustomAttributes struct {
	Only2WhenYouBuyAnything    Only2WhenYouBuyAnything    `json:"Only £2 When You Buy Anything"`
	Only1249WhenYouBuyAnything Only1249WhenYouBuyAnything `json:"Only £12.49 When You Buy Anything"`
	HPWhenYouBuyAnything       HPWhenYouBuyAnything       `json:"HP When You Buy Anything"`
	TestShippingPromo          TestShippingPromo          `json:"tests-shipping-promo"`
}
type CTopPriorityVariation struct {
	ID                        string                    `json:"id"`
	Name                      string                    `json:"name"`
	Brand                     string                    `json:"brand"`
	CProductFormat            string                    `json:"c_productFormat"`
	CCollection               string                    `json:"c_collection"`
	CContributor              string                    `json:"c_contributor"`
	CRenderingType            string                    `json:"c_renderingType"`
	CReservableInStore        interface{}               `json:"c_reservableInStore"`
	CCoverPrice               string                    `json:"c_coverPrice"`
	Price                     float64                   `json:"price"`
	PriceMax                  float64                   `json:"price_max"`
	PricePerUnit              float64                   `json:"price_per_unit"`
	PricePerUnitMax           float64                   `json:"price_per_unit_max"`
	CPreviousPriceLabel       string                    `json:"c_previousPriceLabel"`
	CPageURL                  string                    `json:"c_page_url"`
	PrimaryCategoryID         string                    `json:"primary_category_id"`
	CSupplierPrecedenceValue  string                    `json:"c_supplierPrecedenceValue"`
	ListPrice                 float64                   `json:"listPrice"`
	SalePrice                 string                    `json:"salePrice"`
	Inventory                 Inventory                 `json:"inventory"`
	Prices                    Prices                    `json:"prices"`
	PromotionCustomAttributes PromotionCustomAttributes `json:"promotionCustomAttributes"`
}
type Hits struct {
	Type                  string                `json:"_type"`
	Currency              string                `json:"currency"`
	HitType               string                `json:"hit_type"`
	Link                  string                `json:"link"`
	Price                 float64               `json:"price"`
	PricePerUnit          float64               `json:"price_per_unit"`
	ProductID             string                `json:"product_id"`
	ProductName           string                `json:"product_name"`
	ProductType           ProductType           `json:"product_type"`
	RepresentedProduct    RepresentedProduct    `json:"represented_product"`
	CImages               []CImages             `json:"c_images"`
	CPageURL              string                `json:"c_pageURL"`
	CPrimaryCategoryID    string                `json:"c_primary_category_id"`
	CTopPriorityVariation CTopPriorityVariation `json:"c_topPriorityVariation"`
}
type Values struct {
	Type     string `json:"_type"`
	HitCount int    `json:"hit_count"`
	Label    string `json:"label"`
	Value    string `json:"value"`
}
type Refinements struct {
	Type        string   `json:"_type"`
	AttributeID string   `json:"attribute_id"`
	Label       string   `json:"label"`
	Values      []Values `json:"values"`
}
type SuggestedPhrases struct {
	Type       string `json:"_type"`
	ExactMatch bool   `json:"exact_match"`
	Phrase     string `json:"phrase"`
}
type Terms struct {
	Type       string `json:"_type"`
	Completed  bool   `json:"completed"`
	Corrected  bool   `json:"corrected"`
	ExactMatch bool   `json:"exact_match"`
	Value      string `json:"value"`
}
type SuggestedTerms struct {
	Type         string  `json:"_type"`
	OriginalTerm string  `json:"original_term"`
	Terms        []Terms `json:"terms"`
}
type SearchPhraseSuggestions struct {
	Type             string             `json:"_type"`
	SuggestedPhrases []SuggestedPhrases `json:"suggested_phrases"`
	SuggestedTerms   []SuggestedTerms   `json:"suggested_terms"`
}
type SelectedRefinements struct {
	CProductFormat string `json:"c_productFormat"`
	Price          string `json:"price"`
}
type SortingOptions struct {
	Type  string `json:"_type"`
	ID    string `json:"id"`
	Label string `json:"label"`
}
