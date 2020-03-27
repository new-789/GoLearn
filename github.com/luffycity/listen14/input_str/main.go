package main

import (
	"fmt"
)

func testSscanf() {
	var a int
	var b string
	var c float32
	var str string = "8888 hello 99.9999"
	fmt.Sscanf(str, "%d%s%f", &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func testSscan() {
	var a int
	var b string
	var c float32
	// Sscan 使用空格与换行符作为分隔符，zhoufang 与 9.99 之间为windows系统默认的换行符写法
	var str string = "9999 helloZhouFang\r\n9.99"
	fmt.Sscan(str, &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func testSscanln() {
	var a int
	var b string
	var c float32
	// Sscanln 仅使用空格作为分隔符，遇到换行符则结束，定义一个字符串其中数字类型与字符串类型之间用换行符分隔
	var str string = "90\nhelloZhuFeng 99.09"
	fmt.Sscanln(str, &a, &b, &c)
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	//testSscanf()
	//testSscan()
	//testSscanln()
}
