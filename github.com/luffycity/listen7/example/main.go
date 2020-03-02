package main

import (
	"fmt"
	"math/rand" // 导入数学包中的随机函数包
	"time"
)

func sumArray(a [10]int) int {
	var sum int
	// 第一种遍历数组方式求和
	// for i := 0; i < len(a); i++ {
	// 	sum += a[i]
	// }
	// 第二种遍历方式求和
	for _, val := range a {
		sum += val
	}
	return sum
}

func testArraySum() {
	rand.Seed(time.Now().Unix()) // 使用时间单位秒的戳初始化随机数种子，若不初始化随机种子那么它的随机种子默认为0，将导致每次随机出来的值一样
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000) // 在 0 到 1000-1 之间产生一个随机值
		// b[i] = rand.Int()      // 在 0 到 Int最大值之间产生一个随机值
	}
	sum := sumArray(b)
	fmt.Printf("sumArray = %d\n", sum)
}

func twoSum(a [5]int, target int) {
	// 定义一个函数接收两个参数，第一个参数为一个数组类型，第二个参数为 int 类型，表示数组中任意两个下标数之和
	for i := 0; i < len(a); i++ { // 从第零个元素循环数组
		ather := target - a[i]            // 两个下标之和减去其中一个数找出另一个数是多少，用来在后面判断另一个数
		for j := i + 1; j < len(a); j++ { // 循环数组，j 等于 i+1 表示循环的内容始终在 i 之前
			if a[j] == ather { // 判断另一个数
				fmt.Printf("(%d, %d)\n", i, j) // 输出两个数的下标
			}
		}
	}
}

func testTwoSum() {
	b := [5]int{1, 3, 5, 8, 7}
	twoSum(b, 8)
}

func main() {
	// testArraySum()
	testTwoSum()
}
