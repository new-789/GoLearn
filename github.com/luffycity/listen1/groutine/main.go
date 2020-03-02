package main

import (
	"fmt"
	"time"
)

func calc() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("run", i, "times")
	}
	fmt.Println("calc finish")
}

// 若在 main 函数中开启线程后，主进程结束则所有子线程也都会跟着结束
func main() {
	go calc() // 开启线程方法
	fmt.Println("i exited")
	time.Sleep(11 * time.Second)
}
