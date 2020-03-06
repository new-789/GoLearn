package main

import "fmt"

type Test struct {
	A int32 // 一个 int32 类型的数据占用四个字节的内存空间
	B int32
	C int32
	D int32
}

func main() {
	var test Test
	fmt.Printf("a address:%p\n", &test.A)
	fmt.Printf("b address:%p\n", &test.B)
	fmt.Printf("C address:%p\n", &test.C)
	fmt.Printf("D address:%p\n", &test.D)
}
