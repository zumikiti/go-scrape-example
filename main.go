package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	price float64
	per   float64
	pbr   float64
}

func FindValue(doc *goquery.Document) float64 {
	var value float64

	doc.Find("._38iJU1zx").Each(func(i int, s *goquery.Selection) {
		s.Each(func(i int, ss *goquery.Selection) {
			sss := ss.Find("span")
			t := sss.First().Text()

			if t == "前日終値" {
				v := strings.Replace(sss.Eq(1).Text(), ",", "", -1)

				value, _ = strconv.ParseFloat(v, 32)
			}
		})
	})

	return value
}

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
	var result Result

	result.price = FindValue(doc)

	fmt.Printf("Result: %f\n", result.price)
}

func main() {
	ExampleScrape()
}
