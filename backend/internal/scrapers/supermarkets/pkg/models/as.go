package models

type AsProducts struct {
	Data Data `json:"data"`
}
type AdditionalContents struct {
	Type    string  `json:"type"`
	Configs Configs `json:"configs"`
}
type Facets struct {
	Name  string  `json:"name"`
	Items []Items `json:"items"`
}
type TaxonomyInfo struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategoryRank int    `json:"category_rank"`
	DeptID       string `json:"dept_id"`
	DeptRank     int    `json:"dept_rank"`
	DeptName     string `json:"dept_name"`
	AisleID      string `json:"aisle_id"`
	AisleName    string `json:"aisle_name"`
	AisleRank    int    `json:"aisle_rank"`
	ShelfID      string `json:"shelf_id"`
	ShelfName    string `json:"shelf_name"`
	ShelfRank    int    `json:"shelf_rank"`
	PrimaryShelf string `json:"primary_shelf"`
}
type Images struct {
	Scene7ID            string `json:"scene7_id"`
	Scene7Host          string `json:"scene7_host"`
	SeoImageDescription string `json:"seo_image_description"`
	SeoImageTitle       string `json:"seo_image_title"`
}
type FreshnessInfo struct {
	IsSellByDateRequired bool `json:"is_sell_by_date_required"`
	MaxIdealDayValue     int  `json:"max_ideal_day_value"`
	MinIdealDayValue     int  `json:"min_ideal_day_value"`
}

type Item struct {
	SkuID                string        `json:"sku_id"`
	Cin                  string        `json:"cin"`
	Name                 string        `json:"name"`
	ItemName             string        `json:"item_name"`
	Description          string        `json:"description"`
	Brand                string        `json:"brand"`
	CartDesc             string        `json:"cart_desc"`
	PickerDesc           string        `json:"picker_desc"`
	AdditionalPickerDesc string        `json:"additional_picker_desc"`
	AdditionalInfo       string        `json:"additional_info"`
	SalesUnit            string        `json:"sales_unit"`
	IsPharmacyRestricted bool          `json:"is_pharmacy_restricted"`
	BundleItemCount      int           `json:"bundle_item_count"`
	IsColleagueDiscount  bool          `json:"is_colleague_discount"`
	IsSubstitutable      bool          `json:"is_substitutable"`
	UpcNumbers           []string      `json:"upc_numbers"`
	AllReplacementSkus   []interface{} `json:"all_replacement_skus"`
	TaxonomyInfo         TaxonomyInfo  `json:"taxonomy_info"`
	Images               Images        `json:"images"`
	Icons                []interface{} `json:"icons"`
	FreshnessInfo        FreshnessInfo `json:"freshness_info"`
	UntraitedStores      []string      `json:"untraited_stores"`
	IsHfss               bool          `json:"is_hfss"`
	IsClothingProduct    bool          `json:"is_clothing_product"`
	ClothingEssentials   interface{}   `json:"clothing_essentials"`
	BundledItems         []interface{} `json:"bundled_items"`
}
type AvailabilityInfo struct {
	Availability      string      `json:"availability"`
	MaxQuantity       int         `json:"max_quantity"`
	HasAlternates     bool        `json:"has_alternates"`
	AvailableQuantity interface{} `json:"available_quantity"`
}

