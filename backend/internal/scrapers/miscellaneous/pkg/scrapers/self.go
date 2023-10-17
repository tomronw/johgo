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
	"strings"
	"time"
)

func GetSelf(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0

	cli := http.ScraperHttpclient("")

	getSelfReq, err := http.BuildSelfridgesRequest()

	if err == nil {

		selfResponse, err := cli.Do(getSelfReq)

		if err == nil {

			if selfResponse.StatusCode == 200 {

				defer selfResponse.Body.Close()

				bodyBytes, err := io.ReadAll(selfResponse.Body)
				if err != nil {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}

				var productStruct models.SelfProducts

				err = json.Unmarshal(bodyBytes, &productStruct)
				if err == nil {
					if !(len(productStruct.Products) == 0) {
						for i := 0; i < len(productStruct.Products); i++ {

							productStorageModel := elastic.ElasticProduct{}
							if len(productStruct.Products[i].Im) == 0 {
								productStorageModel.Image = config.DefaultImage
							} else {
								productStorageModel.Image = fmt.Sprintf("https://images.selfridges.com/is/image/selfridges/%s?$PDP_M_ZOOM$", productStruct.Products[i].Im)
							}
							productStorageModel.Price = strings.ReplaceAll(productStruct.Products[i].P, "£", "")
							productStorageModel.Title = productStruct.Products[i].Na
							productStorageModel.Url = fmt.Sprintf("https://www.selfridges.com/GB/en/cat/prod_%s", productStruct.Products[i].ID)
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
