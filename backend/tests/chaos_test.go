package tests

import (
	"fmt"
	"johgo-search-engine/internal/core/coreModels"
	"johgo-search-engine/internal/scrapers/miscellaneous/pkg/scrapers"
	"testing"
)

func TestChaosScrape(t *testing.T) {

	products, err, _ := scrapers.GetChaos(coreModels.Site{
		Name:    "Chaos Cards",
		URL:     "https://www.chaoscards.co.uk",
		ISO:     "GB",
		Keyword: "pokemon",
	})

	if err != nil {
		t.Errorf("failed to get chaos")
	} else {
		for _, product := range products.Products {
			fmt.Println(product.Url)
			fmt.Println(product.Title)
			fmt.Println(product.Price)
			fmt.Println(product.Image)
		}
	}

}
