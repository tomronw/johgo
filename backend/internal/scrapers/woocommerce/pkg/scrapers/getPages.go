package scrapers

import (
	"github.com/PuerkitoBio/goquery"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/core/http"
	"strings"
	"time"
)

func ScrapeWooCommerceSite(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	anotherPage := true
	retries := 0

	for anotherPage {

		cli := http.ScraperHttpclient("")

		getShopifyReq, err := http.BuildWooRequest(site.URL, currentPage)

		if err == nil {

			wooResponse, err := cli.Do(getShopifyReq)

			if err == nil {

				if wooResponse.StatusCode == 200 {

					defer wooResponse.Body.Close()

					doc, err := goquery.NewDocumentFromReader(wooResponse.Body)
					if err != nil {
						core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
						return pulledProducts, err, site.Name
					}

					if err == nil {
						it := 0 // iterator for products
						doc.Find(".product.type-product").Each(func(i int, s *goquery.Selection) {

							relAttr := s.Find("a[rel=nofollow]").Text()
							if strings.ToLower(relAttr) != "read more" {

								productStorageModel := elastic.ElasticProduct{}

								href, exists := s.Find(".woocommerce-LoopProduct-link.woocommerce-loop-product__link").Attr("href")
								if exists && strings.Contains(href, site.Keyword) {
									productStorageModel.Url = href

									title := s.Find(".woocommerce-loop-product__title").Text()

									if title == "" {
										productStorageModel.Title = "No title"
									} else {
										productStorageModel.Title = title
									}

									price := s.Find(".woocommerce-Price-amount.amount").Last().Text()
									if price == "" {
										productStorageModel.Price = "0"
									} else {
										productStorageModel.Price = strings.ReplaceAll(price, "£", "")
									}

									imgSrc, exists := s.Find("img.attachment-woocommerce_thumbnail").First().Attr("src")
									if exists {
										productStorageModel.Image = imgSrc
									} else {
										productStorageModel.Image = config.DefaultImage
									}

									productStorageModel.SiteName = site.Name
									productStorageModel.SiteUrl = site.URL
									pulledProducts.Products = append(pulledProducts.Products, productStorageModel)
								} else {
									core.ErrorLogger.Printf("Link not found on site: %s", site.Name)
									if it == 0 {
										anotherPage = false

									}
								}
							}
							it++
						})

					} else {
						anotherPage = false
						core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
						return pulledProducts, err, site.Name
					}
					currentPage++
				} else {
					retries++
					core.ErrorLogger.Printf("Banned on: %s, retries left: %d", site.Name, retries)
					if !core.CheckRetries(retries) {
						core.ErrorLogger.Printf("Retries exceeded on: %s, returning products...", site.Name)
						return pulledProducts, err, site.Name
					}
					time.Sleep(7 * time.Second)
				}

			} else {

				core.ErrorLogger.Printf("Error getting response: ", site.Name, err.Error())

				anotherPage = false
			}

		} else {
			core.ErrorLogger.Printf("Error creating request: ", site.Name, err.Error())
			anotherPage = false
		}
	}

	return pulledProducts, nil, site.Name

}
