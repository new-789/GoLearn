package main

import (
	"fmt"
	"github.com/luffycity/iniconfig"
)

type Config struct {
	// 声名一个结构体对应我们配置文件中 [] 号中的内容，并通过反射的 tag 属性使 ini 文件中 [] 节点名与结构体中的字段名进行匹配
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	// 声名 ServerConfig 结构体对应 [server] 节点下的所有名称，并设置 tag 映射关系
	Ip   string `ini:"ip"`
	Port uint   `ini:"port"`
}

type MysqlConfig struct {
	// 声名 MysqlConfig 结构体对应 [mysql] 节点下的所有名称，并设置 tag 映射关系
	Username string  `ini:"username"`
	Passwd   string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

func unmarshal(filename string, conf interface{}) {
	err := iniconfig.UnMarshalFile(filename, conf)
	if err != nil {
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
