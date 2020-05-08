package concurrent_parser

import (
	"bytes"
	concurrent_engine "crawler/concurrent_crawler/engine"
	concurrent_model "crawler/concurrent_crawler/model"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var reg = regexp.MustCompile("\\s*")
var idRe = regexp.MustCompile(`.*/car/select/s([\d]+)/`)

func ParserProfile(contents []byte, name string, url string) concurrent_engine.ParserRestul {
	body := ioutil.NopCloser(bytes.NewReader(contents))
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Printf("Parser  body err: %s", err)
	}
	var cardetail concurrent_model.Car
	cardetail.Name = name
	id := idRe.FindSubmatch([]byte(url))
	Cardate(doc, &cardetail)
	if cardetail.Price == "" {
		cardetail.Price = "暂无"
	}
	var result concurrent_engine.ParserRestul
	result.Items = append(result.Items, concurrent_engine.Item{
		Url:     url,
		Id:      string(id[1]),
		Payload: cardetail,
	})
	return result
}

func Cardate(doc *goquery.Document, cardetail *concurrent_model.Car) {
	doc.Find("div.price_menu_box").Each(func(i int, s *goquery.Selection) {
		value := s.Text()
		value = RepString(value)
		value = strings.Replace(value, "厂商指导价：", "", -1)
		cardetail.Price = value
	})

	docs := doc.Find("div.data_list").Find("ul").Find("li")

	docs.Eq(0).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)

		cardetail.Level = value
	})
	docs.Eq(1).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Structure = value
	})
	docs.Eq(2).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Fuel = value
	})
	docs.Eq(3).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Displacement = value
	})
	docs.Eq(4).Each(func(i int, s *goquery.Selection) {
		value := s.After("label").Text()
		value = RepString(value)
		cardetail.Guarantee = value
	})
	docs.Eq(5).Each(func(i int, s *goquery.Selection) {
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
