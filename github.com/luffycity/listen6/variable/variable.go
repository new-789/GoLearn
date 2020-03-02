package main

import (
	"fmt"
)

var num int = 100 // 定义在函数外的为全局变量

func testGlobalVariable() {
	var num int = 300
	fmt.Printf("num=%d\n", num)
}

func testLocalVariable() {
	var localVariable int = 1000 // 在函数内部定义的局部变量,仅能在该函数内部调用
	fmt.Printf("localVariable = %d\n", localVariable)
	if localVariable == 1000 {
		var local int = 200 // 在if语句块中定义的局部变量，仅在 if 语句块中中有效
		fmt.Printf("local = %d\n", local)
	}
	if d := 10; d > 0 { // 此种语法定义的变量 d 同样算作是 if 语句块中的局部变量
		fmt.Printf("d = %d\n", d)
	} else {
		fmt.Printf("d = %d\n", d)
	}
	for i := 0; i < 5; i++ { // 在 for 代码块中定义局部变量i,仅在 for 语句中可调用
		fmt.Printf("i = %d\n", i)
	}
}

func main() {
	// testGlobalVariable()
	testLocalVariable()
	fmt.Printf("nums = %d\n", num)
}
