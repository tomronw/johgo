package elastic

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"johgo-search-engine/api/logger"
	"johgo-search-engine/config"
	"johgo-search-engine/internal/core"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// Client for all elasticsearch calls
/*
This is a Go implementation of an Elasticsearch client. It defines an ElasticEngineClient type that represents a client for
all Elasticsearch calls. The CreateClient function initializes the client by creating a new Elasticsearch client object with
the provided Elasticsearch address and security credentials, and tests if the instance is running by calling the Info function.
The AddProduct function sends a product to Elasticsearch and indexes it with the provided index name and an auto-generated ID.
The WipeIndex function deletes the index with the provided index name. The IndexResults function indexes all scraped products
that are received through the provided channel once all of the results are in. The other variables defined are error messages.
*/

var (
	ErrorInstanceNotRunning = fmt.Sprintf("Elastic instance is not running, please make sure elasticsearch is running and listening, address: %s", config.ElasticSearchAddress)
	ErrorAddingProduct      = "Failed to add product to elastic index"
	ErrorSendingRequest     = "Error creating request to elastic instance"
	ErrorFormattingPayload  = "Error formatting payload for elastic"
	ErrorGettingResponse    = "Error getting response from elastic"
	ErrorParsingBody        = "Error parsing body from elastic"
	ErrorDeletingIndex      = "Error deleting index"
	negativeWords           = []string{"mystery", "joblot", "bulk", "paperback"}
)

// elastic client struct
type ElasticEngineClient struct {
	Instance *elasticsearch.Client
	Index    string
	MinScore string
}

// CreateClient Init elasticsearch client
func CreateClient(index string) (engineClient ElasticEngineClient, e error) {
	// Load elastic credentials from .env
	elasticCreds := LoadElasticSecurity()
	// will try and create client until it succeeds, usually fails because elastic isnt running
	for true {
		// Put elasticsearch config here, given initially when you launch ES
		ec, err := elasticsearch.NewClient(elasticsearch.Config{
			RetryOnStatus:          []int{502, 503, 504, 429},
			RetryBackoff:           func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond },
			MaxRetries:             5,
			Addresses:              config.ElasticSearchAddressConfig,
			Username:               elasticCreds.Username,
			Password:               elasticCreds.Password,
			CertificateFingerprint: elasticCreds.CertificateFingerprint,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Second,
				DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		},
		)
		if err == nil {
			res, err := ec.Info()
			if err == nil {

				fmt.Printf("min score: %s\n", elasticCreds.MinScore)

				defer res.Body.Close()
				core.InfoLogger.Println("elastic client started, version:", elasticsearch.Version)
				return ElasticEngineClient{
					Instance: ec,
					Index:    index,
					MinScore: elasticCreds.MinScore,
				}, nil

			} else {
				core.ErrorLogger.Printf("%s %s", ErrorInstanceNotRunning, err.Error())
				time.Sleep(5 * time.Second)
			}

		} else {

			core.ErrorLogger.Printf("%s %s", ErrorGettingResponse, err.Error())
			time.Sleep(5 * time.Second)

		}

	}
	return
}

func (ec ElasticEngineClient) BulkAddProducts(scrapedProducts chan IndexChannel, ctx context.Context, start time.Time) (e error, sent bool) {

	// https://github.com/elastic/go-elasticsearch/blob/main/_examples/bulk/indexer.go
	// bulk indexer for faster indexing, loops over the channel results and adds them to the bulk indexer
	// before finally sending them to elastic

	// create bulk indexer
	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:     ec.Instance,
		Index:      ec.Index,
		NumWorkers: 2,
		FlushBytes: 5e+6,
	})
	totalProducts := 0
	productsIndexed := 0
	docId := 0
	// loop over capacity of channel
	for i := 0; i < cap(scrapedProducts); i++ {
		// get results from channel
		siteResults := <-scrapedProducts
		// add to total products
		totalProducts += len(siteResults.ReturnProduct.Products)
		// loop over products
		for _, product := range siteResults.ReturnProduct.Products {
			// if product doesn't contain any negative words, add to bulk indexer
			if !containsAny(product.Title) {

				p, _ := json.Marshal(product)
				// add as bulk indexer item
				err = indexer.Add(
					ctx,
					esutil.BulkIndexerItem{
						DocumentID: strconv.Itoa(docId),
						Index:      ec.Index,
						Action:     "index",
						Body:       strings.NewReader(string(p)),

						OnSuccess: func(
							ctx context.Context,
							item esutil.BulkIndexerItem,
							res esutil.BulkIndexerResponseItem,
						) {
							productsIndexed++
						},

						OnFailure: func(
							ctx context.Context,
							item esutil.BulkIndexerItem,
							res esutil.BulkIndexerResponseItem, err error,
						) {
							if err != nil {
								fmt.Printf("error: %s", err)
							} else {
								fmt.Printf("error: %s: %s", res.Error.Type, res.Error.Reason)
							}
						},
					},
				)
				// increment doc id
				docId++
				if err != nil {
					core.ErrorLogger.Printf("Unexpected error: %s", err.Error())
				}
			}

		}
	}
	if err != nil {
		core.ErrorLogger.Printf("Unexpected error: %s", err.Error())
		return err, false
	}

	// close bulk indexer and get stats
	if err := indexer.Close(context.Background()); err != nil {
		core.ErrorLogger.Printf("Unexpected error: %s", err.Error())
		return err, false
	}
	// Get the stats and log them
	stats := indexer.Stats()
	// Code to measure
	duration := time.Since(start)
	if stats.NumFailed > 0 {
		core.InfoLogger.Printf("Indexed [%d] documents on [%v] in [%v] with [%d] errors\n", stats.NumFlushed, ec.Index, duration, stats.NumFailed)
	} else {
		fmt.Printf("Successfully indexed [%d] on [%v] documents in [%v]\n", stats.NumFlushed, ec.Index, duration)
	}
	return nil, true
}

