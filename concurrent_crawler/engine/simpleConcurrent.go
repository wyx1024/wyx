package concurrent_engine

import (
	"log"
)

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}
type SchedulerEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (s SchedulerEngine) Run(seeds ...Request) {
	out := make(chan ParserRestul)
	in := make(chan Request)

	s.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < s.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, req := range seeds {
		s.Scheduler.Submit(req)
	}

	id := 0
	for {
		result := <-out
		for _, item := range result.Items {
			id++
			log.Printf("id: %d City: %s", id, item)
		}
		for _, req := range result.Requests {
			s.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request, out chan ParserRestul) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
