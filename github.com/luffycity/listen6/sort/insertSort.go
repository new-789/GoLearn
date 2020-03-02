package main

import (
	"fmt"
)

func insertSort(ar [11]int) [11]int {
	// i := 1 表示数组的第二个数，如果 i 小于整个数组的长度，i++ 则不断加一
	for i := 1; i < len(ar); i++ {
		// 将 i 的值赋予给 j 表示 index,如果 j 大于 0 则 j 减 1，只要 j 大于 0  则一直在该层循环
		for j := i; j > 0; j-- {
			if ar[j] < ar[j-1] { // 如果 ar[j] 当前循环的数比它左边的数小则进行交换
				ar[j], ar[j-1] = ar[j-1], ar[j]
			} else { // 如果 ar[j] 不小于它左边的数则表示顺序正确，则终止循环
				break // 退出当前 for 循环，再次进入外部 for 循环
			}
		}
	}
	return ar
}

func choiceSort(a [11]int) [11]int {
	for i := 0; i < len(a); i++ { // 第一次循环数组，从第一个元素开始
		for j := i + 1; j < len(a); j++ { // 第二次循环数组，从第二个元素开始
			// 一直循环，不断将当前元素与前一个元素进行对比，后一个元素小于第一个元素则进行交换,否则进入下一次循环最后一个元素和第二个元素进行对比依次类推
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

// bubbleSort 冒泡排序
func bubbleSort(a [11]int) [11]int {
	// 第一次循环数组，当 i 小于数组的整个长度则 i 不断加一
	for i := 0; i < len(a); i++ {
		// 第二次循环数组，len(a)-i 表示配合 i 不断缩减数组的长度，再减 1 表示当前循环的前一个元素
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] < a[j+1] { // 使用当前元素与它前一个元素进行对不，小于则进行交换，大于则进入下一次循环
				a[j], a[j+1] = a[j+1], a[j] // 交换两个数
			}
		}
	}
	return a
}

func main() {
	// Go 语言中定义数组语法 [11] 表示数组的长度，int 表示数组中值的类型
	var numarray [11]int = [11]int{6, 10, 1, 3, 8, 5, 2, 4, 9, 0, 7}
	fmt.Printf("原始数组：%d\n", numarray)
	// insertArray := insertSort(numarray)
	// choiceArray := choiceSort(numarray)
	bubbleArray := bubbleSort(numarray)
	fmt.Printf("选择排序之后的数组：%d\n", bubbleArray)
}
