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
					if !(len(productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items) == 0) {
						for i := 0; i < len(productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items); i++ {

							if strings.Contains(strings.ToLower(productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items[i].Item.Name), strings.ToLower(site.Keyword)) {

								if productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items[i].Inventory.AvailabilityInfo.Availability == "A" {
									productStorageModel := elastic.ElasticProduct{}
									if len(productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items[i].Item.Images.Scene7ID) == 0 {
										productStorageModel.Image = config.DefaultImage
									} else {
										productStorageModel.Image = fmt.Sprintf("https://ui.assets-asda.com/dm/asdagroceries/%s", productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items[i].Item.Images.Scene7ID)
									}
									productStorageModel.Price = strings.ReplaceAll(productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items[i].Price.PriceInfo.Price, "£", "")
									productStorageModel.Title = productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items[i].Item.Name
									productStorageModel.Url = fmt.Sprintf("https://groceries.asda.com/product/%s", productStruct.Data.TempoCmsContent.Zones[1].Configs.Products.Items[i].Item.SkuID)
									productStorageModel.SiteName = site.Name
									productStorageModel.SiteUrl = site.URL
									pulledProducts.Products = append(pulledProducts.Products, productStorageModel)
								}
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
