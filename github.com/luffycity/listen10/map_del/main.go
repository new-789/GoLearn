package main

import (
	"fmt"
)

func main() {
	// 声名一个 map 类型的变量 a 并进行初始化
	var a map[string]int = map[string]int{
		"stu01": 1000,
		"stu02": 2000,
		"stu03": 3000,
		"stu04": 4000,
		"stu05": 5000,
	}
	fmt.Printf("原始 map 数据:%#v\n", a)
	//// 声名一个变量 del 用来保存需要删除的 key
	//var del string = "stu04"
	//delete(a, del) // 删除记录，也可以直接使用 delete(a, "stu04") 语法
	//fmt.Printf("删除其中一个元素之后的数据:%#v\n", a)

	// 使用 for 循环删除 map 中的所有数据
	for key, _ := range a {
		delete(a, key)
	}
	fmt.Printf("删除map中所有数据后:%#v\n", a)
}
