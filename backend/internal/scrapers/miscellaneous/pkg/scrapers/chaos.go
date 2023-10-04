package scrapers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/core/http"
	"time"
)

func GetChaos(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from:", site.Name)

	pageUris := []string{"collection-boxes-pokemon", "booster-boxes-pokemon", "elite-trainer-boxes-pokemon",
		"gift-tins-pokemon", "theme-decks-pokemon", "other-pokemon", "booster-packs-pokemon"}
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0

	for _, uri := range pageUris {

		cli := http.ScraperHttpclient("")

		getChaos, err := http.BuildChaosRequest(uri)

		if err == nil {

			chaosResponse, err := cli.Do(getChaos)

			if err == nil {

				if chaosResponse.StatusCode == 200 {

					defer chaosResponse.Body.Close()

					doc, err := goquery.NewDocumentFromReader(chaosResponse.Body)
					if err != nil {
						core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
						return pulledProducts, err, site.Name
					}

					if err == nil {

						doc.Find(".prod-list__element").Each(func(i int, s *goquery.Selection) {

							addToBasket := s.Find(".prod-el__quick-buy").Text()

							if addToBasket == "Add to basket" {

								productStorageModel := elastic.ElasticProduct{}

								href, exists := s.Find(".prod-el__link").Attr("href")
								if exists {
									productStorageModel.Url = fmt.Sprintf("%s%s", site.URL, href)
								} else {
									productStorageModel.Url = site.URL
								}

								title := s.Find(".prod-el__title span").Text()
								if title == "" {
									productStorageModel.Title = "No title"
								} else {
									productStorageModel.Title = title
								}

								imgSrc, exists := s.Find(".prod-el__image").First().Attr("src")
								if exists {
									productStorageModel.Image = imgSrc
								} else {
									productStorageModel.Image = config.DefaultImage
								}

								price := s.Find(".prod-el__pricing-price.prod-el__pricing-price--sale").Text()
								if price == "" {
									price = s.Find(".prod-el__pricing-price").Text()
									if price == "" {
										productStorageModel.Price = "0"
									} else {
										productStorageModel.Price = price
									}
								} else {
									productStorageModel.Price = price
								}
								productStorageModel.SiteName = site.Name
								productStorageModel.SiteUrl = site.URL
								pulledProducts.Products = append(pulledProducts.Products, productStorageModel)
							} else {

							}
						})

					} else {
						retries++
						core.ErrorLogger.Printf("Banned on: %s, retries left: %d", site.Name, retries)
						time.Sleep(7 * time.Second)
					}
				} else {
					retries++
					core.ErrorLogger.Printf("Banned on: %s, retries left: %d", site.Name, retries)
					time.Sleep(7 * time.Second)
				}

			} else {

				core.ErrorLogger.Printf("Error getting response: ", site.Name, err.Error())

			}

		} else {
			core.ErrorLogger.Printf("Error creating request: ", site.Name, err.Error())
		}
	}
	return pulledProducts, nil, site.Name
}
