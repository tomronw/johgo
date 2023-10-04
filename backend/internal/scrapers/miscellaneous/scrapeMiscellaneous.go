package miscellaneous

import (
	"context"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/scrapers/miscellaneous/pkg/scrapers"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ScrapeMiscellaneous() {

	ec, _ := elastic.CreateClient("miscellaneous")

	for true {

		core.InfoLogger.Printf("Scraping miscellaneous sites...")

		miscellaneousSites, err := core.GetStoreList("miscellaneous")

		// get new seed
		rand.Seed(time.Now().UnixNano())
		start := time.Now()

		if err == nil || miscellaneousSites.Success {
			scrapedProducts := make(chan elastic.IndexChannel, len(miscellaneousSites.ReturnedSites.Sites))

			for _, site := range miscellaneousSites.ReturnedSites.Sites {

				wg.Add(1)
				site := site
				go func() {
					switch site.Name {
					case "WHSmith":
						go scrapeWHSmith(site, scrapedProducts)
						defer wg.Done()
					case "Selfridges":
						go scrapeSelfridges(site, scrapedProducts)
						defer wg.Done()
					case "John Lewis":
						go scrapeJohnLewis(site, scrapedProducts)
						defer wg.Done()
					case "Argos":
						go scrapeArgos(site, scrapedProducts)
						defer wg.Done()
					case "ToysRUs":
						go scrapeToysRUs(site, scrapedProducts)
						defer wg.Done()
						//default:
						//	defer wg.Done()
					}
				}()

			}
			wg.Wait()

			core.InfoLogger.Printf("Finished waiting for all goroutines to finish. Parsing results...")
			ctx := context.Background()

			e, sent := ec.BulkAddProducts(scrapedProducts, ctx, start)
			if e == nil && sent {
				core.InfoLogger.Printf("Miscellaneous sleeping for %s seconds...", strconv.Itoa(config.GlobalDelay))
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

func scrapeWHSmith(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetWhs(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func scrapeSelfridges(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetSelf(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func scrapeJohnLewis(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetJL(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func scrapeArgos(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetArg(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func scrapeToysRUs(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetToys(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}
