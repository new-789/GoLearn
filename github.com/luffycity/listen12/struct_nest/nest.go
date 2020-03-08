package main

import "fmt"

type Address struct {
	Province string
	City     string
}

type Users struct {
	Username string
	Age      int
	Sex      string
	*Address // 匿名结构体嵌套
}

func test1() {
	var user Users
	// 初始化匿名结构体嵌套方式一
	user.Address = &Address{
		Province: "HuBei",
		City:     "WuHan",
	}

	// 初始化结构体嵌套方式二
	user.Province = "河南"
	user.City = "信阳"

	//fmt.Printf("user:%#v\n", user)
	// adrdress 嵌套的是一个指针类型的结构体，所以获取相应值时也应该加上 * 号，否则获取到的为一个内存地址，如 *user.address
	fmt.Printf("user in address:%#v\n", *user.Address)
}

func main() {
	test1()
}
