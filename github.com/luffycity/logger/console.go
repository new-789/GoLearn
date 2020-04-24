package logger

import (
	"fmt"
	"os"
)

type ConsoleLog struct {
	level int
}

func NewConsoleLogger(config map[string]string) (logger LogInterface, err error) {
	logLeve, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found logLevel")
		return
	}
	level := GetLogLevel(logLeve)
	logger = &ConsoleLog{
		level: level,
	}
	return
}

func (c *ConsoleLog) Init() {

}

func (c *ConsoleLog) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	c.level = level
}

func (c *ConsoleLog) Debug(format string, args ...interface{}) {
	if c.level > LogLevelDebug {
		return
	}
	logData := WriteLog("debug", format, args...)
	fmt.Fprintf(os.Stdout, "%s %s 【%s：%s：%d】 %s\n",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNum, logData.Message)
}

func (c *ConsoleLog) Trace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	logData := WriteLog("trace", format, args...)
	fmt.Fprintf(os.Stdout, "%s %s 【%s：%s：%d】 %s\n",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNum, logData.Message)
}

func (c *ConsoleLog) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	logData := WriteLog("info", format, args...)
	fmt.Fprintf(os.Stdout, "%s %s 【%s：%s：%d】 %s\n",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNum, logData.Message)
}

func (c *ConsoleLog) Warn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	logData := WriteLog("warn", format, args...)
	fmt.Fprintf(os.Stdout, "%s %s 【%s：%s：%d】 %s\n",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNum, logData.Message)
}

func (c *ConsoleLog) Error(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	logData := WriteLog("error", format, args...)
	fmt.Fprintf(os.Stdout, "%s %s 【%s：%s：%d】 %s\n",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNum, logData.Message)
}

func (c *ConsoleLog) Fatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	logData := WriteLog("fatal", format, args...)
	fmt.Fprintf(os.Stdout, "%s %s 【%s：%s：%d】 %s\n",
		logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNum, logData.Message)
}

func (c *ConsoleLog) Close() {

}
