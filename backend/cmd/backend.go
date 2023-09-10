package main

import (
	"johgo-search-engine/api"
	"johgo-search-engine/internal/core/healthcheck"
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
