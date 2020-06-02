package main

import (
	"fmt"
	"sync"
	"time"
)

// 注，此处 WaitGroup 参数的指定必须为一个指针类型的数据
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine", i)
	time.Sleep(time.Second * 2)
	fmt.Println("Goroutine end", i)
	wg.Done() // 执行完一个 goroutine 时，使  WaitGroup 计数器减一
}

func main() {
	num := 3
	var wg sync.WaitGroup // 声名一个 WaitGroup 类型的变量
	for i := 0; i < num; i++ {
		wg.Add(1)          // 在开启一个 goroutine 之前使 WaitGroup 计数器加一
		go process(i, &wg) // 开启 goroutine
	}
	wg.Wait() // 检测 WaitGroup 计数器的值是否为零，是则表示所有的 goroutine 都执行完毕了
	fmt.Println("All goroutines finished executing")
}
