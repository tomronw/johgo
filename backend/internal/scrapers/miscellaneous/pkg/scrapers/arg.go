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

func GetArg(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0

	cli := http.ScraperHttpclient("")

	getArgReq, err := http.BuildArgosRequest(site.URL, currentPage)

	if err == nil {

		argResponse, err := cli.Do(getArgReq)

		if err == nil {

			if argResponse.StatusCode == 200 {

				defer argResponse.Body.Close()

				bodyBytes, err := io.ReadAll(argResponse.Body)
				if err != nil {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}

				var productStruct models.ArgProduct

				err = json.Unmarshal(bodyBytes, &productStruct)
				if err == nil {
					if !(len(productStruct.Data) == 0) {
						for i := 0; i < len(productStruct.Data); i++ {

							productStorageModel := elastic.ElasticProduct{}
							if len(productStruct.Data[i].Attributes.WcsID) == 0 {
								productStorageModel.Image = config.DefaultImage
							} else {
								productStorageModel.Image = fmt.Sprintf("https://media.4rgos.it/s/Argos/%s_R_SET", productStruct.Data[i].ID)
							}
							productStorageModel.Price = strconv.Itoa(int(productStruct.Data[i].Attributes.Price))
							productStorageModel.Title = productStruct.Data[i].Attributes.Name
							productStorageModel.Url = fmt.Sprintf("https://www.argos.co.uk/product/%s", productStruct.Data[i].ID)
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
