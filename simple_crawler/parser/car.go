package simple_parser

import (
	"bytes"
	simple_config "crawler/simple_crawler/config"
	simple_engine "crawler/simple_crawler/engine"
	"io/ioutil"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ParserCar(contents []byte) simple_engine.ParserRestul {
	body := ioutil.NopCloser(bytes.NewReader(contents))
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Printf("Parser  body err: %s", err)
	}
	var result simple_engine.ParserRestul
	doc.Find("div.car_list_ul").Find("div.car_col2").Find("div.txt_list").Find("div.title").Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		name := s.Text()
		url := simple_config.URLPATH + href
		//result.Items = append(result.Items, s.Text())
		result.Requests = append(result.Requests, simple_engine.Request{
			Url: url,
			ParseFun: func(i []byte) simple_engine.ParserRestul {
				return ParserProfile(i, name)
			},
		})

	})
	return result
}
