package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// 声名一个到文件的日志模块,用来在其中定义日志打印的特定内容字段.
// 我们需要的日志格式为：2006-01-02 15:04:05.999 LEVEL fileName lineNum logInfo
type FileLogger struct {
	Level   int    // 日志的级别
	LogPath string // 日志输出到文件的路径
	LogName string // 日志文件的名称
	// 定义文件句柄，用来操作两个文件目的是为了将日志内容进行区分
	file     *os.File
	warnFile *os.File
	// 使用 chan 定义一个队列, chan 后面的为队列中的数据类型，此处为 util.go 文件中定义的结构体 LogData
	LogDataChan chan *LogData
	// 该字段用来确认日志切分方式
	LogSplitType  int
	logSplitSize  int64
	lastSplitHour int
}

// 声名一个构造函数，返回一个日志接口和初始化失败后的错误信息
func NewFileLog(config map[string]string) (logger LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found logPath")
		return
	}
	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found logName")
		return
	}
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found logLevel")
	}
	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}

	// 获取用户设置切分日志的规则，如果没有设置则默认以时间进行切分
	var logSplitType int = LogSplitTypeHour
	var logSplitSize int64
	logSplitStr, ok := config["log_split_type"]
	if !ok { // 如果不存在切分参数则默认按时间段进行切分
		logSplitStr = "hour"
	} else {
		if logSplitStr == "size" { // 传入的参数为大小则按大小进行切分
			logSplitSizeStr, ok := config["log_split_size"]
			if !ok { // 如果没有传入切分的大小则设置默认大小为 100M
				logSplitSizeStr = "104857600"
			}
			/* 将 logSplitSize 对应的字符串类型的数字转换为整数
			 ParseInt 接收三个参数:
				第一个参数为需要转换的字符串
			    第二个参数为b表示转换为多少进制
				第三个参数为表示转换为多少位的整数类型(如32位、64位)
			该方法返回两个值，第一个值为转换成功之后的内容，第二个值为错误信息
			*/
			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64)
			if err != nil {
				logSplitSize = 104857600
			}
			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeHour
		}
	}

	// 将字符串类型的数据转换为整数，返回一个转换之后的数据以及一个错误
	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}
	level := GetLogLevel(logLevel)
	logger = &FileLogger{
		Level:   level,
		LogPath: logPath,
		LogName: logName,
		// 初始化队列使用 make 方法，第一个参数为 type， 第二个参数为队列的 size
		LogDataChan: make(chan *LogData, chanSize),
		// 日志文件切分字段
		LogSplitType:  logSplitType,
		logSplitSize:  logSplitSize,
		lastSplitHour: time.Now().Hour(),
	}
	logger.Init() // 打开文件操作
	return
}

// init 方法在程序运行前率先加载到内存中进行初始化
func (f *FileLogger) Init() {
	// 普通日志的文件
	fileName := fmt.Sprintf("%s/%s.log", f.LogPath, f.LogName)
	// 错误日志的文件
	errorFile := fmt.Sprintf("%s/%s.log.wf", f.LogPath, f.LogName)

	// 打开文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	// 打开文件如果出错则直接让程序停止运行，并输出错误
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, error: %v", fileName, err))
	}
	f.file = file

	// 错误日志和 Fatal 日志的文件
	errFile, e := os.OpenFile(errorFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if e != nil {
		panic(fmt.Sprintf("open errfile %v failed, error: %v", errFile, e))
	}
	f.warnFile = errFile

	// 开启线程方法，writeLogBackground 为线程的入口函数
	go f.writeLogBackground()
}

// 按照时间切分日志文件方法
func (f *FileLogger) splitFileHour(WarnFile bool) {
	// 获取当前时间对象,以达到复用时间对象的目的
	nowTimeObj := time.Now()
	hour := nowTimeObj.Hour()
	// 判断最后一次写入文件的时间是否与当前写入日志时间是否相等，相等则不作切分
	if hour == f.lastSplitHour {
		return
	}
	// 上一次更新日志时间不相等则更新为当前时间，因为本次也需要写内容到文件所以将时间更新为本次写入的时间
	f.lastSplitHour = hour

	// 切分日志前先对文件进行备份
	var backupFileName string
	var fileName string // 声名变量用来记录新写入的文件名
	if WarnFile {       // 通过 bool 类型的值判断文件类型
		backupFileName = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02d%02d",
			f.LogPath, f.LogName, nowTimeObj.Year(), nowTimeObj.Month(), nowTimeObj.Day(), f.lastSplitHour)
		fileName = fmt.Sprintf("%s/%s.log.wf", f.LogPath, f.LogName)
	} else {
		backupFileName = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d",
			f.LogPath, f.LogName, nowTimeObj.Year(), nowTimeObj.Month(), nowTimeObj.Day(), f.lastSplitHour)
		fileName = fmt.Sprintf("%s/%s.log", f.LogPath, f.LogName)
	}
	// 备份成功后写入日志到一个新的文件
	file := f.file
	if WarnFile {
		file = f.warnFile
	}
	file.Close()
	os.Rename(fileName, backupFileName) // 备份操作即对当前文件重命名
	// 重新打开备份前的文件名，如果不存在则会创建，达到新建文件的目的
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	// 判断当前被关闭的文件，然后对其重新赋值
	if WarnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
}

