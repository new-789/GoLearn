package main

import (
	"fmt"
	"sync"
	"time"
)

var x int

func write(wg *sync.WaitGroup, rw *sync.Mutex) {
	for i := 0; i < 100; i++ {
		rw.Lock()
		x += 1
		time.Sleep(time.Millisecond * 10)
		rw.Unlock()
	}
	wg.Done()
}

func read(i int, wg *sync.WaitGroup, rw *sync.Mutex) {
	for i := 0; i < 100; i++ {
		rw.Lock()
		//fmt.Printf("goroutine:%d, x: %d\n", i, x)
		time.Sleep(time.Millisecond)
		rw.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	//var rw sync.RWMutex
	var mutex sync.Mutex

	start := time.Now().UnixNano()
	wg.Add(1)
	go write(&wg, &mutex)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read(i, &wg, &mutex)
	}
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println("time:", (end-start)/1000/1000, "ms")
}
