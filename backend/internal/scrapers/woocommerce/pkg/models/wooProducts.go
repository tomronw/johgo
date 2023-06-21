package models

type WooCommerceProducts []struct {
	ID                          int           `json:"id"`
	Date                        string        `json:"date"`
	DateGmt                     string        `json:"date_gmt"`
	GUID                        GUID          `json:"guid"`
	Modified                    string        `json:"modified"`
	ModifiedGmt                 string        `json:"modified_gmt"`
	Slug                        string        `json:"slug"`
	Status                      string        `json:"status"`
	Type                        string        `json:"type"`
	Link                        string        `json:"link"`
	Title                       Title         `json:"title"`
	Content                     Content       `json:"content"`
	Excerpt                     Excerpt       `json:"excerpt"`
	FeaturedMedia               int           `json:"featured_media"`
	CommentStatus               string        `json:"comment_status"`
	PingStatus                  string        `json:"ping_status"`
	Template                    string        `json:"template"`
	Meta                        Meta          `json:"meta"`
	ProductCat                  []int         `json:"product_cat"`
	ProductTag                  []interface{} `json:"product_tag"`
	JetpackPublicizeConnections []interface{} `json:"jetpack_publicize_connections"`
	JetpackSharingEnabled       bool          `json:"jetpack_sharing_enabled"`
	JetpackLikesEnabled         bool          `json:"jetpack_likes_enabled"`
	Links                       Links         `json:"_links"`
}
type GUID struct {
	Rendered string `json:"rendered"`
}
type Title struct {
	Rendered string `json:"rendered"`
}
type Content struct {
	Rendered  string `json:"rendered"`
	Protected bool   `json:"protected"`
}
type Excerpt struct {
	Rendered  string `json:"rendered"`
	Protected bool   `json:"protected"`
}
type Meta struct {
	JetpackPostWasEverPublished    bool          `json:"jetpack_post_was_ever_published"`
	JetpackPublicizeMessage        string        `json:"jetpack_publicize_message"`
	JetpackIsTweetstorm            bool          `json:"jetpack_is_tweetstorm"`
	JetpackPublicizeFeatureEnabled bool          `json:"jetpack_publicize_feature_enabled"`
	JetpackSocialPostAlreadyShared bool          `json:"jetpack_social_post_already_shared"`
	JetpackSocialOptions           []interface{} `json:"jetpack_social_options"`
}
type Self struct {
	Href string `json:"href"`
}
type Collection struct {
	Href string `json:"href"`
}
type About struct {
	Href string `json:"href"`
}
type Replies struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}
type WpFeaturedmedia struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}
type WpAttachment struct {
	Href string `json:"href"`
}
type WpTerm struct {
	Taxonomy   string `json:"taxonomy"`
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}
type Curies struct {
	Name      string `json:"name"`
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}
type Links struct {
	Self            []Self            `json:"self"`
	Collection      []Collection      `json:"collection"`
	About           []About           `json:"about"`
	Replies         []Replies         `json:"replies"`
	WpFeaturedmedia []WpFeaturedmedia `json:"wp:featuredmedia"`
	WpAttachment    []WpAttachment    `json:"wp:attachment"`
	WpTerm          []WpTerm          `json:"wp:term"`
	Curies          []Curies          `json:"curies"`
}
