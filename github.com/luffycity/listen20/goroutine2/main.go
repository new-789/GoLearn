package main

import (
	"fmt"
)

func Hello1(i int) {
	fmt.Println("hello goroutine", i)
}

func main() {
	// 使用 for 循环启动多个线程
	for i := 0; i < 10; i++ {
		go Hello1(i)
	}
}
