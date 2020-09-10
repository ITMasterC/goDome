package parser

import (
	"demo-project/crawler/engine"
	"regexp"
)

//const cityRe  = `<a href="http://www.7799520.com/user/*.html" target="_blank">*([^<]+)\s </a>`
const cityRe  = `<a href="(http://www.7799520.com/user/[0-9]+.html)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParserResult  {
	re := regexp.MustCompile(cityRe)
	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", matches)
	result := engine.ParserResult{}
	for _, m := range matches{
		name := string(m[2])
		result.Items = append(result.Items, "User " +name)
		result.Requests = append(result.Requests, engine.Request{Url:string(m[1]),
			ParserFunc: func(c []byte ) engine.ParserResult {
				return ParseProfile(c, name)
			}})
		
	}
	return result
}
