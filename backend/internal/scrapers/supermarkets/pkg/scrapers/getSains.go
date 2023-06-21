package scrapers

import (
	"encoding/json"
	"io"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/core/http"
	"johgo-search-engine/internal/scrapers/supermarkets/pkg/models"
	"strconv"
	"strings"
	"time"
)

func GetSains(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0

	cli := http.ScraperHttpclient("")

	getSainsreq, err := http.BuildSainsburysRequest(site.URL, currentPage)

	if err == nil {

		sainsResponse, err := cli.Do(getSainsreq)

		if err == nil {

			if sainsResponse.StatusCode == 200 {

				defer sainsResponse.Body.Close()

				bodyBytes, err := io.ReadAll(sainsResponse.Body)
				if err != nil {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}

				var productStruct models.SainsProduct

				err = json.Unmarshal(bodyBytes, &productStruct)
				if err == nil {
					if !(len(productStruct.Productsz) == 0) {
						for i := 0; i < len(productStruct.Productsz); i++ {

							if strings.Contains(strings.ToLower(productStruct.Productsz[i].Name), strings.ToLower(site.Keyword)) {

								if productStruct.Productsz[i].IsAvailable {
									productStorageModel := elastic.ElasticProduct{}
									if len(productStruct.Productsz[i].Image) == 0 {
										productStorageModel.Image = config.DefaultImage
									} else {
										productStorageModel.Image = productStruct.Productsz[i].Image
									}
									productStorageModel.Price = strconv.Itoa(int(productStruct.Productsz[i].RetailPrice.Price))
									productStorageModel.Title = productStruct.Productsz[i].Name
									productStorageModel.Url = productStruct.Productsz[i].FullURL
									productStorageModel.SiteName = site.Name
									productStorageModel.SiteUrl = site.URL
									pulledProducts.Products = append(pulledProducts.Products, productStorageModel)
								}
							}

						}
					} else {
						core.InfoLogger.Printf("No more items for: %s ", site.Name)
						return pulledProducts, err, site.Name
					}

				} else {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}
				// fmt.Println("Page: ", strconv.Itoa(currentPage), site.Name)
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

		}

	} else {
		core.ErrorLogger.Printf("Error creating request: ", site.Name, err.Error())
	}

	return pulledProducts, nil, site.Name

}
