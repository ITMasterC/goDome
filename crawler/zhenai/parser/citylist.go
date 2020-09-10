package parser

import (
	"demo-project/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.7799520.com/zhenghun/[\s0-9a-z]+)"[^>]*>([^>]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	//limit := 9
	for _, m := range matches{
		result.Items = append(result.Items, "City " +string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url:string(m[1]), ParserFunc: ParseCity})
		
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1] )
		//for _, subMatch := range m{
		//	fmt.Printf("%s ", subMatch)
		//}
		//fmt.Println()
		//limit--
		//if limit == 0{
		//	break
		//}
	}
	//fmt.Printf("%d\n", len(matches))
	return result
}