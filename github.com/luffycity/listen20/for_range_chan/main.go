package main

import (
	"fmt"
)

// 声名 produce 生产者函数往管道中插入数据
func produce(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

// 声名 consume 消费者函数从管道中获取数据
func consume(c chan int) {
	// 使用 for range 循环管道取出管道中的数据
	for v := range c {
		fmt.Println("ReadChan", v)
	}
}

func main() {
	var c chan int
	c = make(chan int)
	go produce(c)
	consume(c)
}
