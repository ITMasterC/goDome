package main

import (
	"fmt"
	"regexp"
)

const text  = `
My email is ITMaster_C@163.com
email1 is abc@qq.com
email2 is     kkk@ttt.com
email3 is ddd@abc.com.cn
`

func main()  {
	//re := regexp.MustCompile("ITMaster_C@163.com")
	//. 表示前面所有的字符
	//[a-zA-Z0-9] 表示所有字母和数字
	//re := regexp.MustCompile(`.+@.+\..+`)
	re := regexp.MustCompile(`([a-zA-Z0-9_.]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)//添加括号可以提取相关的内容
	//match := re.FindString(text)//匹配第一个符合的字符串
	//match := re.FindAllString(text, -1)//-1 表示查找所有符合项
	math := re.FindAllStringSubmatch(text, -1)//匹配内容 并进行拆分
	for _, m := range math{
		fmt.Println(m)
	}
}