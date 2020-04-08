package main

import "fmt"

func test1() int {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func test2() (x int) {
	defer func() {
		x += 1
	}()
	return 5
}

func test3() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func test4() (x int) {
	defer func(x int) {
		x += 1
	}(x) // 自调用并自传参方法
	return 5
}

func main() {
	//fmt.Println(test1())
	//fmt.Println(test2())
	//fmt.Println(test3())
	fmt.Println(test4())
}
