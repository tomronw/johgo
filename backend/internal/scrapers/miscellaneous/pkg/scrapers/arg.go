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

func GetArg(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {
	// argos scraper
	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0
	// build client
	cli := http.ScraperHttpclient("")
	// return headers and url
	getArgReq, err := http.BuildArgosRequest()

	if err == nil {
		// make request
		argResponse, err := cli.Do(getArgReq)
		// if no errors
		if err == nil {
			// if 200
			if argResponse.StatusCode == 200 {
				// close body
				defer argResponse.Body.Close()
				// read body
				bodyBytes, err := io.ReadAll(argResponse.Body)
				if err != nil {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}
				// unmarshal body
				var productStruct models.ArgProduct

				err = json.Unmarshal(bodyBytes, &productStruct)
				if err == nil {
					// check if len of data is 0
					if !(len(productStruct.Data) == 0) {
						for i := 0; i < len(productStruct.Data); i++ {
							// loop through data and append to products
							productStorageModel := elastic.ElasticProduct{}
							productStorageModel.Image = core.ValidateString(productStruct.Data[i].Attributes.WcsID, config.DefaultImage)
							productStorageModel.Price = fmt.Sprintf("%.2f", core.ValidateFloat64(productStruct.Data[i].Attributes.Price, 0.00))
							productStorageModel.Title = core.ValidateString(productStruct.Data[i].Attributes.Name, "error")
							productStorageModel.Url = fmt.Sprintf("https://www.argos.co.uk/product/%s", core.ValidateString(productStruct.Data[i].ID, "404"))
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
				defer argResponse.Body.Close()
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
