package concurrent_scheduler

import concurrent_engine "crawler/concurrent_crawler/engine"

type SimpleConcurrentQ struct {
	WorkerChan2 chan concurrent_engine.Request
}

func (s *SimpleConcurrentQ) Submit(r concurrent_engine.Request) {
	go func() {
		s.WorkerChan2 <- r
	}()

}

func (s *SimpleConcurrentQ) WorkerChan() chan concurrent_engine.Request {
	return s.WorkerChan2
}

func (s *SimpleConcurrentQ) Run() {
	s.WorkerChan2 = make(chan concurrent_engine.Request)
}

func (s *SimpleConcurrentQ) WorkerReader(r chan concurrent_engine.Request) {
}
