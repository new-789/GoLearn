package main

import (
	"fmt"
)

func testPoint1() {
	var a int32 = 100
	var b *int32  // 定义一个指针类型的变量语法
	if b == nil { // 判断一个指针变量是否为空语法(有没有感觉和判断数组为空语法一样)
		fmt.Printf("b is nil addr\n")
	}
	b = &a // 获取变量 a 的内存地址赋值给变量 b
	fmt.Printf("b is addr:%p, b is value: %v\n", &b, b)
}

func testPoint2() {
	var a int = 100 // 定义一个 int 类型的变量 a，并初始化值为 100
	var b *int = &a // 定义一个 Int 类型的指针变量 b，并初始化指向变量 a 的内存地址
	*b = 200        // 通过指针变量语法 *b 更改指向内存地址中存储的值
	fmt.Printf("b指针变量指向内存地址中存储的值：%d\n", *b)
	fmt.Printf("a is value:%d\n", a) // 再次打印变量 a 的内容，同样被更改为 200
}

func modify(point *int) {
	// 定义 testPoint3 函数，接收一个指针类型的传参
	*point = 88 // 更改指针类型参数 point 接收到的值为 88
}

func testPoint3() {
	var a int = 100                               // 声名一个 int 类型的变量 a，并初始化为 100
	var b *int = &a                               // 声名一个 int 类型的指针类型变量 b， 并初始化为变量 a 的内存地址
	modify(b)                                     // 调用函数并将指针类型变量 b 当做参数传递, 相当于将指针变量 b 的内存地址拷贝一份给函数中定义的参数，即 point := &b
	fmt.Printf("a value:%d, b value:%d\n", a, *b) // 函数中修改指针参数值后输出变量 a 与指针变量 b 相对应的值
}

func modifyArray(arr *[3]int) {
	// 定义函数接收一个指针变量的数组作为参数
	(*arr)[1] = 88 // 更改数组中的值,注意指针参数获取数组内容语法
}

func testPoint4() {
	// 定义一个类型为 int 长度为 3 的数组 ar
	ar := [...]int{78, 89, 98}
	// 调用 modifyArray 函数，并获取数组 ar 的内存地址当做参数传递给函数
	modifyArray(&ar)
	// 查看在函数中将指针变量数组中的元素更改之后原 ar 数组中的内容
	fmt.Printf("ar 数组更改之后的值:%v\n", ar)
}

func testPoint5() {
	// 创建一个指针型类型的变量 a ，且设定为 int 类型，new(int) 表示为该指针变量创建一个内存地址，这样就可以直接对指针变量像值类型变量那样进行操作了
	var a *int = new(int)
	*a = 100 // 直接操作指针变量，为其赋值 100
	fmt.Printf("*a=%d\n", *a)

	// 使用 new 语法创建一个指针类型的空切片
	var b *[]int = new([]int)
	fmt.Printf("*b=%v\n", *b)
	//(*b)[0] = 100  由于指针切片为空且指针类型的变量指向的是内存地址，所以不支持直接操作赋值
	(*b) = make([]int, 5, 10) // 将指针切片使用 make 方法进行初始化
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("*b=%v\n", *b)
}

func main() {
	//testPoint1()
	//testPoint2()
	//testPoint3()
	//testPoint4()
	testPoint5()
}
