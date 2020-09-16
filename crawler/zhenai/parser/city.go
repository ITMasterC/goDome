package parser

import (
	"demo-project/crawler/engine"
	"regexp"
)

//const cityRe  = `<a href="http://www.7799520.com/user/*.html" target="_blank">*([^<]+)\s </a>`
var profileRe  = regexp.MustCompile(`<a href="(http://www.7799520.com/user/[0-9]+.html)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.7799520.com/zhenghun/[^"]+)">`)

func ParseCity(contents []byte) engine.ParserResult  {
	//matches := re.FindAll(contents, -1)
	matches := profileRe.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", matches)
	result := engine.ParserResult{}
	for _, m := range matches{
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{Url:string(m[1]),
			ParserFunc: func(c []byte ) engine.ParserResult {
				return ParseProfile(c, name)
			}})
		
	}
	
	submatch := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range submatch{
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	
	return result
}
