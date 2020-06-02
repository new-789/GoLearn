package main

import (
	"fmt"
	"time"
)

func hello(exitChan chan string) (str string) {
	// 此处让子线程等待五秒使其执行时间更长，看主线程是否会等待子线程执行完毕后退出
	time.Sleep(time.Second * 5)
	hello := fmt.Sprintln("~~~~~~~~~~ hello goroutine ~~~~~~~~~~")
	// 往队列中插入一个 bool 类型的数据
	exitChan <- hello
	return
}

func main() {
	// 声名一个 bool 类型的队列
	var exitChan chan string
	// 初始化队列，不指定长度
	exitChan = make(chan string)
	go hello(exitChan)
	fmt.Println("========== main thread terminate ==========")
	fmt.Println("等待子线程执行中....... ")
	// 当执行到此步骤的时候会等待子线程中往队列中插入数据，然后取出来之后才会结束执行整个程序，以此来完成与主线程的同步
	str := <-exitChan
	fmt.Printf("hello return str: %s\n", str)
}
