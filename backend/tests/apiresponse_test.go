package tests

import (
	"fmt"
	"johgo-search-engine/elastic"
	"testing"
)

func TestApiResponse(t *testing.T) {
	// Create an Elasticsearch client
	ec, err := elastic.CreateClient("")
	// Set up the search request body
	err, successful, res := ec.Query("elite trainer box", true)
	if err != nil || !successful {
		fmt.Printf("Error getting response: %s", err)
	}

	fmt.Printf("%s", res)

}
