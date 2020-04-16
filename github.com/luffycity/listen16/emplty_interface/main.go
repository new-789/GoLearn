package main

import (
	"fmt"
)

// 声名一个函数，并接受一个空接口类型的变量
func describe(a interface{}) {
	// 打印输出接收到的值所对应的类型和值
	fmt.Printf("Type = %T, value = %v\n", a, a)
}

// 声名一个结构体
type Student struct {
	Name string
	Sex  string
	Age  int
}

func main() {
	// 声名一个 int 类型的变量 b
	var b int = 100
	// 调用 describe 方法并将 b当做参数
	describe(b)

	// 声名一个 string 类型的变量 c
	var c string = "hello world"
	describe(c)

	// 通过结构体初始化一个 stu 对象
	var stu Student = Student{
		Name: "user01",
		Sex:  "man",
		Age:  99,
	}
	describe(stu)
}
