package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	code  int
	price float64
	per   float64
	pbr   float64
}

func FindValue(doc *goquery.Document, key string) float64 {
	var value float64

	doc.Find("._38iJU1zx").Each(func(i int, s *goquery.Selection) {
		s.Each(func(i int, ss *goquery.Selection) {
			sss := ss.Find("span")
			t := sss.First().Text()

			if t == key {
				var v string

				// key 別に値を取得
				switch key {
				case "前日終値":
					v = strings.Replace(sss.Eq(1).Text(), ",", "", -1)
				case "PER", "PBR":
					v = sss.Eq(3).Text()
					v = strings.Replace(v, "(連)", "", -1)
					v = strings.Replace(v, "倍", "", -1)
				}

				value, _ = strconv.ParseFloat(v, 32)
			}
		})
	})

	return value
}

func ExampleScrape(code string) {
	// Request the HTML page.
	url := fmt.Sprintf("https://finance.yahoo.co.jp/quote/%s.T", code)
	res, err := http.Get(url)
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
	result.code, _ = strconv.Atoi(code)
	result.price = FindValue(doc, "前日終値")
	result.per = FindValue(doc, "PER")
	result.pbr = FindValue(doc, "PBR")

	fmt.Printf("Result: %f\n", result.price)
	fmt.Printf("Result: %f\n", result.per)
	fmt.Printf("Result: %f\n", result.pbr)

	saveData(result)
}

func main() {
	// 第一引数を取得する
	flag.Parse()
	code := flag.Arg(0)
	if code == "" {
		log.Fatal("第一引数に企業コードを指定してください")
	}

	ExampleScrape(code)
}
