package main

import (
	"github.com/luffycity/logger"
)

func initLogger(name, logPath, logName string, level string) (err error) {
	m := make(map[string]string, 8)
	m["log_path"] = logPath      // 设置日志文件的路径
	m["log_name"] = logName      // 设置日志文件的名称
	m["log_level"] = level       // 设施日志级别
	m["log_split_type"] = "size" // 设置文件的切分方式
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}
	logger.Debug("init logger success")
	return
}

func Run() {
	for {
		logger.Debug("user server is running E:\\CodingFiles\\GolangCode\\src\\github.com\\luffycity\\listen17\\user_server")
		//time.Sleep(time.Second)
	}
}

func main() {
	initLogger("file", "E:/logs", "user_server", "debug")
	Run()
	return
}
