package ItemSaver

import (
	"crawler/distributed_crawler/engine"
	"crawler/distributed_crawler/model"
	"crawler/distributed_crawler/rpc/rpcSupper"
	"log"
)

func ItemSaverClient(host string) (chan engine.Item, error) {
	client, err := rpcSupper.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	var result bool

	go func() {
		itemCount := 0
		for {
			item := <-out
			car, _ := model.FormJsonObj(item.Payload)
			item.Payload = car
			itemCount++
			log.Printf("item: %d  Car: %v", itemCount, item)

			err = client.Call("ItemSaverService.Saver", item, &result)
			if err != nil {
				log.Printf("Item saver :error:saving item %v: %v", item, err)
			}
		}
	}()
	return out, err
}
