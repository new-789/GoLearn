package main

import (
	"fmt"
	"sync"
	"time"
)

var x int

func write(rw *sync.RWMutex, wg *sync.WaitGroup) {
	rw.Lock() // 给写操作加锁操作,当一个线程获取写锁之后其它线程无论是读操作还是写操作都处于等待状态
	fmt.Println("write Lock")
	x += 1
	time.Sleep(time.Second * 10)
	rw.Unlock() // 释放写锁
	fmt.Println("write UnLock")
	wg.Done()
}

func read(i int, rw *sync.RWMutex, wg *sync.WaitGroup) {
	fmt.Println("wait for rlock")
	rw.RLock() // 给读操作加读锁,当一个线程获取到读锁后，写操作会处于等待状态，读操作则是并行的
	fmt.Printf("goroutine:%d, x:%d\n", i, x)
	rw.RUnlock() // 释放读锁
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var rw sync.RWMutex
	// 通过循环启 10 个写线程
	wg.Add(1)
	go write(&rw, &wg)
	time.Sleep(time.Millisecond * 5)

	for i := 0; i < 10; i++ { // 通过循环启 10 个读线程
		wg.Add(1)
		go read(i, &rw, &wg)
	}

	wg.Wait()
}
