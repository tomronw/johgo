package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"johgo-search-engine/config"
	"johgo-search-engine/internal/core/coreModels"
	"net/http"
)

// To get site list from API, will return array of json values with site name, url and ISO

var (
	Shopify = ApiStore{
		ECommerce:        "Shopify",
		ApiRoute:         "shopify",
		ProductsEndpoint: config.ShopifyEP,
		IndexName:        "shopify",
	}

	Woocommerce = ApiStore{
		ECommerce:        "Woocommerce",
		ApiRoute:         "woocommerce",
		ProductsEndpoint: config.WooEP,
		IndexName:        "woocommerce",
	}
)

type ApiStore struct {
	ECommerce        string
	ApiRoute         string
	ProductsEndpoint string
	IndexName        string
}
type ApiResponse struct {
	ReturnedSites coreModels.SiteList
	Success       bool
}

func GetStoreList(site string) (ApiResponse, error) {

	url := fmt.Sprintf("%s/v1/stores/%s", config.SitesAPIURL, site)

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err == nil {
		defer res.Body.Close()

		if res.StatusCode != 200 {

			return ApiResponse{
				ReturnedSites: coreModels.SiteList{Sites: nil},
				Success:       false,
			}, errors.New("failed getting API, please check status of it")

		}
		body, _ := io.ReadAll(res.Body)

		var siteList coreModels.SiteList

		err = json.Unmarshal(body, &siteList)

		if err == nil {

			response := ApiResponse{
				ReturnedSites: siteList,
				Success:       true,
			}

			return response, nil
		} else {

			return ApiResponse{
				ReturnedSites: coreModels.SiteList{Sites: nil},
				Success:       false,
			}, err
		}
	} else {

		return ApiResponse{
			ReturnedSites: coreModels.SiteList{Sites: nil},
			Success:       false,
		}, err
	}
}
