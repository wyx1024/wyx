package main

import (
	"crawler/distributed_crawler/rpc/rpcSupper"
	"crawler/distributed_crawler/rpc/service"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("worker_host", 0, "worker_server_host")

func main() {
	flag.Parse()
	err := rpcSupper.ServerRpc(fmt.Sprintf(":%d", *port), service.CralService{})
	if err != nil {
		panic(err)
	}
	log.Printf("connection on: %d", *port)
}
