package concurrent_engine

type SchedulerQu interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	WorkerReader(chan Request)
}
type SchedulerEngineQu struct {
	Scheduler   SchedulerQu
	WorkerCount int
	ItemChan    chan Item
}

func (s SchedulerEngineQu) Run(seeds ...Request) {
	out := make(chan ParserRestul)
	s.Scheduler.Run()

	for i := 0; i < s.WorkerCount; i++ {
		createWorkerQu(s.Scheduler.WorkerChan(), out, s.Scheduler)
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
			s.Scheduler.Submit(req)
		}
	}
}

func createWorkerQu(in chan Request, out chan ParserRestul, s SchedulerQu) {
	go func() {
		for {
			s.WorkerReader(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
