package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 3)
	c <- "hello,"
	c <- "world!"
	fmt.Println("capacity is:", cap(c)) // 获取并输出 chan 的容量
	fmt.Println("length is:", len(c))   // 获取并输出 chan 的长度
	fmt.Println("read value:", <-c)
	fmt.Println("new length is:", len(c)) // 取出一个元素后再次获取并输出 chan 的长度
	fmt.Println("capacity is:", cap(c))   // 取出一个元素后再次获取并输出 chan 的容量
}
