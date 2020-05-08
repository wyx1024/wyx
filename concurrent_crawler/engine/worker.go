package concurrent_engine

import (
	concurrent_fetcher "crawler/concurrent_crawler/fetcher"
	"log"
)

func Worker(r Request) (ParserRestul, error) {
	ioResult, err := concurrent_fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher Url:%s,err:%s", r.Url, err)
		return ParserRestul{}, err
	}
	return r.ParseFun(ioResult, r.Url), nil
}