// 按照文件大小切分日志文件方法
func (f *FileLogger) splitFileSize(WarnFile bool) {
	// 获取当前日志文件的大小,并根据不同的日志类型获取不同的文件
	file := f.file
	if WarnFile {
		file = f.warnFile
	}
	// 获取日志文件的大小,Stat 方法可以获取当前文件的基本信息，如时间大小等，file 对象为 *os.File 类型的时间对象
	statInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := statInfo.Size() // 获取文件大小方法
	// 如果当前文件小于设置的文件大小限制则直接返回不作任何操作
	if fileSize <= f.logSplitSize {
		return
	}

	// 切分日志前先对文件进行备份
	nowTimeObj := time.Now()
	var backupFileName string
	var fileName string // 声名变量用来记录新写入的文件名
	if WarnFile {       // 通过 bool 类型的值判断文件类型
		backupFileName = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02d%02d%02d%02d",
			f.LogPath, f.LogName, nowTimeObj.Year(), nowTimeObj.Month(), nowTimeObj.Day(),
			nowTimeObj.Hour(), nowTimeObj.Minute(), nowTimeObj.Second())
		fileName = fmt.Sprintf("%s/%s.log.wf", f.LogPath, f.LogName)
	} else {
		backupFileName = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d%02d%02d",
			f.LogPath, f.LogName, nowTimeObj.Year(), nowTimeObj.Month(), nowTimeObj.Day(),
			nowTimeObj.Hour(), nowTimeObj.Minute(), nowTimeObj.Second())
		fileName = fmt.Sprintf("%s/%s.log", f.LogPath, f.LogName)
	}

	file.Close()
	os.Rename(fileName, backupFileName) // 备份操作,即对当前文件重命名
	// 重新打开备份前的文件名，如果不存在则会创建，达到新建文件的目的
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	// 判断当前被关闭的文件，然后对其重新赋值
	if WarnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
}

// 日志切分方法实现
func (f *FileLogger) checkSplitFile(WarnFile bool) {
	// 如果切分类型为时间,则按照时间进行切分日式文件
	if f.LogSplitType == LogSplitTypeHour {
		f.splitFileHour(WarnFile)
	}

	// 如果切分类型为大小,则按照文件大小进行切分日式文件
	f.splitFileSize(WarnFile)
}

// 实现一个 go 线程入口函数方法
func (f *FileLogger) writeLogBackground() {
	for logData := range f.LogDataChan {
		var file *os.File = f.file
		if logData.WarnAndFatal {
			file = f.warnFile
		}
		// 存入文件之前调用日志切分功能进行日志的切分
		f.checkSplitFile(logData.WarnAndFatal)
		fmt.Fprintf(file, "%s %s 【%s：%s：%d】 %s\n",
			logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNum, logData.Message)
	}
}

func (f *FileLogger) SetLevel(level int) {
	// 用来判断用户传入的日志级别是否在我们定义的日志级别范围内,不存在则设置一个日志默认级别为 Debug
	if level < LogLevelDebug || level > LogLevelFatal {
		f.Level = LogLevelDebug
	}
	f.Level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.Level > LogLevelDebug {
		return
	}
	logData := WriteLog("debug", format, args...)
	// select 关键字用来判断一个队列是否存满了数据
	select {
	// 将数据放入队列使用 <- 符号，左边为初始化的队列对象，右边为存入队列的数据
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.Level > LogLevelTrace {
		return
	}
	logData := WriteLog("trace", format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.Level > LogLevelInfo {
		return
	}
	logData := WriteLog("info", format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.Level > LogLevelWarn {
		return
	}
	logData := WriteLog("warn", format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.Level > LogLevelError {
		return
	}
	logData := WriteLog("error", format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.Level > LogLevelFatal {
		return
	}
	logData := WriteLog("fatal", format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
