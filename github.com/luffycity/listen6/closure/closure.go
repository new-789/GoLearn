package main

import (
	"fmt"
	"strings"
	"time"
)

// Adder 函数的返回值是一个匿名函数
func Adder() func(int) int {
	var num int // Ader 函数中定义的局部变量 num
	// 定义一个匿名函数接收一个int类型的参数且返回一个 int 类型的参数，并将其赋值给 closure 变量
	closure := func(n int) int {
		num += n   // 在匿名函数内部调用它上级函数中的局部变量 num
		return num // 将局部变量 num 返回
	}
	return closure // 返回 closure 变量，实际上等同于返回一个匿名函数
}

// 定义一个 add 函数接收一个参数，并返回一个函数类型的数据且返回的函数类型接受一个参数和返回一个值
func add(base int) func(int) int {
	// 返回一个匿名函数
	return func(num int) int {
		base += num // 在匿名函数内部调用上级函数的参数，同样为闭包函数
		return base
	}
}

func testClosure() {
	// 调用 Adder 函数，并接受返回的一个函数类型的返回值，此时 closure 实际上是一个函数，可使用函数的调用方式进行调用
	closure := Adder()
	// 调用返回的 closure 变量指向的匿名函数并进行传参，并将返回的值赋予 num 变量，以次便实现了外部访问函数内部局部变量的实现
	num := closure(1)
	fmt.Printf("num=%d\n", num)
	num = closure(100)
	fmt.Printf("num=%d\n", num)
	num = closure(300)
	fmt.Printf("num=%d\n", num)

	closure1 := Adder() // 重新调用 Adder 函数，创建了一个新环境,此时 Adder 函数内部的 num 局部变量重新初始化
	num = closure1(1)   // 相当于修改 Adder 函数内部局部变量 num 的值
	fmt.Printf("num=%d\n", num)
}

func testClosure1() {
	tmp1 := add(10) // 调用 add 函数接收一个匿名函数，并将该匿名函数赋值给 tmp1 变量
	// f(10) 相当于通过变量调用 add 函数内部的闭包函数并传参
	fmt.Printf("第一次调用返回的值%d, 第二次调用返回的值%d\n", tmp1(10), tmp1(20))
	tmp2 := add(100) // 重新调用 add 函数生成一个新的闭包实例，并重新传值
	fmt.Printf("第一次调用返回的值%d, 第二次调用返回的值%d\n", tmp2(100), tmp2(200))
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func testClosure2() {
	func1 := makeSuffixFunc(".bmp\n")
	func2 := makeSuffixFunc(".jpg\n")
	fmt.Printf(func1("test"))
	fmt.Printf(func2("test"))
}

// 定义一个 calc 函数，接收一个参数，并同时返回两个匿名函数
func calc(base int) (func(int) int, func(int) int) {
	// 函数内部定义一个匿名函数并赋值给 sum 变量，同时该匿名函数接收且并返回一个参数
	sum := func(i int) int {
		base += i   // 匿名函数内调用上级 calc 函数的参数做计算
		return base // 返回 base 参数
	}
	// 函数内部定义一个匿名函数并赋值给 sub 变量， 同时该匿名函数接收且并返回一个参数
	sub := func(n int) int {
		base -= n   // 匿名函数内调用上级 calc 函数的参数做计算
		return base // 返回 base 参数
	}
	return sum, sub // 同时返回两个类型为匿名函数的变量
}

func testClosure3() {
	// 调用 calc 函数并传参，同时定义两个变量接收 calc 函数返回的两个匿名函数
	sum, sub := calc(100)
	// sum(788) 表示调用 calc 返回的内部匿名函数(闭包)并传值
	fmt.Printf("one sum closure value: %d\n", sum(788))
	fmt.Printf("one sub closure value: %d\n", sub(12))
	// 再次调用 calc 函数并传参即重新创建一个环境，同时定义两个变量接收 calc 函数返回的两个匿名函数
	sum, sub = calc(100)
	// sum(788) 表示调用 calc 返回的内部匿名函数(闭包)并传值
	fmt.Printf("two sum closure value: %d\n", sum(566))
	fmt.Printf("two sub closure value: %d\n", sub(658))
}

func testClosure4() {
	for i := 0; i < 5; i++ { //
		// 在函数内部开启异步，并执行一个匿名函数(闭包)
		final := i
		go func() {
			// 注 i 属于父级函数中的 for 代码块的变量
			fmt.Println(final)
		}()
	}
	time.Sleep(time.Second)
}

func main() {
	// testClosure()
	// testClosure1()
	// testClosure2()
	// testClosure3()
	testClosure4()
}

// 注：若定义的函数提示 exported function Adder should have comment or be unexported 则需要为该函数添加一行注释，且注释的开头必须是该函数名
