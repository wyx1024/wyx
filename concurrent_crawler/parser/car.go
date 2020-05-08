package concurrent_parser

import (
	"bytes"
	concurrent_config "crawler/concurrent_crawler/config"
	concurrent_engine "crawler/concurrent_crawler/engine"
	"io/ioutil"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ParserCar(contents []byte, _ string) concurrent_engine.ParserRestul {
	body := ioutil.NopCloser(bytes.NewReader(contents))
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Printf("Parser  body err: %s", err)
	}
	var result concurrent_engine.ParserRestul
	doc.Find("div.car_list_ul").Find("div.car_col2").Find("div.txt_list").Find("div.title").Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		url := concurrent_config.URLPATH + href
		result.Requests = append(result.Requests, concurrent_engine.Request{
			Url:      url,
			ParseFun: ParseCarFunc(s.Text()),
		})

	})
	return result
}

func ParseCarFunc(name string) concurrent_engine.ParseFun {
	return func(i []byte, url string) concurrent_engine.ParserRestul {
		return ParserProfile(i, name, url)
	}
}
