package logger

import (
	"fmt"
	//"os"
	"path"
	"runtime"
	"time"
)

// 声名一个结构体用来在 chan 队列中存储日志信息
type LogData struct {
	Message      string
	TimeStr      string
	LevelStr     string
	FileName     string
	FuncName     string
	LineNum      int
	WarnAndFatal bool // 用来区分是什么日志类型并分别写入不同文件
}

// 声名 GetLineInfo 函数用来获取出错的文件名及代码行号
func GetLineInfo() (filename string, funcName string, lineNum int) {
	// 返回的 pc 为指令计数器，即 CPU 执行到那条指令了, 参数 3 表示当前调用日志功能打印日志的地方，根据调用的层级进行调整
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		filename = file
		// 通过 runtime.FuncForPC(pc).Name() 方法获取当前指令所在的函数名
		funcName = runtime.FuncForPC(pc).Name()
		lineNum = line
	}
	return
}

/*
写入文件操作分为以下几个步骤以达到异步写日志的方法
	1. 当业务调用日志的方法时，我们将日志相关的数据写入到 chan(队列)
	2. 然后开启一个后台线程不断的动 chan 中获取这些日志信息，最终写入到文件
*/
func WriteLog(level string, format string, args ...interface{}) *LogData {
	// 获取当前时间并设置时间格式,注设置时间格式的值必须为 2006-1-2 15-04-05，格式可任意设置但是值必须为该时间
	timeNow := time.Now().Format("2006-01-02 15:04:05.999")

	// 调用 GetLevelText 方法获取日志级别信息
	//levelNum := GetLogLevel(level)

	// 调用 GetLineInfo 方法获取文件名、函数名、及行号
	fileName, funcName, lineNum := GetLineInfo()
	fileName = path.Base(fileName)      // 获取整个路径最后的文件名
	funcName = path.Base(funcName)      // 获取整个路径最后的函数名
	msg := fmt.Sprintf(format, args...) // 将日志格式和传入的参数格式化到一个变量中

	// 写入日志信息到文件,通过 LogDATA 生成一个实例
	//_, err := fmt.Fprintf(file, "%s %s 【%s：%s：%d】 %s\n",
	//	timeNow, level, fileName,  funcName, lineNum, msg)
	logData := &LogData{
		Message:      msg,
		TimeStr:      timeNow,
		LevelStr:     level,
		FileName:     fileName,
		FuncName:     funcName,
		LineNum:      lineNum,
		WarnAndFatal: false,
	}
	// 用来判断错误级别，更改 WarnAndFatal 的值用来实现写入文件的不同
	if level == "error" || level == "warn" || level == "fatal" {
		logData.WarnAndFatal = true
	}
	return logData
}
