package main

import (
	"fmt"
	"time"
)

func write(ch chan string) {
	for {
		select {
		case ch <- "hello": // 通过 select 语句往管道中插入数据
			fmt.Println("write success")
		default:
			fmt.Println("channel is full")
		}
		time.Sleep(time.Millisecond * 500) // 等待五百毫秒
	}
}

func main() {
	select {} // 当代码执行到该行时，程序产生阻塞；
	out1 := make(chan string, 10)

	go write(out1)

	for s := range out1 {
		fmt.Println("recv: ", s)
		time.Sleep(time.Second) // 每等待一秒从 out1 管道中取一个数据
	}
}
