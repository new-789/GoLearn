package main

import (
	"fmt"
)

func main() {
	// 定义变量语法一
	// var a int
	// var b string
	// var c bool
	// var d float32 = 10.1223457
	// 定义变量语法二，推荐使用
	var (
		a int
		b string = "word"
		c bool
		d float32 = 12.2343546
	)
	// 格式化输出：使用占位符进行格式化输出:
	// %d ：表示数值类型数据占位符，// %s： 表示字符串类型数据占位符
	// %t ：表示布尔类型数据占位符， // %f：表示浮点类型数据占位符
	fmt.Printf("a=%d b=%s c=%t d=%f\n", a, b, c, d)

	// 给变量赋值语法
	a = 100
	b = "hello" // 注，后面再次给变量赋值时会更改当前变量已存储在的值
	c = true
	d = 3.1415926
	fmt.Printf("a=%d b=%s c=%t d=%f", a, b, c, d)
}
