package main

import (
	"fmt"
)

func testPoint1() {
	var a int32 = 100
	var b *int32  // 定义一个指针类型的变量语法
	if b == nil { // 判断一个指针变量是否为空语法(有没有感觉和判断数组为空语法一样)
		fmt.Printf("b is nil addr\n")
	}
	b = &a // 获取变量 a 的内存地址赋值给变量 b
	fmt.Printf("b is addr:%p, b is value: %v\n", &b, b)
}

func testPoint2() {
	var a int = 100 // 定义一个 int 类型的变量 a，并初始化值为 100
	var b *int = &a // 定义一个 Int 类型的指针变量 b，并初始化指向变量 a 的内存地址
	*b = 200        // 通过指针变量语法 *b 更改指向内存地址中存储的值
	fmt.Printf("b指针变量指向内存地址中存储的值：%d\n", *b)
	fmt.Printf("a is value:%d\n", a) // 再次打印变量 a 的内容，同样被更改为 200
}

func main() {
	//testPoint1()
	testPoint2()
}
