package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	out := make(chan ParserResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()
	
	for i := 0; i < e.WorkerCount; i++{
		createWorker(out, e.Scheduler)
	}
	
	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}
	
	itemCount := 0
	for{
		result := <- out
		for _, item := range result.Items{
			log.Printf(" Got item #%d: %v\n", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParserResult, s Scheduler)  {
	in := make(chan Request)
	go func() {
		for{
			//tell scheduler i'm ready
			s.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}