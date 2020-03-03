package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	// 定义一个 swap 函数并接收两个指针类型的数据类型为  int 的参数
	fmt.Printf("before:a=%d , b=%d\n", *a, *b)
	*a, *b = *b, *a // 通过指针交换两个变量的值
	fmt.Printf("after:a=%d , b=%d\n", *a, *b)
}

func main() {
	var a int = 88
	var b int = 98
	swap(&a, &b) // 获取两个值变量指向的内存地址并当做参数传递给 swqp 函数
	fmt.Printf("in main:a=%d , b=%d\n", a, b)
}
