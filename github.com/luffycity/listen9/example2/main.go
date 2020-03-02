package main

// 使用 golang 标准包 "sort" 对数组进行排序

import (
	"fmt"
	"sort"
)

func testIntSort(slice []int) []int {
	// sort.Ints 方法函数接收一个整数切片类型的数据作为参数，它的作用是用来对一个 int 类型的数组按照从小到大的顺序进行排序
	sort.Ints(slice)
	return slice
}

func intSort() {
	var a [5]int = [5]int{5, 2, 3, 4, 1}
	b := testIntSort(a[:])
	fmt.Printf("array a:%v\n", b)
}

func testStringSort(str []string) []string {
	// sort.Strings 方法函数接收一个字符串切片类型的数据作为参数，它的作用是用来对一个字符串类型的数组按照从a到z的顺序进行排序
	sort.Strings(str)
	return str
}

func strSort() {
	var a [5]string = [5]string{"ds", "bg", "aj", "ek", "co"}
	b := testStringSort(a[:])
	fmt.Printf("string sort a:%v\n", b)
}

func testFloatSort(float []float64) []float64 {
	// sort.Float64s方法函数接收一个浮点切片类型的数据作为参数，它的作用是用来对一个 float 类型的数组按照从小到大的顺序进行排序
	sort.Float64s(float)
	return float
}

func floatSort() {
	var a [5]float64 = [5]float64{9.234, 4.334, 0.123, 6.456, 7.234}
	b := testFloatSort(a[:])
	fmt.Printf("float sort a:%v\n", b)
}

func main() {
	// a := [5]int{5, 2, 3, 4, 1}
	// // sort.Ints 方法函数接收一个整数切片类型的数据作为参数，它的作用是用来对一个 int 类型的数组按照从小到大的顺序进行排序
	// sort.Ints(a[:]) // Ints == func(a []int)
	// fmt.Printf("int sort is a:%v\n", a)

	// var b [5]string = [5]string{"ds", "bg", "aj", "ek", "co"}
	// // sort.Strings 方法函数接收一个字符串切片类型的数据作为参数，它的作用是用来对一个 str 类型的数组按照首字母从 a 到 z 的顺序进行排序
	// sort.Strings(b[:]) // strings == func(a []string)
	// fmt.Printf("string sort is b:%v\n", b)

	// var c [5]float64 = [5]float64{9.234, 4.334, 0.123, 6.456, 7.234}
	// // sort.Float64 方法函数接收一个浮点切片类型的数据作为参数，它的作用是用来对一个 float64 类型的数组按照首字母从大到小的顺序进行排序
	// sort.Float64s(c[:])
	// fmt.Printf("float64 sort is c:%v\n", c)
	intSort()
	strSort()
	floatSort()
}
