package tests

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var r map[string]interface{}

func TestSearch(t *testing.T) {

	e := godotenv.Load(".env")
	if e != nil {
		panic(e.Error())
	}

	ec, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:              []string{"https://localhost:9200"},
		Username:               os.Getenv("ELASTIC_USERNAME"),
		Password:               os.Getenv("ELASTIC_PASSWORD"),
		CertificateFingerprint: os.Getenv("ELASTIC_CERT"),
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
	// Set up the search request body
	//query := "brilliant stars booster box"
	//minScore := "7"
	res, err := ec.Search(
		ec.Search.WithContext(context.Background()),
		ec.Search.WithBody(strings.NewReader("{ \"from\" : 0, \"size\" : 10000,\"query\": {\"match\": {\"title\": \"elite trainer box\"}}}")),
		//ec.Search.WithBody(strings.NewReader("{ \"from\" : 0, \"size\" : 10000,\"query\": {\"match_phrase\": {\"Title\": \"v box\"}}}")),
		//ec.Search.WithBody(strings.NewReader("{ \"from\" : 0, \"size\" : 10000,\"query\": {\"match_phrase\": {\"Title.keyword\": {\"query\": \"v box\", \"analyzer\": \"standard\"}}}")),
		//ec.Search.WithBody(strings.NewReader(fmt.Sprintf("{ \"from\" : 0, \"size\" : 10000, \"min_score\": %s, \"query\": {\"match\": {\"Title\": \"\"}}}", minScore, query))),
		ec.Search.WithTrackTotalHits(true),
		ec.Search.WithPretty(),
	)
	if err != nil {
		fmt.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			fmt.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			fmt.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		fmt.Printf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	fmt.Println(len(r["hits"].(map[string]interface{})["hits"].([]interface{})))
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		fmt.Printf(" * ID=%s, %s\n", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	fmt.Printf(strings.Repeat("=", 37))
}
