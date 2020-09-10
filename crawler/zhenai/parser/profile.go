package parser

import (
	"demo-project/crawler/engine"
	"demo-project/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<span class="age s1">([0-9]+)岁</span>`) //[\d]表示所有的数字
var heightRe = regexp.MustCompile(`<span class="height">([\d]+)cm</span>`)
var incomeRe = regexp.MustCompile(`<li>收入：<span>[\d]+以下</span></li>`)
var educationRe = regexp.MustCompile(`<span class="education">([^<]+)</span>`)
var occupationRe = regexp.MustCompile(`<li>职业：<span>([^<]+)</span></li>`)
var hoKouRe = regexp.MustCompile(`<li>籍贯：<span>([^<]+)</span></li>`)
var xinZouRe = regexp.MustCompile(`<li>星座：<span>([^<]+)</span></li>`)
var bloodRe = regexp.MustCompile(`<li>血型：<span>([^<]+)</span></li>`)
var xianJu = regexp.MustCompile(`<li>现居：<span>([^<]+)</span></li>`)
var genderRe = regexp.MustCompile(`<span>\n([^\s]+)\s</span>`)
var marriageRe = regexp.MustCompile(`<span class="marrystatus">([^<]+)</span>`)

func ParseProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil{
		
		profile.Age = age
	}
	if height, e := strconv.Atoi(extractString(contents, heightRe)); e == nil{
		profile.Height = height
	}
	
	profile.Marriage = extractString(contents, marriageRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Education = extractString(contents, educationRe)
	profile.HoKou = extractString(contents, hoKouRe)
	profile.XinZou = extractString(contents, xinZouRe)
	profile.XianJu = extractString(contents, xianJu)
	profile.Blood = extractString(contents, bloodRe)
	profile.Gender = extractString(contents, genderRe)
	
	return engine.ParserResult{
		Items:[]interface{}{profile},
	}
}

func extractString(contents []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		return string(match[1])
	}else{
		return ""
	}
}
