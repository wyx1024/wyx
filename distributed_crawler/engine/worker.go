package engine

import (
	"crawler/distributed_crawler/fetcher"
	"log"
)

func Worker(r Request) (ParserRestul, error) {
	ioResult, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher Url:%s,err:%s", r.Url, err)
		return ParserRestul{}, err
	}
	return r.Parse.Parse(ioResult, r.Url), nil
}
