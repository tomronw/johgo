package tests

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"testing"
)

func TestWpScrape(t *testing.T) {

	res, err := http.Get("https://cardgalaxy.co.uk/page/1/?s=pokemon")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	print("scraping\n")

	doc.Find(".ast-grid-common-col.ast-full-width").Each(func(i int, s *goquery.Selection) {

		fmt.Println("here")

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

	})

}
