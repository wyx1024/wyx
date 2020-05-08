package main

import (
	"crawler/distributed_crawler/config"
	"crawler/distributed_crawler/engine"
	"crawler/distributed_crawler/parser"
	ItemSaver "crawler/distributed_crawler/rpc/client/Item"
	"crawler/distributed_crawler/rpc/client/worker"
	"crawler/distributed_crawler/rpc/rpcSupper"
	"crawler/distributed_crawler/scheduler"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	Itemhost    = flag.String("item_host", "", "connection on host")
	Workerhosts = flag.String("worker_hosts", "", "connection on host")
)

func main() {
	flag.Parse()
	itemChan, err := ItemSaver.ItemSaverClient(*Itemhost)
	if err != nil {
		panic(err)
	}
	pool := CreateClientPool(strings.Split(*Workerhosts, ","))
	proccess := worker.CreateProccesser(pool)
	s := engine.SchedulerEngineQu{
		Scheduler:       &scheduler.Concurrent{},
		WorkerCount:     100,
		ItemChan:        itemChan,
		RequestProccess: proccess,
	}
	r := engine.Request{
		Url:   config.URL,
		Parse: engine.NewFuncParse(parser.ParserCarTypeList, config.ParseCarTypeList),
	}
	s.Run(r)
}
func CreateClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcSupper.NewClient(h)
		if err == nil {
			clients = append(clients, client)
		} else {
			log.Printf("create client err %s", err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
