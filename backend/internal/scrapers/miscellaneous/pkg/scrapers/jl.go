package scrapers

import (
	"encoding/json"
	"fmt"
	"io"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/core/http"
	"johgo-search-engine/internal/scrapers/miscellaneous/pkg/models"
	"time"
)

func GetJL(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0

	cli := http.ScraperHttpclient("")

	getJLReq, err := http.BuildJohnLewisRequest()

	if err == nil {

		jLResponse, err := cli.Do(getJLReq)

		if err == nil {

			if jLResponse.StatusCode == 200 {

				defer jLResponse.Body.Close()

				bodyBytes, err := io.ReadAll(jLResponse.Body)
				if err != nil {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}

				var productStruct models.JLProductss

				err = json.Unmarshal(bodyBytes, &productStruct)
				if err == nil {
					if !(len(productStruct.JLProducts) == 0) {
						for i := 0; i < len(productStruct.JLProducts); i++ {
							if !productStruct.JLProducts[i].OutOfStock {

								productStorageModel := elastic.ElasticProduct{}
								productStorageModel.Image = core.ValidateString(productStruct.JLProducts[i].Image, config.DefaultImage)
								productStorageModel.Price = core.ValidateString(productStruct.JLProducts[i].Price.Now, "0.00")
								productStorageModel.Title = core.ValidateString(productStruct.JLProducts[i].Title, "error")
								productStorageModel.Url = fmt.Sprintf("https://www.johnlewis.com/product/p%s", core.ValidateString(productStruct.JLProducts[i].ProductID, "404"))
								productStorageModel.SiteName = site.Name
								productStorageModel.SiteUrl = site.URL
								pulledProducts.Products = append(pulledProducts.Products, productStorageModel)

							}

						}
					} else {
						//core.InfoLogger.Printf("No more items for: %s ", site.Name)
						return pulledProducts, err, site.Name
					}

				} else {
					//core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}
				// fmt.Println("Page: ", strconv.Itoa(currentPage), site.Name)
				currentPage++
			} else {
				retries++
				//core.ErrorLogger.Printf("Banned on: %s, retries left: %d", site.Name, retries)
				if !core.CheckRetries(retries) {
					//core.ErrorLogger.Printf("Retries exceeded on: %s, returning products...", site.Name)
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