type Inventory struct {
	SkuID                 string           `json:"sku_id"`
	Cin                   string           `json:"cin"`
	Upc                   string           `json:"upc"`
	StoreNumber           string           `json:"store_number"`
	AvailabilityInfo      AvailabilityInfo `json:"availability_info"`
	BundledItems          []interface{}    `json:"bundled_items"`
	AvailableReplacements []interface{}    `json:"available_replacements"`
}
type PriceInfo struct {
	Price            string      `json:"price"`
	PricePerUom      string      `json:"price_per_uom"`
	SalesUnit        string      `json:"sales_unit"`
	MinimumUnitPrice interface{} `json:"minimum_unit_price"`
	AvgWeight        int         `json:"avg_weight"`
	SalePrice        string      `json:"sale_price"`
}
type Price struct {
	Cin                string      `json:"cin"`
	Upc                string      `json:"upc"`
	SkuID              string      `json:"sku_id"`
	StoreNumber        string      `json:"store_number"`
	IsOnSale           bool        `json:"is_on_sale"`
	BundledItems       interface{} `json:"bundled_items"`
	PriceInfo          PriceInfo   `json:"price_info"`
	BundleDiscount     interface{} `json:"bundle_discount"`
	BundleDiscountInfo interface{} `json:"bundle_discount_info"`
}
type BasePromotion struct {
	ItemPromoType string `json:"item_promo_type"`
}
type PromotionInfo struct {
	SkuID         string        `json:"sku_id"`
	Cin           string        `json:"cin"`
	BasePromotion BasePromotion `json:"base_promotion"`
	Rollback      interface{}   `json:"rollback"`
	Linksave      interface{}   `json:"linksave"`
	LoyaltyInfo   interface{}   `json:"loyalty_info"`
}
type Items struct {
	ItemID        string          `json:"item_id"`
	IsBundle      bool            `json:"is_bundle"`
	Item          Item            `json:"item"`
	Inventory     Inventory       `json:"inventory"`
	Price         Price           `json:"price"`
	PromotionInfo []PromotionInfo `json:"promotion_info"`
}
type OverallInvalidItems struct {
	InvalidItems       []interface{} `json:"invalid_items"`
	InvalidPrices      []interface{} `json:"invalid_prices"`
	InvalidInventories []interface{} `json:"invalid_inventories"`
}
type Products struct {
	Items               []Items             `json:"items"`
	OverallInvalidItems OverallInvalidItems `json:"overall_invalid_items"`
}
type PositionChange struct {
	ID                     string `json:"id"`
	PositionChangeByMargin int    `json:"position_change_by_margin"`
}
type PreviouslyPurchased struct {
	ID                    string `json:"id"`
	IsPreviouslyPurchased bool   `json:"is_previously_purchased"`
}
type Configs struct {
	AdditionalContents               []AdditionalContents  `json:"additional_contents"`
	CurrentPage                      int                   `json:"current_page"`
	TotalRecords                     int                   `json:"total_records"`
	MaxPages                         int                   `json:"max_pages"`
	Skus                             []string              `json:"skus"`
	MonetizedItems                   []interface{}         `json:"monetized_items"`
	Facets                           []Facets              `json:"facets"`
	Products                         Products              `json:"products"`
	MonetizedProducts                interface{}           `json:"monetized_products"`
	AutoCorrectedTerm                string                `json:"auto_corrected_term"`
	DidYouMeanTerm                   string                `json:"did_you_mean_term"`
	RedirectURL                      interface{}           `json:"redirect_url"`
	QueryRelaxationStrategy          interface{}           `json:"query_relaxation_strategy"`
	Reranked                         bool                  `json:"reranked"`
	PersonalisedItems                interface{}           `json:"personalised_items"`
	PositionChange                   []PositionChange      `json:"position_change"`
	PreviouslyPurchased              []PreviouslyPurchased `json:"previously_purchased"`
	IsHookLogicInsert                bool                  `json:"is_hook_logic_insert"`
	HookLogicInsertProductCount      int                   `json:"hook_logic_insert_product_count"`
	SkipFirstHooklogicInsertPosition bool                  `json:"skip_first_hooklogic_insert_position"`
	Errors                           []interface{}         `json:"errors"`
}
type Viewports struct {
	Name         string `json:"name"`
	ProductCount int    `json:"product_count"`
}
type Zones struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	ModuleID string  `json:"module_id"`
	Configs  Configs `json:"configs,omitempty"`
}
type TempoCmsContent struct {
	Zones []Zones `json:"zones"`
}
type Data struct {
	TempoCmsContent TempoCmsContent `json:"tempo_cms_content"`
}
