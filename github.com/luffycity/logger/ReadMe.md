##  日志模块使用几板斧
1. 将日志模块复制到 ``$GOPATH$`` 变量的 src 目录下，并在功能模块中导入开发的日志包 logger
2. 定义一个函数，接收日志的输出方式，文件保存路径，文件名，日志级别, 切分方式参数
        输出方式有：'file'、'console' 文件及控制台可选
        日志级别有：'debug'、'trace'、'info'、'warn'、'error'、'fatal' 几个级别可选，debug 建议在开发阶段使用
        切分类型有：'size'、'hour' 可选
3. 使用 make 方法初始化一个 map[string]string 类型的 Map 对象
4. 调用日志模块的 InitLogger 方法，并将日志显示方式及初始化 map 获得的对象当做参数传给该函数

---
**使用示例**
```go
package main

import (
	"github.com/luffycity/logger"  // 1. 导入开发的日志包
)

// 2. 定义一个函数
func initLogger(name, logPath, logName string, level string, logSplitType string) (err error) {
  // 3. 使用 make 方法初始化一个 map[string]string 类型的 Map 对象
	m := make(map[string]string, 8)
	m["log_path"] = logPath  // 设置日志文件的路径
	m["log_name"] = logName  // 设置日志文件的名称
	m["log_level"] = level  // 设施日志级别
	m["log_split_type"] = logSplitType  // 设置文件的切分方式
	// 4. 调用日志模块的 InitLogger 方法，并将日志显示方式及初始化 map 获得的对象当做参数传给该函数
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}
	logger.Debug("init logger success")
	return
}

// 打印日志方法，伪代码，请根据实际情况编写相关逻辑
func Run(){
	for {
		logger.Debug("user server is running E:\\CodingFiles\\GolangCode\\src\\github.com\\luffycity\\listen17\\user_server")
		//time.Sleep(time.Second)
	}
}

func main() {
	initLogger("file", "E:/logs", "user_server", "debug", "size")
	Run()
	return
}
```
