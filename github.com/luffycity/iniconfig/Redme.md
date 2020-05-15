### 配置文件库使用说明
**1. 摘要**  
- 本配置文件库实现了两个功能  
   - 将 ini 配置文件通过映射的方式序列化为 struct 类型的数据  
   - 将 struct 类型的数据通过映射的方式反序列化为 ini 配置文件信息并存入文件  
   
**2. 使用说明**  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2.1. 按照常规在使用之前通过下面语法先导入 ``iniConfig`` 包
```go
import ("github.com/src/.../iniConfig")  // 路径根据自己的项目目录进行填写
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2.2. 根据 ini 配置文件中的节点信息建立相对应的结构体数据  
 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;本次使用的 ini 配置文件信息结构如下:
```
[server]
ip=192.168.8.100
port=8080

[mysql]
username=test
passwd=root
database=test
host=localhost
port=3306
timeout=1.2
```
 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;根据 ini 配置文件信息结构建立的 Go 语言对应的结构体数据如下:
 ```go
type Config struct {
	// 声名一个结构体对应我们配置文件中 [] 号中的内容，并通过反射的 tag 属性使 ini 文件中 [] 节点名与结构体中的字段名进行匹配
	ServerConf ServerConfig `ini:"server"`
	MysqlConf MysqlConfig `ini:"mysql"`
}

type ServerConfig struct {
	// 声名 ServerConfig 结构体对应 [server] 节点下的所有名称，并设置 tag 映射关系
	Ip string  `ini:"ip"`
	Port uint   `ini:"port"`
}

type MysqlConfig struct {
	// 声名 MysqlConfig 结构体对应 [mysql] 节点下的所有名称，并设置 tag 映射关系
	Username string  `ini:"username"`
	Passwd string    `ini:"passwd"`
	Database string  `ini:"database"`
	Host string      `ini:"host"`
	Port int         `ini:"port"`
	Timeout float32  `ini:"timeout"`
}
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2.3. 声名一个 filename 变量用来指定 ini 文件的路径，及使用 ``Config`` 结构体初始化一个变量
```go
filename := "./xxx.ini"  // 路径请根据自己的文件路径进行填写
var conf Config
``` 
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2.4. 调用 ``MarshalFile`` 接口方法，该方法用来将用户传入的结构体数据序列化后存入 ini 配置文件,该方法接收两个参数,并返回一个错误信息
- 参数说明：
   - filename: 该参数接收一个 ini 文件名/或文件路径
   - data interface{}: 使用结构体初始化声名的变量名称  

示例
```go
err := MarshalFile(filename, conf)
if err != nil{
    // 错误提示内容请自行定义
    return
}
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2.5. 调用 ``UnMarshalFile`` 接口方法，该方法用来将用户传入的 ini 配置文件内容序列化为 struct 类型的数据，该方法接收两个参数,并返回一个错误信息
- 参数说明：
   - filename: 该参数接收一个 ini 文件名/或文件路径
   - result interface{}: 使用结构体初始化声名的变量名称<font color=red> **注：该参数仅支持接收一个指针类型的数据，否则会抛出异常** </font>  

示例
```go
err := UnMarshalFile(filename, &conf)
if err != nil {
    // 错误提示内容请自行定义
    return
}
```

**3. 完整使用代码示例**
```go
package main

import (
	"fmt"
	"github.com/luffycity/iniconfig"
)

type Config struct {
	// 声名一个结构体对应我们配置文件中 [] 号中的内容，并通过反射的 tag 属性使 ini 文件中 [] 节点名与结构体中的字段名进行匹配
	ServerConf ServerConfig `ini:"server"`
	MysqlConf MysqlConfig `ini:"mysql"`
}

type ServerConfig struct {
	// 声名 ServerConfig 结构体对应 [server] 节点下的所有名称，并设置 tag 映射关系
	Ip string  `ini:"ip"`
	Port uint   `ini:"port"`
}

type MysqlConfig struct {
	// 声名 MysqlConfig 结构体对应 [mysql] 节点下的所有名称，并设置 tag 映射关系
	Username string  `ini:"username"`
	Passwd string    `ini:"passwd"`
	Database string  `ini:"database"`
	Host string      `ini:"host"`
	Port int         `ini:"port"`
	Timeout float32  `ini:"timeout"`
}

func unmarshal(filename string, conf interface{}) {
	err := iniconfig.UnMarshalFile(filename, conf)
	if err != nil{
		fmt.Printf("unmarshal failed. error:%s", err)
		return
	}
	fmt.Printf("unmarshal succ, conf:%#v\n", conf)
}

func marshal(filename string, result interface{}) {
	err := iniconfig.MarshalFile(filename, result)
	if err != nil {
		fmt.Printf("marshal failed. error:%s", err)
		return
	}
	fmt.Println("============= save file success ===============")
}

func main() {
	filename := "D:/test.ini"
	// 将 ini 格式类型的配置文件反序列化为结构体
	var conf Config
	unmarshal(filename, &conf)

	//
	filename2 := "test.ini"
	marshal(filename2, conf)
}
```