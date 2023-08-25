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
	"strconv"
	"time"
)

func GetWhs(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0

	cli := http.ScraperHttpclient("")

	getWhsReq, err := http.BuildWhsmithRequest(site.URL, currentPage)

	if err == nil {

		whsResponse, err := cli.Do(getWhsReq)

		if err == nil {

			if whsResponse.StatusCode == 200 {

				defer whsResponse.Body.Close()

				bodyBytes, err := io.ReadAll(whsResponse.Body)
				if err != nil {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}

				var productStruct models.WHSProducts

				err = json.Unmarshal(bodyBytes, &productStruct)
				if err == nil {
					if !(len(productStruct.Hits) == 0) {
						for i := 0; i < len(productStruct.Hits); i++ {

							productStorageModel := elastic.ElasticProduct{}
							if productStruct.Hits[i].CImages[0].Alt != nil {
								productStorageModel.Image = productStruct.Hits[i].CImages[0].URL
							} else {
								productStorageModel.Image = config.DefaultImage
							}
							productStorageModel.Price = strconv.Itoa(int(productStruct.Hits[i].Price))
							productStorageModel.Title = productStruct.Hits[i].ProductName
							productStorageModel.Url = fmt.Sprintf("https://www.whsmith.co.uk/%s/%s.html", productStruct.Hits[i].CPageURL, productStruct.Hits[i].ProductID)
							productStorageModel.SiteName = site.Name
							productStorageModel.SiteUrl = site.URL
							pulledProducts.Products = append(pulledProducts.Products, productStorageModel)

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
				core.ErrorLogger.Printf("Banned on: %s, retries left: %d", site.Name, retries)
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
