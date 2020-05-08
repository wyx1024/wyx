package service

import (
	"crawler/distributed_crawler/engine"
	worker "crawler/distributed_crawler/rpc"
	"log"
)

type CralService struct {
}

func (CralService) Proccess(req worker.Request, result *worker.ParseResult) error {
	engineReq, err := worker.DeSerializedRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = worker.SerizlizedParseResult(engineResult)
	log.Printf("%v", req.Url)
	return nil
}
