package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func testFunc1() {
	f1 := add                         // 将 add 函数赋予给 f1 变量
	fmt.Printf("f1 type is %T\n", f1) // 打印 f1 的类型，打印类型用 %T 表示
	sum := f1(10, 50)
	fmt.Printf("sum=%d\n", sum)
}

func testFunc2() {
	f1 := func(a, b int) int { // 定义一个匿名函数方法，将匿名函数赋值给 f1 变量
		return a - b
	}
	fmt.Printf("type of f1=%T\n", f1)
	sum := f1(30, 20) // 通过 f1 调用 匿名函数
	fmt.Printf("sum=%d\n", sum)
}

func testFunc3() {
	var num int = 88
	defer func() { // 在 defer 后直接定义一个匿名函数，表示函数返回前会执行该匿名函数
		fmt.Printf("num=%d\n", num)
	}() // 在匿名函数后加上括号标志直接调用该匿名函数
	num = 99
	fmt.Printf("num=%d\n", num)
	return
}

func sub(a, b int) int {
	return a - b
}

func ride(a, b int32) int32 {
	return a * b
}

func calc(a, b int, op func(int, int) int) int {
	/* op 参数是一个函数类型的参数，接收一个函数类型的参数
	   注：在传递函数做为参数时，定义的函数参数如果有参数，
		   那么被传递的函数的参数类型必须一致, 如 op 作为一个函数参数，
		   接收两个 int 类型的数据，如果传递给它的函数所接受的参数不符合 int 类型则报错；
	*/
	return op(a, b)
}

func testFunc4() {
	sum := calc(100, 200, add) // 调用 calc 函数并将 add 函数当做参数传递给 calc 的第三个参数
	sub := calc(300, 100, sub) // 调用 calc 函数并将 sub 函数当做参数传递给 calc 的第三个参数
	// 调用 calc 函数并将 ride 函数当做参数传递给 calc 的第三个参数,由于 ride 接收的参数类型与 calc 函数中第三个参数 op 接收的参数类型不同所以报错
	// ride := calc(10, 10, ride)
	fmt.Printf("sum=%d, sub=%d\n", sum, sub)
}

func main() {
	// testFunc1()
	// testFunc2()
	// testFunc3()
	testFunc4()
}
