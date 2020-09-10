package main

import "github.com/gin-gonic/gin"

// 不知道为什么你的系统环境设置一直改变不了
// set GOPROXY=GOPROXY=https://goproxy.io 这一行设置不得 是set GOPROXY=https://goproxy.io才
// // 终于可以了  系统没有重新加载环境 你得重启终端才可以

// 升级一下golang版本 这个太低了 环境设置的一些功能遍历没有

// 环境配置 设置go path： win: set GOPATH="your go path dir" linux: export GOPATH="go path dir"  设置 goproxy 做代理，国内网络不行
// set GOPROXY=https://goproxy.io 设置go module set GO111MODULE=on
// golang 一般很少用vendor环境的，环境是你开发的时候用的而已，有一个go mod就行了， 正式环境一般直接执行 go build出来的二进制文件
func main()  {
	r :=gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"hello":"wolrd",
		})
	})
	r.Run(":8080")
}