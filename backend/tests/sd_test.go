package tests

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"johgo-search-engine/config"
	"johgo-search-engine/elastic"
	"johgo-search-engine/internal/core"
	"johgo-search-engine/internal/scrapers/miscellaneous/pkg/models"
	"net/http"
	"strconv"
	"testing"
)

func TestSdScrape(t *testing.T) {

	url := "https://www.sportsdirect.com/product/search?categoryId=&page=1&productsPerPage=59&sortOption=rank&selectedFilters=ABRA%5EPokemon%7CWEBSTYLE%5ECard%20Games&isSearch=true&searchText=pokemon&columns=3&mobileColumns=2&clearFilters=false&pathName=%2Fsearchresults&searchTermCategory=&selectedCurrency=GBP&portalSiteId=12&searchCategory="

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:121.0) Gecko/20100101 Firefox/121.0")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.5")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Referer", "https://www.sportsdirect.com/searchresults?descriptionfilter=pokemon")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("TE", "trailers")

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		log.Fatal("bad response code: " + strconv.Itoa(res.StatusCode))
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

	var productStruct models.SDResModel

	err := json.Unmarshal(body, &productStruct)

	if err == nil {
		if !(len(productStruct.Products) == 0) {
			for i := 0; i < len(productStruct.Products); i++ {
				productStorageModel := elastic.ElasticProduct{}
				productStorageModel.Image = core.ValidateString(productStruct.Products[i].Image, config.DefaultImage)
				productStorageModel.Price = fmt.Sprintf("%.2f", core.ValidateFloat64(productStruct.Products[i].PriceUnFormatted, 0.00))
				productStorageModel.Title = core.ValidateString(productStruct.Products[i].Name, "name not found")
				productStorageModel.Url = fmt.Sprintf("https://www.sportsdirect.com/%s", core.ValidateString(productStruct.Products[i].URL, "404"))
				productStorageModel.SiteName = "SD test"
				productStorageModel.SiteUrl = "https://sportsdirect.com"
				//pulledProducts.Products = append(pulledProducts.Products, productStorageModel)

				fmt.Println(productStorageModel.Url)
				fmt.Println(productStorageModel.Title)
				fmt.Println(productStorageModel.Price)
				fmt.Println(productStorageModel.Image)
				fmt.Println(productStorageModel.SiteName)
				fmt.Println(productStorageModel.SiteUrl)

			}
		} else {
			log.Fatal("no products found")
		}

	} else {
		log.Fatal(err)
	}

}
