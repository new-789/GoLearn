package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		// 判断 n 是否为一个素数，可以整除则说明不是素数,判断一个数是否为素数的关键点在于这个数是否可以被任何数整除
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 声名一个生产者函数
func produceSushu(c chan int) {
	var i int = 1
	for {
		i += 1
		result := isPrime(i)
		if result {
			// 将元素插入队列语法
			c <- i
		}
		time.Sleep(time.Second)
	}
}

// 声名一个消费者函数
func consumeSushu(c chan int) {
	// 通过 for 循环从队列中取元素方法
	for v := range c {
		fmt.Printf("%d is prime\n", v)
	}
}

func main() {
	// 声名一个队列类型的变量，并通过 make 关键字初始化队列容量为 1000，声名队列使用 chan 关键字
	var intChan chan int = make(chan int, 1000)
	// go 语言开启线程方法
	go produceSushu(intChan)
	go consumeSushu(intChan)

	time.Sleep(time.Hour)
}
