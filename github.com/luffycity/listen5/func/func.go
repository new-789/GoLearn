package main

import (
	"fmt"
)

func sayheloo() { // 定义一个无参数和无返回值函数，并在其中打印 hello sorld
	fmt.Printf("hello world\n")
}

func funcAdd(a int, b int) int {
	/*
		a，b：即为接受的函数
		int: 参数后的 int 表示接受参数的类型
		int: 括号与大括号之间的 int 表示返回值的类型
	*/
	sum := a + b
	return sum
	// 我们还可以直接返回两个数相加，如 return a + b
}

func testFunc(a, c, b int) int {
	return a + b + c
}

func testreturn(vara, varb int) (int, int) {
	sum := vara + varb
	sub := vara - varb
	return sum, sub
}

func testReturn(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a + b
	return
}

func calc_v1(a ...int) int {
	sum := 0
	// 写法一
	for _, val := range a {
		sum += val
	}
	// 写法二
	// for i := 0; i < len(a); i++ {
	// 	sum += a[i]
	// }
	return sum
}

func calc_v2(a int, b ...int) int {
	sum := 0
	for _, val := range b {
		sum += val
	}
	return sum
}

func main() {
	// sayheloo() // 调用函数：通过函数名加括号即可调用该函数
	// sum := funcAdd(200, 688)
	// fmt.Printf("相加后的结果等于:%d\n", sum)
	// fmt.Printf("返回结果：%d\n", testFunc(666, 222, 112))
	// sum, sub := testreturn(222, 111) // 调用函数并同时接收多个值语法
	// sum, _ := testReturn(222, 111)
	// fmt.Printf("sum=%d\n", sum)
	// sum := calc_v1()                            // 不传参
	// sum := calc_v1(100, 200, 300, 400, 500)
	sum := calc_v2(10, 20, 20, 20)
	fmt.Printf("sum = %d\n", sum)
}
