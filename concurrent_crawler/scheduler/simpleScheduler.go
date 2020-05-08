package concurrent_scheduler

import concurrent_engine "crawler/concurrent_crawler/engine"

type SimpleConcurrent struct {
	WorkerChan chan concurrent_engine.Request
}

func (s *SimpleConcurrent) Submit(r concurrent_engine.Request) {
	go func() {
		s.WorkerChan <- r
	}()
}

func (s *SimpleConcurrent) ConfigureMasterWorkerChan(in chan concurrent_engine.Request) {
	s.WorkerChan = in
}
