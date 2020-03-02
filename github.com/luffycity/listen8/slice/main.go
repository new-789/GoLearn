package main

import (
	"fmt"
)

func testSlice0() {
	var a []int   // 定义一个空切片 a
	if a == nil { // nil 判断一个数组是否为空切片
		fmt.Printf("a is null\n")
	} else {
		fmt.Printf(" a = %v\n", a)
	}
}

func testSlice1() {
	a := [5]int{1, 2, 3, 4, 5} // 创建一个数组 a
	var b []int                // 创建一个空切片 b
	b = a[1:4]                 // 借助数组 a 对空切片 b 进行初始化
	fmt.Printf("slice b: %v\n", b)
	fmt.Printf("b[0] = %d\n", b[0]) // 获取切片中的第零个元素
	fmt.Printf("b[1] = %d\n", b[1]) // 获取切片中的第1个元素
	fmt.Printf("b[2] = %d\n", b[2]) // 获取切片中的第2个元素
}

func testSlice2() {
	a := []int{6, 7, 8}
	fmt.Printf("slice a：%v\n type of a:%T\n", a, a)
}

// 数组切片基本操作示例
func testSlice3() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("a = %v\n", a)
	var b []int
	b = a[1:9] // 切片语法一，从第1个元素到9-1(即倒数第二)个元素
	fmt.Printf("b = %v\n", b)
	c := a[1:] // 切片语法二从第一个元素到结尾
	fmt.Printf("c = %v\n", c)
	e := a[:9] // 切片语法三，从第0个元素到 9-1(即倒数第二) 个元素
	fmt.Printf("e = %v\n", e)
	f := a[:] // 切片语法四，包含所有元素的切片
	fmt.Printf("f = %v\n", f)
}

func testSlice4() {
	// 创建一个数组 a， [...]表示有编译器确定数组的长度
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	fmt.Printf("array a:%v, type of %T\n", a, a)
	b := a[2:6] // 创建切片b，并对数组a进行切片初始化切片b
	fmt.Printf("[]int b:%v, type of: %T\n", b, b)
	// 修改切片操作方式一
	// b[0] += 10 // 修改切片中的第0个元素
	// b[1] += 20 // 修改切片中的第1个元素
	// b[2] += 30 // 修改切片中的第2个元素
	// b[3] += 40 // 修改切片中的第3个元素
	// 修改切片操作方式二，使用 for 循环
	for index := range b { // 使用 for range 循环遍历切片，只需要下标索引所以我们只保留了一个 index，正常写法应该是 for index,val := range b {}
		b[index] += 10 // 修改切片中对应下标的元素
	}
	fmt.Printf("ayyar a : %v, type of : %T\n", a, a)
}

func main() {
	// testSlice0()
	// testSlice1()
	// testSlice2()
	// testSlice3()
	testSlice4()
}
