package main

import (
	"fmt"
)

func testfor1() {
	for i := 1; i < 10; i++ {
		fmt.Printf("i = %d \n", i)
	}
}

func testfor2() {
	for i := 1; i < 10; i++ {
		if i > 5 {
			break // 当 i 的值大于 5 时直接终止 for 循环
		}
		fmt.Printf("testfor2 i=%d\n", i)
	}
}

func testfor3() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("testfor3=%d\n", i)
	}
}

func testfor4() {
	i := 0
	for i <= 10 {
		fmt.Printf("testfor4 i = %d\n", i)
		i += 2
	}
}

func testfor5() {
	for n, i := 10, 1; i <= 10 && n <= 19; i, n = i+1, n+1 {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}

func testfor6() {
	for {
		fmt.Printf("这是一个无限循环\n")
	}
}

func main() {
	testfor1()
	testfor2()
	testfor3()
	testfor4()
	testfor5()
	testfor6()
}
