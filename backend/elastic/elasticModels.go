package elastic

// JSON models for elastic

type ProductsToStore struct {
	Products []ElasticProduct `json:"products"`
}
type ElasticProduct struct {
	Title    string `json:"title"`
	Price    string `json:"price"`
	Url      string `json:"url"`
	Image    string `json:"image"`
	SiteName string `json:"siteName"`
	SiteUrl  string `json:"siteUrl"`
}

type ElasticCreds struct {
	Username               string
	Password               string
	CertificateFingerprint string
	MinScore               string
}

type ElasticReqError struct {
	Error  Error `json:"error"`
	Status int   `json:"status"`
}
type RootCause struct {
	Type         string `json:"type"`
	Reason       string `json:"reason"`
	ResourceType string `json:"resource.type"`
	ResourceID   string `json:"resource.id"`
	IndexUUID    string `json:"index_uuid"`
	Index        string `json:"index"`
}
type Error struct {
	RootCause    []RootCause `json:"root_cause"`
	Type         string      `json:"type"`
	Reason       string      `json:"reason"`
	ResourceType string      `json:"resource.type"`
	ResourceID   string      `json:"resource.id"`
	IndexUUID    string      `json:"index_uuid"`
	Index        string      `json:"index"`
}

type IndexChannel struct {
	SiteName      string          `json:"site"`
	ReturnProduct ProductsToStore `json:"returned_product"`
	Error         error           `json:"error"`
}

/* Elastic query response, for api */
type ElasticQuery struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}
type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}
type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}
type Source struct {
	Title    string `json:"Title"`
	Price    string `json:"Price"`
	URL      string `json:"Url"`
	Image    string `json:"Image"`
	SiteName string `json:"SiteName"`
	SiteURL  string `json:"SiteUrl"`
}
type QueryHits struct {
	Index  string  `json:"_index"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source"`
}
type Hits struct {
	Total     Total       `json:"total"`
	MaxScore  float64     `json:"max_score"`
	QueryHits []QueryHits `json:"hits"`
}

type BulkResponse struct {
	Errors bool `json:"errors"`
	Items  []struct {
		Index struct {
			ID     string `json:"_id"`
			Result string `json:"result"`
			Status int    `json:"status"`
			Error  struct {
				Type   string `json:"type"`
				Reason string `json:"reason"`
				Cause  struct {
					Type   string `json:"type"`
					Reason string `json:"reason"`
				} `json:"caused_by"`
			} `json:"error"`
		} `json:"index"`
	} `json:"items"`
}

type BulkAdd struct {
	ID      int
	Product ElasticProduct
}
