package scheduler

import "demo-project/crawler/engine"

type SimpleScheduler struct{
	workerChan chan engine.Request
}

//Submit
func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {s.workerChan <- request}()
	
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request  {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(c chan engine.Request)  {
	//s.workerChan <- c
}

func (s *SimpleScheduler) Run()  {
	s.workerChan = make(chan engine.Request)
}

//func (s *SimpleScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
//	s.workerChan = requests
//}

