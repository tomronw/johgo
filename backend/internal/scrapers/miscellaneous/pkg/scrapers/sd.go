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

func GetSd(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0

	cli := http.ScraperHttpclient("")

	getSdReq, err := http.BuildSDReq()

	if err == nil {

		sdResponse, err := cli.Do(getSdReq)

		if err == nil {

			if sdResponse.StatusCode == 200 {

				defer sdResponse.Body.Close()

				bodyBytes, err := io.ReadAll(sdResponse.Body)
				if err != nil {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}

				var productStruct models.SDResModel

				err = json.Unmarshal(bodyBytes, &productStruct)
				if err == nil {
					if !(len(productStruct.Products) == 0) {
						for i := 0; i < len(productStruct.Products); i++ {
							productStorageModel := elastic.ElasticProduct{}
							productStorageModel.Image = core.ValidateString(productStruct.Products[i].Image, config.DefaultImage)
							productStorageModel.Price = fmt.Sprintf("%.2f", core.ValidateFloat64(productStruct.Products[i].PriceUnFormatted, 0.00))
							productStorageModel.Title = core.ValidateString(productStruct.Products[i].Name, "name not found")
							productStorageModel.Url = fmt.Sprintf("https://www.sportsdirect.com/%s", core.ValidateString(productStruct.Products[i].URL, "404"))
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
				currentPage++
			} else {
				retries++
				if !core.CheckRetries(retries) {
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
