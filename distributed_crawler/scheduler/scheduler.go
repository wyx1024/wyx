package scheduler

import "crawler/distributed_crawler/engine"

type Concurrent struct {
	RequestChan chan engine.Request
	WokreChan   chan chan engine.Request
}

func (c *Concurrent) Run() {
	c.RequestChan = make(chan engine.Request)
	c.WokreChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeR engine.Request
			var activeW chan engine.Request
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

func (c *Concurrent) WorkerReader(w chan engine.Request) {
	c.WokreChan <- w
}

func (c *Concurrent) Submit(r engine.Request) {
	c.RequestChan <- r
}

func (c *Concurrent) WorkerChan() chan engine.Request {
	return c.RequestChan
}
