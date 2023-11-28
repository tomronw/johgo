package main

import (
	"johgo-search-engine/api"
	healthcheck "johgo-search-engine/internal/core/healthCheck"
	"johgo-search-engine/internal/scrapers/miscellaneous"
	"johgo-search-engine/internal/scrapers/shopify"
	"johgo-search-engine/internal/scrapers/supermarkets"
	"johgo-search-engine/internal/scrapers/woocommerce"
	"sync"
)

var wgMain sync.WaitGroup

func main() {
	// go build backend.go
	// launch all scrapers
	// TODO: add pointers to string in product fields to better handle empty strings
	// TODO: trim structs to avoid wasting compute power and add omitempty to pointers
	wgMain.Add(5)
	go miscellaneous.ScrapeMiscellaneous()
	go shopify.ScrapeShopifySites()
	go woocommerce.ScrapeWooCommerce()
	go supermarkets.ScrapeSupermarkets()

	// launch health check
	go healthcheck.SpawnHealthCheck()

	// launch api
	api.ServeRouter()
	wgMain.Wait()
}
