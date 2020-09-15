package engine

import (
	"demo-project/crawler/fetcher"
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}
	
	for len(requests) > 0  {
		r:=requests[0]
		requests = requests[1:]
		
		log.Printf("Fetching %s", r.Url)
		
		parseResult, err := worker(r)
		if err != nil{
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}
		requests = append(requests, parseResult.Requests...)
		//for _, item := range parseResult.Items{
		//	//log.Printf("Got item %v ", item)//%v - 打印时不进行转义，原来是什么就打印什么
		//}
	}
}

func worker(r Request) (ParserResult, error) {
	body, err := fetcher.Fetch(r.Url)
	//fmt.Printf("%s\n", body)
	if err != nil{
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	
	return r.ParserFunc(body), nil
}