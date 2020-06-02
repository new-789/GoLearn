package main

import (
	"fmt"
	"runtime"
	"time"
)

func Hello() {
	fmt.Println("hello goroutine")
}

func main() {
	/*
		使用 go 关键字创建一个线程方法，创建一个 goroutine，当主线程结束则使用 go 创建的线程全部结束
		所以为了防止创建的线程能够执行完毕需要防止主线程在其它线程执行完毕前结束
	*/
	go Hello()
	fmt.Println("main thread terminate", runtime.NumCPU())
	time.Sleep(time.Second)
}
