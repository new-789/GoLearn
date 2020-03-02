package main

import (
	"fmt"
)

func testArray1() {
	var a [3]int // 创建一个名称为 a 的数组，且数组中的元素类型为 int 类型，长度为5
	a[0] = 10
	a[1] = 20      // 使用数组的下标特性对数组中的第零个元素进行赋值
	a[2] = 30      // 使用数组的下标特性对数组中的第 2 个元素进行赋值
	fmt.Println(a) // 不对数组做任何操作直接打印数组，此时数组中的内容默认初始化全部为 0
}

func testArray2() {
	// 定义时数组初始化
	var b [3]int = [3]int{10, 20, 30}
	fmt.Println(b)
}

func testArray3() {
	b := [3]int{10, 20, 30}
	fmt.Println(b)
}

func testArray4() {
	b := [...]int{10, 20, 30}
	fmt.Println(b)
}

func testArray5() {
	b := [3]int{10}
	fmt.Println(b)
}

func testArray6() {
	a := [5]int{0: 10, 2: 20, 4: 40}
	fmt.Println(a)
}

func testArray7() {
	a := [3]int{10, 20, 30}
	fmt.Println(a)
	var b [3]int
	b = a
	fmt.Println(b)
}

func testArray8() {
	a := [5]int{3: 100, 4: 200}
	fmt.Printf("len(a)=%d\n", len(a))
}

func testArray9() {
	a := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(a); i++ { // 遍历数组方式一
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}

func testArray10() {
	a := [5]int{10, 20, 30, 40, 50}
	for index, val := range a {
		fmt.Printf("a[%d] = %d\n", index, val)
	}
}

func testArray11() {
	var a [3][2]int // 定义一个三行两列的二维数组 a,数组中元素类型为 int
	a[0][0] = 10    // 给二维数组赋值方法
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 40
	a[2][1] = 50

	// fmt.Println(a)
	// 遍历二维数组方式一
	// for i := 0; i < len(a); i++ {
	// 	for j := 0; j < len(a[i]); j++ {
	// 		fmt.Printf("%d ", a[i][j])
	// 	}
	// 	fmt.Println() // 打印换行符
	// }
	// 遍历二维数组方式二
	for i, val := range a {
		for j, value := range val {
			fmt.Printf("(%d,%d) = %d ", i, j, value)
		}
		fmt.Println() // 打印换行符
	}
}

func strArray(a [3][2]string) {
	for _, val := range a {
		for _, value := range val {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}

func testArray12() {
	a := [3]int{10, 20, 30} // 定义一个数组 a
	b := a                  // 将数组 a 拷贝一份给 b,此时变量 b 会将属于数组 a 的数组完完整整的肤质一份，且二者所占用的内存空间是不一样的。
	b[0] = 100              // 更改数组 b 中的第 0 个元素为 1000，对 a 数组没有任何影响
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
}

func modify(b [3]int) [3]int {
	// 该函数接收一个数组类型的参数
	b[0] = 100 // 更改接收参数数组 b 的第零个元素为 100
	return b
}

func practice1() {
	var a [5]int = [5]int{1, 3, 5, 8, 7}
	var sumArray int
	for i := 0; i < len(a); i++ {
		sumArray += a[i]
	}
	fmt.Printf("sumArray = %d\n", sumArray)
}

func practice2() {
	var a [5]int = [5]int{1, 3, 5, 8, 7}
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d, %d) = %d\n", i, j, a[i]+a[j])
			}
		}
	}
}

func testArray13() {
	a := [3]int{10, 20, 30} // 定义数组 a
	b := modify(a)          // 将数组 a 当做参数传给 modify 函数，并接收返回值
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b - %v\n", b)
}

func main() {
	// testArray1()
	// testArray2()
	// testArray3()
	// testArray4()
	// testArray5()
	// testArray6()
	// testArray8()
	// testArray9()
	// testArray10()
	// testArray11()
	// 定义一个三行两列的数组 a，且数组中的元素类型为字符串,并对其进行赋值
	// a := [3][2]string{
	// 	{"hello", "world"}, // 二维数组直接定义赋值方法，大括号中同样是大括号然后在写赋值的内容
	// 	{"hello", "zhoufang"},
	// 	{"hello", "zhufeng"},
	// }
	// strArray(a)
	// testArray12()
	// testArray13()
	// practice1()
	practice2()
}
