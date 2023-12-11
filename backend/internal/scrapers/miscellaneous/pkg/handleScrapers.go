package pkg

import (
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/scrapers/miscellaneous/pkg/scrapers"
)

// handles all miscellaneous scrapers

func ScrapeWHSmith(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetWhs(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func ScrapeSelfridges(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetSelf(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func ScrapeJohnLewis(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetJL(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func ScrapeArgos(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetArg(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func ScrapeToysRUs(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetToys(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func ScrapeChaos(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetChaos(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}

func ScrapeSD(site coreModels.Site, scrapedProducts chan<- elastic.IndexChannel) {

	productsFound, err, store := scrapers.GetSd(site)
	siteResults := elastic.IndexChannel{
		SiteName:      store,
		ReturnProduct: productsFound,
		Error:         err,
	}

	scrapedProducts <- siteResults
}
