package main

import (
	"fmt"
	"time"
)

// 声名一个生产者函数，用来往不带缓存区的队列中插入数据
func produce(c chan int) {
	c <- 10000
	fmt.Println("=============== produce start running ===============")
}

// 声名一个消费者函数，用来获取队列中的数据
func consume(c chan int) {
	data := <-c
	fmt.Println(data)
}

func main() {
	var c chan int
	fmt.Println(c)
	c = make(chan int) // 初始化队列，不指定队列的长度
	go produce(c)      // 使用 go 关键字开启子线程执行 produce 函数
	go consume(c)      // 使用 go 关键字开启字线程执行 consume 函数
	// 等待两秒让子线程执行完毕再结束，否则主线程执行完毕则开启的子线程会全部结束
	time.Sleep(time.Second * 2)
}
