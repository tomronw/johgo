package miscellaneous

import (
	"context"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/scrapers/miscellaneous/pkg"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

// maps site name to function
var scrapeFuncs = map[string]func(coreModels.Site, chan<- elastic.IndexChannel){
	"WHSmith":       pkg.ScrapeWHSmith,
	"Selfridges":    pkg.ScrapeSelfridges,
	"John Lewis":    pkg.ScrapeJohnLewis,
	"Argos":         pkg.ScrapeArgos,
	"ToysRUs":       pkg.ScrapeToysRUs,
	"Chaos Cards":   pkg.ScrapeChaos,
	"Sports Direct": pkg.ScrapeSD,
}

// ScrapeMiscellaneous scrapes all miscellaneous sites
func ScrapeMiscellaneous() {
	// creates elastic client
	ec, _ := elastic.CreateClient("miscellaneous")

	for true {

		core.InfoLogger.Printf("scraping miscellaneous sites...")
		// gets all miscellaneous sites
		miscellaneousSites, err := core.GetStoreList("miscellaneous")

		start := time.Now()
		// if no error and sites are returned
		if err == nil || miscellaneousSites.Success {
			// create a channel for each site
			scrapedProducts := make(chan elastic.IndexChannel, len(miscellaneousSites.ReturnedSites.Sites))

			for _, site := range miscellaneousSites.ReturnedSites.Sites {
				// call map of functions and call if exists
				if scrapeFunc, exists := scrapeFuncs[site.Name]; exists {
					// add routine to wait group
					wg.Add(1)
					// spawn routine with correct function
					go func(site coreModels.Site, scrapeFunc func(coreModels.Site, chan<- elastic.IndexChannel)) {
						// defer wait group
						defer wg.Done()
						// call function
						scrapeFunc(site, scrapedProducts)
					}(site, scrapeFunc)
				}
			}
			wg.Wait()

			core.InfoLogger.Printf("waiting for all goroutines to finish...")
			ctx := context.Background()
			// send products to elastic
			e, sent := ec.BulkAddProducts(scrapedProducts, ctx, start)
			if e == nil && sent {
				core.InfoLogger.Printf("miscellaneous sleeping for %s seconds...", strconv.Itoa(config.GlobalDelay))
				time.Sleep(time.Duration(config.GlobalDelay) * time.Second)
			} else {
				core.ErrorLogger.Printf("error sending products to elastic: %s", e.Error())
				time.Sleep(time.Duration(config.GlobalDelay) * time.Second)
			}

		} else {
			core.ErrorLogger.Printf("error: [%s], please make sure apiCalls is live | Polling...", err.Error())
			time.Sleep(5 * time.Second)
		}

	}

}
