package main

import (
	"fmt"
	"time"
)

func process(i int, c chan bool) {
	fmt.Println("started Goroutine", i)
	time.Sleep(time.Second * 2)
	fmt.Println("Goroutine end", i)
	c <- true
}

func main() {
	num := 3
	c := make(chan bool, num)
	for i := 0; i < num; i++ {
		go process(i, c)
	}
	for i := 0; i < num; i++ {
		<-c
	}
}
