package simple_engine

import (
	simple_fetcher "crawler/simple_crawler/fetcher"
	"log"
)

func Worker(r Request) (ParserRestul, error) {
	ioResult, err := simple_fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher Url:%s,err:%s", r.Url, err)
		return ParserRestul{}, err
	}
	return r.ParseFun(ioResult), nil
}
