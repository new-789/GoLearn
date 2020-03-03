package main

import (
	"fmt"
)

func modify(a *int) {
	// 声名一个 modify 函数，接收一个指针类型的参数
	*a = 100
}

func testExample() {
	var b int = 10 // 声名一个 int 类型的值变量 b
	fmt.Printf("修改之前的值:%d\n", b)
	modify(&b) // 调用 modify 函数，并传入值变量 b 的内存地址
	fmt.Printf("修改之后的值:%d\n", b)
}

func main() {
	testExample()
}
