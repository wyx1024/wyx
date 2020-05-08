package simple_parser

import (
	simple_config "crawler/simple_crawler/config"
	simple_engine "crawler/simple_crawler/engine"
	"regexp"
)

var CarList = regexp.MustCompile(`<a href="(/car/[\d+-]+\d+/)" [^<]*><span [^<]*><img [^<]*></span>([^<]+)</a>`)

func ParserCarTypeList(contents []byte) simple_engine.ParserRestul {
	matchs := CarList.FindAllSubmatch(contents, -1)
	var result simple_engine.ParserRestul
	for _, m := range matchs {
		url := simple_config.URLPATH + string(m[1])
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, simple_engine.Request{
			Url:      url,
			ParseFun: ParserCar,
		})
	}
	return result
}
