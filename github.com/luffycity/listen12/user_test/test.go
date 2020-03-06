package main

import (
	"fmt"
	"github.com/luffycity/listen12/user" // 导入自定义包
)

func main() {
	// 通过导入的 user 包中定义的 Users 结构体定义一个变量 u ,
	// 注意此处变量名最好不要与包名同名
	//var u user.Users
	// 通过定义的变量初始化其中的两个字段
	//u.Username = "Test"
	//u.Age = 88
	//fmt.Printf("user:%#v\n", u)

	// 通过调用包中的构造函数来创建一个用户对象
	userInfo := user.NewUser("testUser", 88, "男", "小学")
	fmt.Printf("user:%#v\n", userInfo)
}
