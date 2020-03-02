package main

import (
	"flag" // 导入命令行参数解析包
	"fmt"
	"math/rand" // 导入随机算法相关包
	"time"
)

var (
	// 这里定义的两个变量用来获取用户输入的命令行参数
	length  int    // 用来获取用户输入的命令行参数 -l 所指定的内容，用来确定指定的密码长度
	charset string // 用来获取用户输入的命令行参数 -t 所指定的内容，用来确定指定的密码类型
)

const ( // 定义常量，设定密码组成的字符内容
	NumStr  = "0123456789"                                         // 密码为数字的内容来源
	CharStr = "ABCDEFGHIJKMNOPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz" // 密码为字符串内容来源
	SpecStr = "!@#$%^{=}[+]<>&(*?/~)"                              // 密码为特殊字符的内容来源
)

func parseArgs() {
	// flag.IngVar 指定参数绑定的变量，类型为整数类型
	flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
	// flag.StringVar 指定参数绑定的变量，类型为字符串类型
	flag.StringVar(&charset, "t", "num",
		`-t 指定密码生成的字符集,
    	num:只使用数字[0-9],
	char:只使用英文字母[a-z/A-Z]，mix:使用英文字母和数字,
	advance: 使用英文字母加数字和特殊符号`)
	flag.Parse() // 解析命令行定义的 flag
}

func generatePassword() string {
	var password []byte = make([]byte, 0, length) // 定义一个切片，用来保存生成的密码
	var sourceStr string                          // 用来保存用户对密码的内容的字符串拼接，以便获取随机内容当做密码
	// 当用户使用 -t 参数指定为 num 时更新 sourceStr 的内容为数字
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" { // 当用户使用 -t 参数指定为 char 时更新 sourceStr 的内容为字母
		sourceStr = CharStr
	} else if charset == "mix" { // 当用户使用 -t 参数指定为 mix 时更新 sourceStr 的内容为数字加字母
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr) // 格式化打印输出将其中的数字格式化微字符串
	} else if charset == "advance" { // 当用户使用 -t 参数指定为 advance 时更新 sourceStr 的内容为数字加字母和特殊字符
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr) // 格式化打印输出将其中的数字格式化微字符串
	} else { // 如果用户输入有误则直接让生成的密码为数字
		sourceStr = NumStr
	}
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))            // 使用 rand 模块随机生成字符串的下标,rand.Intn 表示随机生成 0 到 n 的数字
		password = append(password, sourceStr[index]) // 通过随机生成的字符串下标获取内容添加到 password 切片中
	}
	return string(password) // 强制将 password 切片转换为字符串类型的数据
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间的纳秒初始化随机化种子
	parseArgs()
	//fmt.Printf("length:%d\n string:%s\n", length, charset)
	password := generatePassword()
	fmt.Printf("password: %v\n", password)
}
