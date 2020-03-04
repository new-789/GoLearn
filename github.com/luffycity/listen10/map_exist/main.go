package main

import (
	"fmt"
)

func testMapExist() {
	var a map[string]int = map[string]int{
		"stu01": 100,
		"stu02": 200,
		"stu03": 300,
	}
	// 定义一个变量用来存储需要判断的 key
	var key string = "stu02"
	// 通过该语法获取 map 中 key 对应的值与返回的状态
	value, state := a[key]
	// 直接判断返回的状态是否为 true，也可以直接使用 if state {...} 语法
	if state == true {
		fmt.Printf("the's a[%s] value:%d\n", key, value)
	}
}

func main() {
	testMapExist()
}
