package worker

import (
	"crawler/distributed_crawler/engine"
	worker "crawler/distributed_crawler/rpc"
	"net/rpc"
)

func CreateProccesser(ClientChan chan *rpc.Client) engine.Proccesser {
	return func(r engine.Request) (engine.ParserRestul, error) {
		req := worker.SerializedRequest(r)
		var result worker.ParseResult
		c := <-ClientChan
		err := c.Call("CralService.Proccess", req, &result)
		if err != nil {
			return engine.ParserRestul{}, err
		}
		engineResult := worker.DeSerizlizedParseResult(result)
		return engineResult, nil
	}
}
