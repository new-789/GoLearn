package main

import (
	"fmt"
)

func main() {
	// 定义一个队列 c
	var c chan int
	fmt.Printf("c=%v\n", c)
	// 初始化队列 c，并制定长度为 10
	c = make(chan int, 10)
	fmt.Printf("c=%v\n", c)
	// 入队操作
	c <- 100
	// 出队操作
	data := <-c
	fmt.Printf("data=%d\n", data)
	<-c // 出队并丢弃数据,在 go 1.14 版本中如果队列中只有一个数据做此操作在执行时则会抛出异常
}
