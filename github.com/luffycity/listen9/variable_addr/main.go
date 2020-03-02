package main

import (
	"fmt"
)

func main() {
	var variable int32 // 创建一个 int32 类型的变量即创建四字节的内存地址控件
	variable = 100
	fmt.Printf("variable:%d\n", variable)
	fmt.Printf("variable_addr:%p\n", &variable) // 使用 & 符号获取变量的内存地址，使用 %p 打印输出内存地址
}
