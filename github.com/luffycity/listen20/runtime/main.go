package main

import (
	"fmt"
	"runtime"
	"time"
)

var j int

func calc() {
	for {
		j++
	}
}

func main() {
	// 获取 cpu 的核心数
	cpu := runtime.NumCPU()
	fmt.Println("cpu:", cpu)
	// 设置该程序使用 cpu 的核心数
	runtime.GOMAXPROCS(cpu)

	for i := 0; i < 10; i++ {
		// 开启线程方法
		go calc()
	}
	time.Sleep(time.Hour)
}
