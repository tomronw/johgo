package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"johgo-search-engine/elastic"
	"strconv"
	"strings"
	"testing"
)

func TestBulkAdding(t *testing.T) {
	productsToStore := elastic.ProductsToStore{
		Products: []elastic.ElasticProduct{
			{
				Title:    "http://product1.com",
				Price:    "100",
				Url:      "http://product1.com",
				Image:    "http://product1.com/image.jpg",
				SiteName: "Site 1",
				SiteUrl:  "http://site1.com",
			},
			{
				Title:    "Product 2",
				Price:    "200",
				Url:      "http://product2.com",
				Image:    "http://product2.com/image.jpg",
				SiteName: "Site 2",
				SiteUrl:  "http://site2.com",
			},
			{
				Title:    "Product 3",
				Price:    "300",
				Url:      "http://product3.com",
				Image:    "http://product3.com/image.jpg",
				SiteName: "Site 3",
				SiteUrl:  "http://site3.com",
			},
			{
				Title:    "Product 4",
				Price:    "400",
				Url:      "http://product4.com",
				Image:    "http://product4.com/image.jpg",
				SiteName: "Site 4",
				SiteUrl:  "http://site4.com",
			},
			{
				Title:    "Product 5",
				Price:    "500",
				Url:      "http://product5.com",
				Image:    "http://product5.com/image.jpg",
				SiteName: "Site 5",
				SiteUrl:  "http://site5.com",
			},
			{
				Title:    "Product 6",
				Price:    "600",
				Url:      "http://product6.com",
				Image:    "http://product6.com/image.jpg",
				SiteName: "Site 6",
				SiteUrl:  "http://site6.com",
			},
		},
	}

	ec, _ := elastic.CreateClient("testindex")

	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:     ec.Instance, // The Elasticsearch client
		Index:      "testindex", // The default index name
		NumWorkers: 4,           // The number of worker goroutines (default: number of CPUs)
		FlushBytes: 5e+6,        // The flush threshold in bytes (default: 5M)
	})

	var res *esapi.Response

	if res, err = ec.Instance.Indices.Delete([]string{"testindex"}, ec.Instance.Indices.Delete.WithIgnoreUnavailable(true)); err != nil || res.IsError() {
		fmt.Printf("Cannot delete index: %s", err)
	}
	res.Body.Close()
	res, err = ec.Instance.Indices.Create("testindex")
	if err != nil {
		fmt.Printf("Cannot create index: %s", err)
	}
	if res.IsError() {
		fmt.Printf("Cannot create index: %s", res)
	}
	res.Body.Close()

	for n, product := range productsToStore.Products {

		p, _ := json.Marshal(product)
		fmt.Println(string(p))

		err = indexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				DocumentID: strconv.Itoa(n),
				Index:      "testindex",
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",
				// Body is an `io.Reader` with the payload
				Body: strings.NewReader(string(p)),

				// OnSuccess is the optional callback for each successful operation
				OnSuccess: func(
					ctx context.Context,
					item esutil.BulkIndexerItem,
					res esutil.BulkIndexerResponseItem,
				) {
					fmt.Printf("[%d] %s test/%s", res.Status, res.Result, item.DocumentID)
				},

				// OnFailure is the optional callback for each failed operation
				OnFailure: func(
					ctx context.Context,
					item esutil.BulkIndexerItem,
					res esutil.BulkIndexerResponseItem, err error,
				) {
					if err != nil {
						fmt.Printf("ERROR: %s", err)
					} else {
						fmt.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)

	}

	if err != nil {
		fmt.Printf("Unexpected error: %s", err)
	}

	// Close the indexer channel and flush remaining items
	//
	if err := indexer.Close(context.Background()); err != nil {
		fmt.Printf("Unexpected error: %s", err)
	}

	// Report the indexer statistics
	//
	stats := indexer.Stats()
	if stats.NumFailed > 0 {
		fmt.Printf("Indexed [%d] documents with [%d] errors", stats.NumFlushed, stats.NumFailed)
	} else {
		fmt.Printf("Successfully indexed [%d] documents", stats.NumFlushed)
	}

}
