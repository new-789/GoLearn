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

func modify(point *int) {
	// 定义 testPoint3 函数，接收一个指针类型的传参
	*point = 88 // 更改指针类型参数 point 接收到的值为 88
}

func testPoint3() {
	var a int = 100                               // 声名一个 int 类型的变量 a，并初始化为 100
	var b *int = &a                               // 声名一个 int 类型的指针类型变量 b， 并初始化为变量 a 的内存地址
	modify(b)                                     // 调用函数并将指针类型变量 b 当做参数传递, 相当于将指针变量 b 的内存地址拷贝一份给函数中定义的参数，即 point := &b
	fmt.Printf("a value:%d, b value:%d\n", a, *b) // 函数中修改指针参数值后输出变量 a 与指针变量 b 相对应的值
}

func main() {
	//testPoint1()
	//testPoint2()
	testPoint3()
}
