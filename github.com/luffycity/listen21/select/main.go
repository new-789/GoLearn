package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second) // 此处等待六秒往管道中仍一个字符串
	ch <- "from sever1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second) // 此处等待三秒往管道中仍一个字符串
	ch <- "from server2"
}

func main() {
	out1 := make(chan string)
	out2 := make(chan string)

	go server1(out1)
	go server2(out2)

	//s1 := <- out1
	//fmt.Printf("s1: %s\n", s1)
	//s2 := <- out2
	//fmt.Printf("s2:%s\n", s2)  // 按照常理 s2 应该先输出，因为 server2 函数只等待了三秒钟

	select {
	case s1 := <-out1:
		fmt.Println("s1:", s1)
	case s2 := <-out2:
		fmt.Println("s2:", s2)
	default:
		fmt.Println("run default")
	}
}
