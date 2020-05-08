package parser

import (
	"crawler/distributed_crawler/config"
	"crawler/distributed_crawler/engine"
	"regexp"
)

var CarList = regexp.MustCompile(`<a href="(/car/[\d+-]+\d+/)" [^<]*><span [^<]*><img [^<]*></span>([^<]+)</a>`)

func ParserCarTypeList(contents []byte, _ string) engine.ParserRestul {
	matchs := CarList.FindAllSubmatch(contents, -1)
	var result engine.ParserRestul
	for _, m := range matchs {
		url := config.URLPATH + string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:   url,
			Parse: engine.NewFuncParse(ParserCar, config.ParseCar),
		})
	}
	return result
}