func (ec ElasticEngineClient) ParseToNDJson(data []map[string]interface{}, dst *bytes.Buffer) error {
	enc := json.NewEncoder(dst)
	for _, element := range data {
		if err := enc.Encode(element); err != nil {
			if err != io.EOF {
				return fmt.Errorf("failed to parse NDJSON: %v", err)
			}
			break
		}
	}
	return nil
}

// Send product to ES

func (ec ElasticEngineClient) AddProduct(product ElasticProduct, iNum int, ctx context.Context) (e error, sent bool) {
	// Index is channel for product to go into, Document ID is incremented
	req := esapi.IndexRequest{
		Index:      ec.Index,
		DocumentID: strconv.Itoa(iNum),
		Body:       strings.NewReader(JsonStruct(product)),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, ec.Instance)

	if err == nil {
		defer res.Body.Close()

		if res.IsError() {
			core.ErrorLogger.Printf("%s ERROR indexing document ID=%d", res.Status(), iNum)
			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(string(bodyBytes))
			return errors.New(ErrorAddingProduct), false
		}
		return nil, true
	} else {
		core.ErrorLogger.Printf("%s: %s", ErrorAddingProduct, err)
		return err, false
	}
}

// Load ES config from .env

func LoadElasticSecurity() ElasticCreds {

	e := godotenv.Load(".env")
	if e != nil {
		panic(e.Error())
	}

	return ElasticCreds{
		Username:               os.Getenv("ELASTIC_USERNAME"),
		Password:               os.Getenv("ELASTIC_PASSWORD"),
		CertificateFingerprint: os.Getenv("ELASTIC_CERT"),
		MinScore:               os.Getenv("ELASTIC_MIN_SCORE"),
	}

}

func (ec ElasticEngineClient) WipeIndex(ctx context.Context) (e error) {
	// deletes index with provided index name
	// create delete request
	req := esapi.IndicesDeleteRequest{
		Index: []string{ec.Index},
	}
	// send request
	res, err := req.Do(ctx, ec.Instance)
	// no error on deleting index
	if err == nil {
		defer res.Body.Close()
		// double check if response is error
		if res.IsError() {
			core.ErrorLogger.Printf("%s- %s [%s]", ErrorDeletingIndex, ec.Index, res.Status())
			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			// load response into ElasticRequestError struct
			var elasticError ElasticReqError
			err = json.Unmarshal(bodyBytes, &elasticError)
			if err != nil {
				return errors.New("Error parsing body from elastic: " + err.Error())
			}

			return errors.New(elasticError.Error.Reason)
		}
		core.InfoLogger.Printf("Deleted index: %s", ec.Index)
		return nil
	} else {
		core.ErrorLogger.Printf("IndexRequest ERROR: %s", err)
		return err
	}
}

// Index once all results are in

func (ec ElasticEngineClient) IndexResults(scrapedProducts chan IndexChannel, ctx context.Context, start time.Time) (e error, successful bool) {
	// deprecated function, replaced by bulk indexer, kept for reference and testing
	totalProducts := 0
	productsIndexed := 0
	docIds := 1

	for i := 0; i < cap(scrapedProducts); i++ {

		siteResults := <-scrapedProducts

		totalProducts += len(siteResults.ReturnProduct.Products)

		for _, product := range siteResults.ReturnProduct.Products {

			if !containsAny(product.Title) {
				pulledP := ElasticProduct{
					Title:    product.Title,
					Price:    product.Price,
					Url:      product.Url,
					Image:    product.Image,
					SiteName: product.SiteName,
					SiteUrl:  product.SiteUrl,
				}
				err, sent := ec.AddProduct(pulledP, docIds, ctx)
				if !sent {
					core.ErrorLogger.Printf("%s: %s", ErrorSendingRequest, err.Error())
				} else {
					productsIndexed++
				}
				docIds++
			}
		}

	}

	// Code to measure
	duration := time.Since(start)

	core.InfoLogger.Printf("Finished process. Products found: %d | Products sent to elastic: %d | Taken: %v", totalProducts, productsIndexed, duration)
	return nil, true
}

