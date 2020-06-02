package main

import (
	"fmt"
)

func write(c chan int) {
	for i := 0; i < 5; i++ { // 循环往 chan 队列中插入数据，插入两个后开始阻塞
		c <- i
	}
	close(c)
}

func main() {
	// 定义一个 chan 变量，并指定容量为 2
	c := make(chan int, 2)
	go write(c) // 调用 write 函数往队列中插入元素
	//time.Sleep(time.Second * 2)  // 此处等到两秒
	for v := range c { // 循环从 chan 队列中获取数据，取出两个后开始阻塞
		fmt.Println("read_channel_value：", v)
		//time.Sleep(time.Second * 2)  // 等待两秒等待往队列中继续插入数据下次循环可以彩瓷取出
	}
}
