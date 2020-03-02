package main

import (
	"fmt"
)

// go 语言函数接收参数及多个返回值写法，
// 第一个括号表示接收的参数，即接收参数的类型；
// 第二个括号内里的内容表示返回值的类型，一个类型表示一个返回值
func add(a int, b int) (int, int) {
	// 返回多个类容，直接使用 ，逗号隔开即可
	return a + b, a - b
}

func main() {
	// go 语言接收多个返回值语法，sum 表示接收第一个返回值，sub 表示接收第二个返回值
	sum, sub := add(2, 6)
	fmt.Println(sum, sub)
}
