package main

import (
	"fmt"
)

// 声名一个 produce 生产者函数用来往管道中插入数据
func produce(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c) // 关闭管道语法
}

// 声名一个 consume 消费者函数用从往管道中获取数据
func consume(c chan int) {
	// 使用死循环从管道中获取数据
	for {
		// 从管道中获取数据，返回两个值，一个为插入管道的值，一个为状态类型的 bool,如果为 false 则表示管道被关闭了
		v, ok := <-c
		if ok == false { // 如果 ok 等于 false 则管道被关闭则退出循环
			fmt.Println("====== channel id closed ======")
			break
		}
		fmt.Println("ReadChan ====", v, ok)
	}
}

func main() {
	c := make(chan int)
	go produce(c)
	consume(c)
}
