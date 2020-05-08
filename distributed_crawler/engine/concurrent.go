package engine

import (
	"log"
)

type SchedulerQu interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	WorkerReader(chan Request)
}
type SchedulerEngineQu struct {
	Scheduler       SchedulerQu
	WorkerCount     int
	ItemChan        chan Item
	RequestProccess Proccesser
}
type Proccesser func(Request) (ParserRestul, error)

func (s SchedulerEngineQu) Run(seeds ...Request) {
	out := make(chan ParserRestul)
	s.Scheduler.Run()

	for i := 0; i < s.WorkerCount; i++ {
		s.createWorkerQu(s.Scheduler.WorkerChan(), out, s.Scheduler)
	}

	for _, req := range seeds {
		s.Scheduler.Submit(req)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(i Item) {
				s.ItemChan <- i
			}(item)
		}
		for _, req := range result.Requests {
			if isurl := IsURL(req.Url); isurl {
				log.Printf("Fectch url ISRUL %v", isurl)
			}
			s.Scheduler.Submit(req)
		}
	}
}

var URLBool = make(map[string]bool)

func IsURL(url string) bool {
	if URLBool[url] {
		return true
	}
	URLBool[url] = true
	return false
}

func (s SchedulerEngineQu) createWorkerQu(in chan Request, out chan ParserRestul, e SchedulerQu) {
	go func() {
		for {
			e.WorkerReader(in)
			request := <-in
			result, err := s.RequestProccess(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
