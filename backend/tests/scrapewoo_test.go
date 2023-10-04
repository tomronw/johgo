package tests

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"testing"
)

func TestWooScrape(t *testing.T) {

	res, err := http.Get("https://bathtcg.co.uk/page/1/?s=pokemon&post_type=product")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".product.type-product").Each(func(i int, s *goquery.Selection) {

		href, exists := s.Find(".woocommerce-LoopProduct-link.woocommerce-loop-product__link").Attr("href")
		if exists {
			fmt.Println("Link source:", href)
		} else {
			fmt.Println("Link not found")
		}

		title := s.Find(".woocommerce-loop-product__title").Text()
		fmt.Println("title:", title)

		price := s.Find(".woocommerce-Price-amount.amount").Last().Text()
		fmt.Println("price:", price)

		imgSrc, exists := s.Find("img.attachment-woocommerce_thumbnail").First().Attr("src")
		if exists {
			fmt.Println("Image source:", imgSrc)
		} else {
			fmt.Println("Image source not found")
		}

		relAttr := s.Find("a[rel=nofollow]").Text()
		if exists {
			fmt.Println("rel value:", relAttr)
		} else {
			fmt.Println("rel attribute does not exist")
		}

		fmt.Println()
	})

}
