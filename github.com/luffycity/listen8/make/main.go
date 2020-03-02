package main

import (
	"fmt"
)

func testMake1() {
	// 使用 make 创建一个切片，第一个参数表示是一个切片类型的数组，第二个参数表示切片的长度，第三个参数表示切片的容量
	a := make([]int, 5, 10)
	a[0] = 10
	a[1] = 20
	fmt.Printf("a = %v\n", a)
}

func testMake2() {
	var a []int                                                              // 定义一个空切片
	a = make([]int, 1, 5)                                                    // 使用 make 定义切片的长度和容量
	a[0] = 10                                                                // 给切片的第 0 个元素增加内容
	fmt.Printf("a = %v, address:%p, len:%d, cap:%d\n", a, a, len(a), cap(a)) // 查看切片的内容，长度和容量
	for i := 11; i < 15; i++ {                                               // 使用 for 循环对切片进行长度的扩充
		a = append(a, i) // 往切片中不断的增加元素
	}
	fmt.Printf("长度扩容后：a = %v, address:%p, len:%d, cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 20) // 当超过容量后继续往切片中增加元素，此时会扩容切片的容量
	fmt.Printf("容量扩容后：a = %v, address:%p, len:%d, cap:%d\n", a, a, len(a), cap(a))
}

func testMake3() {
	a := make([]int, 1, 5)
	a = append(a, 10)
	fmt.Printf("a = %v\n", a)
}

func testMake4() {
	a := make([]int, 0, 5)
	a = append(a, 10)
	fmt.Printf("a = %v\n", a)
}

func testCap1() {
	// 定义一个字符串类型的数组，并初始化8个元素
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := a[1:3] // 定义一个切片 b，初始化数据来自数组 a 的切片
	fmt.Printf("b = %v, len(b) = %d, cap(b) = %d\n", b, len(b), cap(b))
}

func testCap2() {
	// 定义一个字符串类型的数组，并初始化8个元素
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := a[1:3] // 定义一个切片 b，初始化数据来自数组 a 的切片
	fmt.Printf("b = %v, len(b) = %d, cap(b) = %d\n", b, len(b), cap(b))
	c := b[:cap(b)] // 对切片 b 在进行切片，此时切片 c 的长度和容量均来自切片 b 的容量，因为切片的内容是从切片 b 的第零个元素开始至所有容量结束
	fmt.Printf("对切片b在进行切片：c = %v, len(c) = %d, cap(c) = %d\n", b, len(c), cap(c))
}

func testSlice() {
	var a []int // 定义一个空切片
	fmt.Printf("address:%p, len:%d, cap:%d\n", a, len(a), cap(a))
	if a == nil { // nil 判断一个切片是否空切片方法
		fmt.Printf("slice a is nil\n")
	}
	a = append(a, 100) // 往空切片中增加元素方法
	fmt.Printf("address:%p, len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 200) // 往空切片中增加元素方法
	fmt.Printf("address:%p, len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 300) // 往空切片中增加元素方法
	fmt.Printf("address:%p, len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 300) // 往空切片中增加元素方法
	fmt.Printf("address:%p, len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 400) // 往空切片中增加元素方法
	fmt.Printf("address:%p, len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 500) // 往空切片中增加元素方法
	fmt.Printf("address:%p, len:%d, cap:%d\n", a, len(a), cap(a))
}

func testAppend() {
	var a []int = []int{1, 3, 4}
	var b []int = []int{4, 5, 6}
	a = append(a, 6, 7, 8)
	fmt.Printf("a = %v\n", a)
	// b 后面的三个点表示将切片 b 展开为一个个元素
	a = append(a, b...)
	fmt.Printf("a = %v\n", a)
}

func sumArray(a []int) int {
	// 定义一个函数接收一个切片
	var sum int
	for _, v := range a { // 循环该切片
		sum += v
	}
	return sum
}

func testSumArray() {
	// 定义一个长度为 10 的数组 a
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 调用 sumArray 函数并将数组 a 进行切片取所有元素当做参数传递给 sumArray 函数
	sumA := sumArray(a[:])
	fmt.Printf("sumA = %d\n", sumA) // 打印结果
	// 定义一个长度为 3 的数组 b
	var b [3]int = [3]int{4, 5, 6}
	// 调用 sumArray 函数并将数组 b 进行切片取所有元素当做参数传递给 sumArray 函数
	sumB := sumArray(b[:])
	fmt.Printf("sumB = %d\n", sumB) // 打印结果
}

func testModifySlice(a []int) {
	a[0] = 1000
}

func modifySlice() {
	var a [3]int = [3]int{1, 2, 3}
	testModifySlice(a[:])
	fmt.Printf("a = %v\n", a)
}

func testModifyCopy() {
	a := []int{1, 2, 3}
	b := []int{4}
	copy(a, b) // 将 b 切片拷贝到切片 a
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
}

func main() {
	// testMake1()
	// testMake2()
	// testMake3()
	// testMake4()
	// testCap1()
	// testCap2()
	// testSlice()
	// testAppend()
	// testSumArray()
	// testModifyCopy()
	modifySlice()
}
