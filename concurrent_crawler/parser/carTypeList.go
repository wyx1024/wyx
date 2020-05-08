package concurrent_parser

import (
	concurrent_config "crawler/concurrent_crawler/config"
	concurrent_engine "crawler/concurrent_crawler/engine"
	"regexp"
)

var CarList = regexp.MustCompile(`<a href="(/car/[\d+-]+\d+/)" [^<]*><span [^<]*><img [^<]*></span>([^<]+)</a>`)

func ParserCarTypeList(contents []byte, _ string) concurrent_engine.ParserRestul {
	matchs := CarList.FindAllSubmatch(contents, -1)
	var result concurrent_engine.ParserRestul
	for _, m := range matchs {
		url := concurrent_config.URLPATH + string(m[1])
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, concurrent_engine.Request{
			Url:      url,
			ParseFun: ParserCar,
		})
	}
	return result
}
