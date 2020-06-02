package main

import "fmt"

// 声名一个 sendData 函数用来往队列中插入数据
func sendData(s chan<- int) {
	/* s chan<- int 表示该队列只能往队列中插入数据*/
	s <- 100
}

// 声名一个 readData 函数用来从队列中读取数据
func readData(r <-chan int) {
	/* r <-chan int 表示该队列只能从队列中获取数据*/
	data := <-r
	fmt.Println("readData：", data)
}

func main() {
	// 声名一个管理变量并进行初始化
	c := make(chan int)
	go sendData(c)
	readData(c)
}
