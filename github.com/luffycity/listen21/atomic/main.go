package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int32

func add(wg *sync.WaitGroup) {
	for i := 0; i < 50000; i++ {
		// 原子增操作,参数一为需要增加的类型变量，参数二为需要增加的值
		atomic.AddInt32(&x, 1)
	}
	wg.Done()
}

func addMutex(wg *sync.WaitGroup, mutex *sync.Mutex) {
	for i := 0; i < 50000; i++ {
		mutex.Lock()
		x += 1
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	start := time.Now().UnixNano()
	//for i := 0; i < 1000; i++{
	//	wg.Add(1)
	//	go add(&wg)
	//}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addMutex(&wg, &mutex)
	}
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println("x=", x, "time:", (end-start)/1000/1000, "ms")
}
