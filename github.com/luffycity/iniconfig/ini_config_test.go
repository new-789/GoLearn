package iniconfig

import "testing"

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

//func TestIniConfig(t *testing.T) {
//	//fmt.Println("hello world")
//	// 读取 ini 配置文件, 返回一个 byte 数组类型的数据，和一个错误
//	data, err := ioutil.ReadFile("./config.ini")
//	if err != nil{
//		t.Error("读取文件错误")
//	}
//
//	var conf Config
//	er := UnMarshal(data, &conf)  // 调用定义的反序列化方法
//	if er != nil {
//		t.Errorf("unmarshal failed, err:%v", er)
//	}
//	// 打印日志
//	t.Logf("unmarshal success, conf:%#v", conf)
//
//	// 调用 Marshal 方法将结构体类型的信息通过反射序列化为配置文件内容的结构
//	result, er := Marshal(conf)
//	if er != nil {
//		t.Errorf("Marshal failed, errorL%v", err)
//	}
//	t.Logf("Marshal success, conf:%s", string(result))
//
//	// 调用 MarshalField 方法直接将内容存入文件
//	MarshalFile(conf, "./test.ini")
//}

func TestIniConfigFile(t *testing.T) {
	filename := "./test.ini"
	var conf Config = Config{
		ServerConf: ServerConfig{"192.168.8.100", 8080},
		MysqlConf: MysqlConfig{"test",
			"root",
			"test",
			"localhost",
			3306,
			1.2},
	}
	er := MarshalFile(filename, conf)
	if er != nil {
		t.Errorf("marshal fialed, error:%s", er)
		return
	}

	var conf2 Config
	err := UnMarshalFile(filename, &conf2)
	if err != nil {
		t.Errorf("unmarshal fialed, error:%s", err)
	}

	t.Logf("unmarshal success, conf:%#v", conf2)
}
