package main

import "fmt"

type Address struct {
	Province   string
	City       string
	CreateTime string // Address 结构体中的 Ceatertime 字段
}

type Email struct {
	Account    string
	CreateTime string // Email 结构体中的 CreateTime 字段
}

//type Users struct {
//	Username string
//	Age      int
//	Sex      string
//	*Address // 匿名结构体嵌套
//}

type Users01 struct {
	City     string // 定义 City 字段
	*Address        // 匿名结构体，继承自 Address 结构体，且该结构体中同样存在 CreateTime 字段
	*Email          // 匿名结构体，继承自 Email 结构体，且该结构体中同样存在 CreateTime 字段
}

func test02() {
	var user01 Users01
	// 访问 Users01 结构体中的 City 字段
	//user01.City = "武汉市"
	//fmt.Printf("user02:%#v\n", user01)

	user01.Address = new(Address) // 使用 new 方法初始化 Address 匿名结构体字段
	user01.Email = new(Email)     // 使用 new 方法初始化 Email 匿名结构体字段
	//user01.Address.City = "宜昌市"  // 解决冲突访问匿名结构体中的字段方法
	// user02.Address.City 表示访问 Address 结构体中的 city 方法，
	//fmt.Printf("user01:%#v , City of address:%s \n", user01, user01.Address.City)

	user01.Email.CreateTime = "2020-02-02"   // 指定 Email 匿名结构体名指定需要访问 Email 结构体中的 CreateTime 字段名
	user01.Address.CreateTime = "2020-03-08" // 指定 Address 匿名结构体名指定需要访问 Address 结构体中的 CreateTime 字段名
	fmt.Printf("email of createtime:%s , address of createtime:%s \n", user01.Email.CreateTime,
		user01.Address.CreateTime)
}

//func test1() {
//	var user Users
//	// 初始化匿名结构体嵌套方式一
//	//user.Address = &Address{
//	//	Province: "HuBei",
//	//	City:     "WuHan",
//	//}
//
//	// 初始化结构体嵌套方式二
//	user.Address = &Address{}
//	user.Province = "河南"
//	user.City = "信阳"
//
//	//fmt.Printf("user:%#v\n", user)
//	// adrdress 嵌套的是一个指针类型的结构体，所以获取相应值时也应该加上 * 号，否则获取到的为一个内存地址，如 *user.address
//	fmt.Printf("user in address:%#v\n", *user.Address)
//}

func main() {
	//test1()
	test02()
}
