package scrapers

import (
	"encoding/json"
	"io"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/core/http"
	"johgo-search-engine/internal/scrapers/shopify/pkg/models"
	"strings"
	"time"
)

func GetPages(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	anotherPage := true
	retries := 0

	for anotherPage {

		cli := http.ScraperHttpclient("")

		getShopifyReq, err := http.BuildShopifyRequest(site.URL, currentPage)

		if err == nil {

			shopifyResponse, err := cli.Do(getShopifyReq)

			if err == nil {

				if shopifyResponse.StatusCode == 200 {

					defer shopifyResponse.Body.Close()

					bodyBytes, err := io.ReadAll(shopifyResponse.Body)
					if err != nil {
						core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
						return pulledProducts, err, site.Name
					}

					var productStruct models.ShopifyProducts

					err = json.Unmarshal(bodyBytes, &productStruct)
					if err == nil {
						if !(len(productStruct.Products) == 0) {
							for i := 0; i < len(productStruct.Products); i++ {

								if strings.Contains(strings.ToLower(productStruct.Products[i].Title), strings.ToLower(site.Keyword)) {

									if !(len(productStruct.Products[i].Variants) == 0) {

										for v := 0; v < len(productStruct.Products[i].Variants); v++ {

											if productStruct.Products[i].Variants[v].Available {

												productStorageModel := formatToStorageModel(productStruct, i, v, site)

												pulledProducts.Products = append(pulledProducts.Products, productStorageModel)
											}
										}
									}
								}
							}
						} else {
							anotherPage = false
						}

					} else {
						core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
						return pulledProducts, err, site.Name
					}
					currentPage++
				} else {
					retries++
					if !core.CheckRetries(retries) {
						//core.ErrorLogger.Printf("Retries exceeded on: %s, returning products...", site.Name)
						return pulledProducts, err, site.Name
					}
					time.Sleep(7 * time.Second)
				}

			} else {
				if strings.Contains(err.Error(), "context deadline exceeded") {
					core.ErrorLogger.Printf("timeout error (context deadline exceeded) on site: ", site.Name)
				} else {
					core.ErrorLogger.Printf("error getting response: ", site.Name, err.Error())

					anotherPage = false
				}
			}

		} else {
			core.ErrorLogger.Printf("Error creating request: ", site.Name, err.Error())
			anotherPage = false
		}
	}

	return pulledProducts, nil, site.Name

}

func formatToStorageModel(productStruct models.ShopifyProducts, i int, v int, site coreModels.Site) elastic.ElasticProduct {

	productStorageModel := elastic.ElasticProduct{}

	if len(productStruct.Products[i].Images) == 0 {
		productStorageModel.Image = config.DefaultImage
	} else {
		productStorageModel.Image = productStruct.Products[i].Images[0].Src
	}
	productStorageModel.Price = productStruct.Products[i].Variants[v].Price
	productStorageModel.Title = productStruct.Products[i].Title
	productStorageModel.Url = site.URL + "/products/" + productStruct.Products[i].Handle
	productStorageModel.SiteName = site.Name
	productStorageModel.SiteUrl = site.URL
	return productStorageModel
}
