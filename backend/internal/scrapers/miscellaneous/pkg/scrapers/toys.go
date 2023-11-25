package scrapers

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/core/http"
	"johgo-search-engine/internal/scrapers/miscellaneous/pkg/models"
	"strings"
	"time"
)

func GetToys(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	anotherPage := true
	retries := 0

	for anotherPage {

		cli := http.ScraperHttpclient("")

		getToysReq, err := http.BuildToysRUsRequest(currentPage)

		if err == nil {

			toysResponse, err := cli.Do(getToysReq)

			if err == nil {

				if toysResponse.StatusCode == 200 {

					defer toysResponse.Body.Close()

					bodyBytes, err := io.ReadAll(toysResponse.Body)
					if err != nil {
					}

					var productStruct models.ToysResponse

					err = json.Unmarshal(bodyBytes, &productStruct)

					doc, err := goquery.NewDocumentFromReader(strings.NewReader(productStruct.HTML))
					if err != nil {
						core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
						return pulledProducts, err, site.Name
					}

					if err == nil {
						doc.Find(".product-item").Each(func(i int, s *goquery.Selection) {

							dupe := false

							productStorageModel := elastic.ElasticProduct{}

							productLink, exists := s.Find("a").Attr("href")
							if exists {
								productStorageModel.Url = productLink
							} else {
								productStorageModel.Url = site.URL
							}

							productName := strings.TrimSpace(s.Find(".h3").Text())
							if len(productName) > 0 {
								productStorageModel.Title = productName
							} else {
								productStorageModel.SiteName = site.Name
							}

							productPrice := strings.TrimSpace(s.Find(".new-price").Text())
							if len(productPrice) > 0 {
								productStorageModel.Price = strings.ReplaceAll(productPrice, "£", "")
							} else {
								productStorageModel.Price = "??"
							}

							productImage, exists := s.Find("img").Attr("src")
							if exists {
								productStorageModel.Image = productImage
							} else {
								productStorageModel.Image = config.DefaultImage
							}
							productStorageModel.SiteName = site.Name
							productStorageModel.SiteUrl = site.URL

							for _, pulledProduct := range pulledProducts.Products {

								if pulledProduct.Url == productStorageModel.Url {
									dupe = true
								}
							}
							if !dupe {
								pulledProducts.Products = append(pulledProducts.Products, productStorageModel)
							}
						})
						currentPage++
						if currentPage == 4 {
							anotherPage = false
						}

					} else {
						anotherPage = false
						//core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
						return pulledProducts, err, site.Name
					}
					currentPage++
				} else {
					retries++
					///core.ErrorLogger.Printf("Banned on: %s, retries left: %d", site.Name, retries)
					if !core.CheckRetries(retries) {
						//core.ErrorLogger.Printf("Retries exceeded on: %s, returning products...", site.Name)
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
