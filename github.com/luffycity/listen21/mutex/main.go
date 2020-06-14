package main

import (
	"fmt"
	"sync"
)

var x int

func add(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x += 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add(&wg, &mutex) // 注：此处传递的参数为地址类型，因为 WaitGroup 以及 Mutex 为一个结构体
	}
	wg.Wait()
	fmt.Println("x:", x)
}
