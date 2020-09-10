package scheduler

import "demo-project/crawler/engine"

type SimpleScheduler struct{
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {s.workerChan <- request}()
	
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
	s.workerChan = requests
}

