package scrap

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zumikiti/go-scrap-example/src/db"
	"github.com/zumikiti/go-scrap-example/src/typefile"
)

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
	var result typefile.Result
	result.Code, _ = strconv.Atoi(code)
	result.Price = FindValue(doc, "前日終値")
	result.Per = FindValue(doc, "PER")
	result.Pbr = FindValue(doc, "PBR")

	fmt.Printf("Result: %f\n", result.Price)
	fmt.Printf("Result: %f\n", result.Per)
	fmt.Printf("Result: %f\n", result.Pbr)

	db.SaveData(result)
}

func Scrap() {
	// 第一引数を取得する
	flag.Parse()
	code := flag.Arg(0)
	if code == "" {
		log.Fatal("第一引数に企業コードを指定してください")
	}

	ExampleScrape(code)
}
