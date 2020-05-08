package simple_parser

import (
	"bytes"
	simple_engine "crawler/simple_crawler/engine"
	"crawler/simple_crawler/model"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var reg = regexp.MustCompile("\\s*")

func ParserProfile(contents []byte, name string) simple_engine.ParserRestul {
	body := ioutil.NopCloser(bytes.NewReader(contents))
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Printf("Parser  body err: %s", err)
	}
	var cardetail model.Car
	cardetail.Name = name
	Cardate(doc, &cardetail)
	var result simple_engine.ParserRestul
	result.Items = append(result.Items, cardetail)
	return result
}

func Cardate(doc *goquery.Document, cardetail *model.Car) {
	doc.Find("div.data_ck dealer_price_box").Find("ul").Find("li").Each(func(i int, s *goquery.Selection) {
		value := s.After("span").Text()
		value = RepString(value)
		cardetail.Price = value
	})
	doc.Find("div.data_list").Find("ul").Find("li").Eq(0).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Level = value
	})
	doc.Find("div.data_list").Find("ul").Find("li").Eq(1).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Structure = value
	})
	doc.Find("div.data_list").Find("ul").Find("li").Eq(2).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Fuel = value
	})
	doc.Find("div.data_list").Find("ul").Find("li").Eq(3).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Displacement = value
	})
	doc.Find("div.data_list").Find("ul").Find("li").Eq(4).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Guarantee = value
	})
	doc.Find("div.data_list").Find("ul").Find("li").Eq(5).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Transmission = value
	})
}

func RepString(value string) string {
	value = strings.Replace(value, " ", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	return value
}
