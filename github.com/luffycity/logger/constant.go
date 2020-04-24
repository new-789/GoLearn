package logger

// 使用 const 关键字定义日志级别，当配置文件使用,iota 表示从 0 开始自动加 1
const (
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

// 声名常量用来定义日志的切分模式
const (
	LogSplitTypeHour = iota // 按照小时切分
	LogSplitTypeSize        // 按照大小切分
)

// 声名一个方法将数字级别的日志等级更改为可读性较高的内容进行打印输出
func GetLogLevel(level string) int {
	switch level {
	case "debug":
		return LogLevelDebug
	case "trace":
		return LogLevelTrace
	case "info":
		return LogLevelInfo
	case "warn":
		return LogLevelWarn
	case "error":
		return LogLevelError
	case "fatal":
		return LogLevelFatal
	}
	return LogLevelDebug
}
