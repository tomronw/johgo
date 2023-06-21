package tests

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestApiResponse(t *testing.T) {
	// Create an Elasticsearch client
	ec, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:              []string{"https://localhost:9200"},
		Username:               "elastic",
		Password:               "u5d_J_V3*G6kQgQ=r73+",
		CertificateFingerprint: "3e0a8e827c6b4663e39ee35052370a72fbeba820b8cfdd3c60367e7c0dade47c",
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
	res, err := ec.Search(
		ec.Search.WithContext(context.Background()),
		ec.Search.WithBody(strings.NewReader("{ \"from\" : 0, \"size\" : 10000,\"query\": {\"match\": {\"Title\": \"elite\"}}}")),
		ec.Search.WithTrackTotalHits(true),
		ec.Search.WithPretty(),
	)
	if err != nil {
		fmt.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s", bodyBytes)

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
		fmt.Println("Error parsing the response body: ", err)
	}

}
