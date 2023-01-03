package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://finance.yahoo.co.jp/quote/2685.T")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("._1-yujUee").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("span").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

func main() {
	ExampleScrape()
}
