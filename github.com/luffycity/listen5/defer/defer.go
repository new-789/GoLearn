package main

import (
	"fmt"
)

func testDefer1() {
	defer fmt.Printf("derfer 语句\n") // 此条语句会在下面所有代码执行完毕函数返回前执行
	fmt.Printf("函数中执行顺序的第一条语句\n")
	fmt.Printf("函数中执行顺序的第二条语句\n")
}

func testDefer2() {
	defer fmt.Printf("defer 1\n")
	defer fmt.Printf("defer 2\n")
	defer fmt.Printf("defer 3\n")
	fmt.Printf("one\n")
	fmt.Printf("two\n")
	fmt.Printf("there\n")
}

func main() {
	// testDefer1()
	testDefer2()
}
