package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	out := make(chan ParserResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()
	
	for i := 0; i < e.WorkerCount; i++{
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	
	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}
	
	for{
		result := <- out
		for _, item := range result.Items{
			//为每一个item开一个goroutine,消耗的速度比生成的速度快
			go func() {e.ItemChan <- item}()
		}
		
		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParserResult, ready ReadyNotifier)  {
	//in := make(chan Request)
	go func() {
		for{
			//tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}