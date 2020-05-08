package concurrent_scheduler

import concurrent_engine "crawler/concurrent_crawler/engine"

type ConcurrentQ struct {
	RequestChan chan concurrent_engine.Request
	WokreChan   chan chan concurrent_engine.Request
}

func (c *ConcurrentQ) Run() {
	c.RequestChan = make(chan concurrent_engine.Request)
	c.WokreChan = make(chan chan concurrent_engine.Request)
	go func() {
		var requestQ []concurrent_engine.Request
		var workerQ []chan concurrent_engine.Request
		for {
			var activeR concurrent_engine.Request
			var activeW chan concurrent_engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeR = requestQ[0]
				activeW = workerQ[0]
			}
			select {
			case r := <-c.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-c.WokreChan:
				workerQ = append(workerQ, w)
			case activeW <- activeR:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

func (c *ConcurrentQ) WorkerReader(w chan concurrent_engine.Request) {
	c.WokreChan <- w
}

func (c *ConcurrentQ) Submit(r concurrent_engine.Request) {
	c.RequestChan <- r
}

func (c *ConcurrentQ) WorkerChan() chan concurrent_engine.Request {
	return c.RequestChan
}
