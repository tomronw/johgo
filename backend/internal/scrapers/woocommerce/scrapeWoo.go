package woocommerce

import (
	"context"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/scrapers/woocommerce/pkg/scrapers"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ScrapeWooCommerce() {

	ec, _ := elastic.CreateClient(core.Woocommerce.IndexName)

	for true {

		core.InfoLogger.Printf("Scraping woocommerce sites")

		wooSites, err := core.GetStoreList(core.Woocommerce.ApiRoute)

		// get new seed
		rand.Seed(time.Now().UnixNano())
		start := time.Now()

		if err == nil || wooSites.Success {
			scrapedProducts := make(chan elastic.IndexChannel, len(wooSites.ReturnedSites.Sites))

			for _, site := range wooSites.ReturnedSites.Sites {

				wg.Add(1)
				site := site
				go func() {
					go scrapeSite(site, scrapedProducts)
					defer wg.Done()
				}()

			}
			wg.Wait()

			core.InfoLogger.Printf("Finished waiting for all goroutines to finish. Parsing results...")
			ctx := context.Background()

			e, sent := ec.BulkAddProducts(scrapedProducts, ctx, start)

			if e == nil && sent {
				core.InfoLogger.Printf("Woo sleeping for %s seconds...", strconv.Itoa(config.GlobalDelay))
				time.Sleep(time.Duration(config.GlobalDelay) * time.Second)
			} else {
				core.ErrorLogger.Printf("Error sending products to elastic: %s, sleeping for global delay", e.Error())
				time.Sleep(time.Duration(config.GlobalDelay) * time.Second)
			}
		} else {
			core.ErrorLogger.Printf("Error: [%s], please make sure apiCalls is live | Polling...", err.Error())
			time.Sleep(5 * time.Second)
		}

	}
}

func scrapeSite(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.ScrapeWooCommerceSite(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}
