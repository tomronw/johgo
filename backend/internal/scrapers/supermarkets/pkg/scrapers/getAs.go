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
	"johgo-search-engine/internal/scrapers/supermarkets/pkg/models"
	"strings"
	"time"
)

func GetAs(site coreModels.Site) (p elastic.ProductsToStore, err error, s string) {

	core.InfoLogger.Println("Getting products from: ", site.Name)

	currentPage := 1
	var ps []elastic.ElasticProduct
	pulledProducts := elastic.ProductsToStore{Products: ps}
	retries := 0

	cli := http.ScraperHttpclient("")

	getAsReq, err := http.BuildAsdaRequest()

	if err == nil {

		asResponse, err := cli.Do(getAsReq)

		if err == nil {

			if asResponse.StatusCode == 200 {

				defer asResponse.Body.Close()

				bodyBytes, err := io.ReadAll(asResponse.Body)
				if err != nil {
					core.ErrorLogger.Printf("Error parsing body [%s], returning products: %s", site.Name, err)
					return pulledProducts, err, site.Name
				}

				var productStruct models.AsProducts

				err = json.Unmarshal(bodyBytes, &productStruct)
				if err == nil {
					// Check if Zones has at least 2 elements before accessing Zones[1]
					if len(productStruct.Data.TempoCmsContent.Zones) > 1 {
						zone := productStruct.Data.TempoCmsContent.Zones[1]

						// Check if there are items in the zone
						if len(zone.Configs.Products.Items) > 0 {
							for _, item := range zone.Configs.Products.Items {
								if strings.Contains(strings.ToLower(item.Item.Name), strings.ToLower(site.Keyword)) {
									if item.Inventory.AvailabilityInfo.Availability == "A" {
										productStorageModel := elastic.ElasticProduct{}

										if len(item.Item.Images.Scene7ID) == 0 {
											productStorageModel.Image = config.DefaultImage
										} else {
											productStorageModel.Image = fmt.Sprintf("https://ui.assets-asda.com/dm/asdagroceries/%s", item.Item.Images.Scene7ID)
										}

										productStorageModel.Price = strings.ReplaceAll(item.Price.PriceInfo.Price, "£", "")
										productStorageModel.Title = item.Item.Name
										productStorageModel.Url = fmt.Sprintf("https://groceries.asda.com/product/%s", item.Item.SkuID)
										productStorageModel.SiteName = site.Name
										productStorageModel.SiteUrl = site.URL

										pulledProducts.Products = append(pulledProducts.Products, productStorageModel)
									}
								}
							}
						} else {
							// Handle the case where there are no items
							return pulledProducts, nil, site.Name
						}
					} else {
						// Handle the case where Zones does not have enough elements
						return pulledProducts, fmt.Errorf("not enough elements in Zones"), site.Name
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
