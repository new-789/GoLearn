package log

// 通过接口对输出日志的方式定义一个规范，即对所有需要用到日志的程序都需要实现定义好的日志接口规范
type LogInterface interface {
	LogDebug(msg string)
	LogWarn(msg string)
}
