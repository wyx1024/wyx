package concurrent_parsiti

import (
	"context"
	concurrent_engine "crawler/concurrent_crawler/engine"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan concurrent_engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan concurrent_engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("item: %d  Car: %s", itemCount, item)
			_, err := client.Index().Index(index).Id(item.Id).BodyJson(item).Do(context.Background())
			if err != nil {
				log.Printf("Item saver :error:saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
