package main

import (
	"johgo-search-engine/api"
	"sync"
)

var wgMain sync.WaitGroup

// Launch whole johgo-search-engine

func main() {
	// go build ./backend.go
	// launch all scrapers
	//wgMain.Add(4)
	//go miscellaneous.ScrapeMiscellaneous()
	//go shopify.ScrapeShopifySites()
	//go woocommerce.ScrapeWooCommerce()
	//go supermarkets.ScrapeSupermarkets()

	// launch api
	api.ServeRouter()
	wgMain.Wait()
}
