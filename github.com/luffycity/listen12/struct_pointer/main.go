package main

import "fmt"

// 定义 struct 结构体类型
type Users struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user *Users
	fmt.Printf("user=%v\n", user)
	//if user == nil{
	//	fmt.Println("这是一个空的结构体指针")
	//}

	// 空指针初始化方式一
	var user01 *Users = &Users{}
	user01.Username = "Test1" // 等同于 (*user01).Username = "Test"
	user01.Age = 19
	fmt.Printf("user01=%#v\n", user01)

	// 空指针初始化方式二
	var user02 *Users = &Users{
		Username:  "test2",
		Age:       90,
		AvatarUrl: "dssdzsdfsd",
	}
	fmt.Printf("user02=%#v\n", user02)

	// 空指针初始化方式三
	var user03 *Users = new(Users)
	user03.Username = "test3"
	user03.Age = 23
	user03.AvatarUrl = "shhhkllasjdjsaj"
	fmt.Printf("user03=%#v\n", user03)
}
