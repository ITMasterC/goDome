package main

import (
	"demo-project/crawler/engine"
	"demo-project/crawler/scheduler"
	"demo-project/crawler/zhenai/parser"
)

func main()  {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	
	//e.Run(engine.Request{
	//	Url:        "http://www.7799520.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	
	e.Run(engine.Request{
		Url: "http://www.7799520.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}

