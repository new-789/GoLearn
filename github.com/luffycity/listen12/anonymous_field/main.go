package main

import (
	"fmt"
)

// 声名一个结构体类型
type Users struct {
	Username string // 定义一个类型为 string Username 字段
	Sex      string
	int      // 定义一个类型为 int 类型的匿名字段
	string
}

func main() {
	var user Users         // 通过声名的结构体 Users 声名一个 user 变量
	user.Username = "test" // 通过 user 变量调用结构体中的 Username 字段名
	user.Sex = "男"
	user.int = 88 // 通过 user 结构体声名的变量访问匿名字段
	user.string = "湖北"
	fmt.Printf("user：%#v", user)
}
