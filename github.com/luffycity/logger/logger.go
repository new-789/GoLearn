package logger

import (
	"fmt"
)

var log LogInterface

/*
初始化日志对象方法，在业务逻辑中直接调用该对象即可
参数说明：
name： 表示传入的内容，并根据内容的不同调用不同的日志原型，暂定有下面两种日志类型
	file: “初始化一个文件日志实例”
	console: “初始化一个控制台的日志实例”
config : 由于控制台与文件日志原型所接收的参数不同，所以此处用一个 map 类型的参数用来保存需要使用到的所有参数数据
*/
func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		log, err = NewFileLog(config)
	case "console":
		log, err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("unsupport logger name:%s ", name)
	}
	return
}

func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}

func Trace(format string, args ...interface{}) {
	log.Trace(format, args...)
}

func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}

func Close() {
	log.Close()
}
