package main

import (
	"fmt"
	"github.com/luffycity/listen17/log"
)

func LogLevelAssert(l interface{}) {
	switch v := l.(type) {
	// 使用接口设定规则标准后，无论断言出来的是何种类型都只需要调用相同的方法
	case log.FileLog:
		v.LogDebug("这是 file 程序的 Debug 日志测试信息")
		v.LogWarn("这是 file 程序的 Warn 日志测试信息")
	case log.Console:
		v.LogDebug("这是 console 程序的 Debug 日志测试信息")
		v.LogWarn("这是 console 程序的 Warn 日志测试信息")
	default:
		fmt.Println("对不起，这是什么类型的内容，我们暂时无法判断")
	}
}

func main() {
	// 将日志内容写入到文件
	// 通过 log 包中的 NewFileLog 方法生成一个文件实例 file
	//file := log.NewFileLog("c:/a.log")
	//file.LogDebug("这是 Debug 日志信息")
	//file.LogWarn("这是 Warn 日志信息")

	// 将日志打印到控制台
	//console := log.NewConsole("ssss")
	//console.LogDebug("这是打印到控制台的 debug 日志测试")
	//console.LogWarn("这是打印到控制台的 warn 日志测试")

	file := log.NewFileLog("c:/test.log")
	LogLevelAssert(file)

	console := log.NewConsole("sss")
	LogLevelAssert(console)
}