// For the api to call, processes frontend query

func (ec ElasticEngineClient) Query(query string, filterSingles bool) (e error, successful bool, results []byte) {
	// handles query from frontend request
	// create products struct
	var ps []ElasticProduct
	pulledProducts := ProductsToStore{Products: ps}
	// create scope for search
	res, err := ec.Instance.Search(
		ec.Instance.Search.WithContext(context.Background()),
		ec.Instance.Search.WithBody(strings.NewReader(fmt.Sprintf("{ \"from\" : 0, \"size\" : 10000, \"min_score\": %s, \"query\": {\"match\": {\"title\": \"%s\"}}}", ec.MinScore, query))),
		ec.Instance.Search.WithTrackTotalHits(true),
		ec.Instance.Search.WithPretty(),
	)
	if err == nil {
		// no error so we read the response body
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			core.ErrorLogger.Printf(err.Error())
			return err, false, nil
		}
		// create elastic query struct
		var productStruct ElasticQuery
		// unmarshal response body into elastic query struct
		err = json.Unmarshal(bodyBytes, &productStruct)
		if err != nil {
			logger.ApiErrorLogger.Printf("%s: %s", ErrorParsingBody, err.Error())
			return err, false, nil
		}

		core.InfoLogger.Printf("total hits for query [%s]: %d", query, productStruct.Hits.Total.Value)
		// loop over hits and add to products struct, checking for duplicates before returning
		for _, hit := range productStruct.Hits.QueryHits {
			duplicate := false
			productStorageModel := ElasticProduct{}
			productStorageModel.Price = hit.Source.Price
			productStorageModel.Title = hit.Source.Title
			productStorageModel.Image = hit.Source.Image
			productStorageModel.Url = hit.Source.URL
			productStorageModel.SiteName = hit.Source.SiteName
			productStorageModel.SiteUrl = hit.Source.SiteURL
			for _, item := range pulledProducts.Products {
				if item.Url == hit.Source.URL {
					duplicate = true
				}
			}
			if !duplicate {
				pulledProducts.Products = append(pulledProducts.Products, productStorageModel)
			}
		}
		// check if we need to filter singles or not, loaded from the switch on the frontend
		if filterSingles {
			pulledProducts = FilterSingleCards(pulledProducts)
		}
		// marshal products struct into json
		jsonData, err := json.Marshal(pulledProducts)
		if err != nil {
			logger.ApiErrorLogger.Printf("%s: %s", ErrorParsingBody, err.Error())
			return err, false, nil
		}
		// return json data
		return nil, true, jsonData
	} else {
		logger.ApiErrorLogger.Printf("%s: %s", ErrorFormattingPayload, err.Error())
	}
	defer res.Body.Close()
	// if error, return error and marshall into json
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			core.ErrorLogger.Printf("Error parsing the response body: %s", err)
			return err, false, nil
		} else {
			core.ErrorLogger.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return err, false, nil
		}
	} else {

	}
	return
}

func FilterSingleCards(products ProductsToStore) ProductsToStore {
	// here we attempt to filter out single cards, we do this by checking if the title contains
	// any variation of set numbers, IE x/xxx or x / x or xxx etc
	// however this was pre release of 151, so we also check for 151 in the title and allow if so
	var filteredProducts ProductsToStore
	// loop over products
	for _, product := range products.Products {
		// match regex
		match, err := regexp.MatchString("[a-zA-Z]{0,2}[0-9]{1,3}/[a-zA-Z]{0,2}[0-9]{1,3}|[a-zA-Z]{0,2}[0-9]{1,3}\\s/\\s[a-zA-Z]{0,2}[0-9]{1,3}|[a-zA-Z]{0,2}[0-9]{3}", product.Title)
		if err != nil {
			// if error, log and return
			core.ErrorLogger.Printf("Error matching regex: %s", err.Error())
			return products
		}
		// if no match, add to filtered products
		if !match {
			filteredProducts.Products = append(filteredProducts.Products, product)
		} else if strings.Contains(product.Title, "151") {
			// if 151 in title, add to filtered products
			filteredProducts.Products = append(filteredProducts.Products, product)
		}
	}

	return filteredProducts
}

func containsAny(s string) bool {
	// check if string contains any of the negative words
	for _, substr := range negativeWords {
		if strings.Contains(strings.ToLower(s), substr) {
			return true
		}
	}
	return false
}

func JsonStruct(product ElasticProduct) string {
	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(product)
	if err != nil {
		fmt.Println("json.Marshal ERROR:", err.Error())
		return err.Error()
	}
	return string(b)
}
