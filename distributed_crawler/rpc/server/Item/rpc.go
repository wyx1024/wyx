package main

import (
	"crawler/distributed_crawler/config"
	"crawler/distributed_crawler/rpc/rpcSupper"
	"crawler/distributed_crawler/rpc/service"
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

var port = flag.Int("itemsaver_host", 0, "Itemsaveer_server_host")

func main() {
	flag.Parse()
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = rpcSupper.ServerRpc(fmt.Sprintf(":%d", *port), service.ItemSaverService{
		Client: client,
		Index:  config.ELASTIC_INDEX,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("connection on: %d", *port)

}
