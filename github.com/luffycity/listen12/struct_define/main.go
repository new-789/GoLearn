package main

import "fmt"

// 定义一个结构体类型的变量 Users
type Users struct {
	// 结构体字段
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	// 通过定义的 Users 结构体声名一个变量 user
	var user Users
	// 通过变量初始化结构体中的字段
	user.Username = "ZhouFang"
	user.Sex = "女"
	user.Age = 32
	user.AvatarUrl = "https://pmlchina.cn/image/aaa.jpg"

	// 访问结构体重的字段同样使用我们通过结构体变量定义的变量名加结构体中的字段名
	//fmt.Printf("Username:%s\n Sex:%s\n Age:%d\n AvatarUrl:%s\n",
	//	user.Username, user.Sex, user.Age,user.AvatarUrl)

	//var user2 Users = Users{
	//	Username:"Test",
	//	Sex:"女",
	//	Age:18,
	//	AvatarUrl:"https://test.test.test/",
	//}
	//fmt.Printf("user2:%#v\n", user2)
	//fmt.Println()
	//user3 := Users{
	//	Username:  "HaHa",
	//	Sex:       "变形",
	//	Age:       0,
	//	AvatarUrl: "http://hshsh.hshsh.hshsh/",
	//}
	//fmt.Printf("user3:%#v\n",user3)

	var user4 Users
	fmt.Printf("user4:%#v\n", user4)
}
