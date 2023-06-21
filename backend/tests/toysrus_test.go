package tests

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
	"testing"
)

type ToysRUsResponse struct {
	Status string      `json:"status"`
	Path   interface{} `json:"path"`
	Pages  int         `json:"pages"`
	HTML   string      `json:"html"`
	Header string      `json:"header"`
}

func TestTRUScrape(t *testing.T) {

	url := "https://www.toysrus.co.uk/api/ajax_load_products.php"

	payload := strings.NewReader("action=autoload&category=1576&offset=1&pages=3&filter=&manufacturer=&order=popularity")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("authority", "www.toysrus.co.uk")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("dnt", "1")
	req.Header.Add("origin", "https://www.toysrus.co.uk")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://www.toysrus.co.uk/pokemon.html")
	req.Header.Add("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `macOS`)
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	resp, _ := http.DefaultClient.Do(req)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
	}

	var productStruct ToysRUsResponse

	err = json.Unmarshal(bodyBytes, &productStruct)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(productStruct.HTML))
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".product-item").Each(func(i int, s *goquery.Selection) {
		// Extract product link
		productLink, exists := s.Find("a").Attr("href")
		if exists {
			fmt.Printf("Product Link: %s\n", productLink)
		}

		// Extract product name
		productName := strings.TrimSpace(s.Find(".h3").Text())
		if len(productName) > 0 {
			fmt.Printf("Product Name: %s\n", productName)
		}

		// Extract product price
		productPrice := strings.TrimSpace(s.Find(".new-price").Text())
		if len(productPrice) > 0 {
			fmt.Printf("Product Price: %s\n", strings.ReplaceAll(productPrice, "£", ""))
		}

		// Extract product image
		productImage, exists := s.Find("img").Attr("src")
		if exists {
			fmt.Printf("Product Image: %s\n", productImage)
		}

		fmt.Println()
	})

}
