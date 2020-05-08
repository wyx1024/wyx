package main

import (
	concurrent_config "crawler/concurrent_crawler/config"
	concurrent_engine "crawler/concurrent_crawler/engine"
	concurrent_parser "crawler/concurrent_crawler/parser"
	concurrent_parsiti "crawler/concurrent_crawler/parsiti"
	concurrent_scheduler "crawler/concurrent_crawler/scheduler"
)

func main() {
	//1 simpleSchedluer 多个对一
	//s := concurrent_engine.SchedulerEngine{
	//	Scheduler:   &concurrent_scheduler.SimpleConcurrent{},
	//	WorkerCount: 100,
	//}

	//2 SchedluerQu  一对一队列
	itemChan, err := concurrent_parsiti.ItemSaver(concurrent_config.ELASTIC_INDEX)
	if err != nil {
		panic(err)
	}
	s := concurrent_engine.SchedulerEngineQu{
		Scheduler:   &concurrent_scheduler.SimpleConcurrentQ{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	r := concurrent_engine.Request{
		Url:      concurrent_config.URL,
		ParseFun: concurrent_parser.ParserCarTypeList,
	}

	s.Run(r)
}
