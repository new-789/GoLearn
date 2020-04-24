package logger

// 声名一个 LogInterface 接口，用来定义需要实现的日志级别方法
type LogInterface interface {
	Init()
	// 定义设置日志级别方法
	SetLevel(level int)
	// 定义日志级别方法... interface{} 表示接收一个可变参数, 第一个参数表示可设置格式化输出样式
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
