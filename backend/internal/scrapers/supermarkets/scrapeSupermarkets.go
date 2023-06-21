package supermarkets

import (
	"context"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/scrapers/supermarkets/pkg/scrapers"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ScrapeSupermarkets() {

	ec, _ := elastic.CreateClient("supermarkets")

	for true {

		core.InfoLogger.Printf("Scraping supermarkets...")

		supermarketSites, err := core.GetStoreList("supermarkets")

		// get new seed
		rand.Seed(time.Now().UnixNano())
		start := time.Now()

		if err == nil || supermarketSites.Success {
			scrapedProducts := make(chan elastic.IndexChannel, len(supermarketSites.ReturnedSites.Sites))

			for _, site := range supermarketSites.ReturnedSites.Sites {

				wg.Add(1)
				site := site
				go func() {
					switch site.Name {
					case "Asda":
						go scrapeAsda(site, scrapedProducts)
						defer wg.Done()
					case "Sainsburys":
						go scrapeSainsburys(site, scrapedProducts)
						defer wg.Done()

					}
				}()

			}
			wg.Wait()

			core.InfoLogger.Printf("Finished waiting for all goroutines to finish. Parsing results...")
			ctx := context.Background()

			e, sent := ec.BulkAddProducts(scrapedProducts, ctx, start)
			if e == nil && sent {
				core.InfoLogger.Printf("Supermarkets sleeping for %s seconds...", strconv.Itoa(config.GlobalDelay))
				time.Sleep(time.Duration(config.GlobalDelay) * time.Second)
			} else {
				core.ErrorLogger.Printf("Error sending products to elastic: %s", e.Error())
				time.Sleep(time.Duration(config.GlobalDelay) * time.Second)
			}

		} else {
			core.ErrorLogger.Printf("Error: [%s], please make sure apiCalls is live | Polling...", err.Error())
			time.Sleep(5 * time.Second)
		}

	}

}

func scrapeAsda(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetAs(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func scrapeSainsburys(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetSains(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}
