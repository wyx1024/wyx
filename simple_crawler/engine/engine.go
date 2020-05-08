package simple_engine

import (
	"log"
)

func Run(seeds ...Request) error {
	var requestQ []Request
	for _, req := range seeds {
		requestQ = append(requestQ, req)
	}
	id := 0
	for len(requestQ) > 0 {
		r := requestQ[0]
		requestQ = requestQ[1:]

		result, err := Worker(r)
		if err != nil {
			continue
		}

		for _, item := range result.Items {
			id++
			log.Printf("id: %d City: %s", id, item)

		}
		requestQ = append(requestQ, result.Requests...)
	}
	return nil
}
