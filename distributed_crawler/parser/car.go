package parser

import (
	"bytes"
	"crawler/distributed_crawler/config"
	"crawler/distributed_crawler/engine"
	"io/ioutil"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ParserCar(contents []byte, _ string) engine.ParserRestul {
	body := ioutil.NopCloser(bytes.NewReader(contents))
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Printf("Parser  body err: %s", err)
	}
	var result engine.ParserRestul
	doc.Find("div.car_list_ul").Find("div.car_col2").Find("div.txt_list").Find("div.title").Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		url := config.URLPATH + href
		result.Requests = append(result.Requests, engine.Request{
			Url:   url,
			Parse: NewProfile(s.Text()),
		})

	})
	return result
}
