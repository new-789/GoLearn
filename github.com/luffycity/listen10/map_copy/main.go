package main

import (
	"fmt"
)

func modify(a map[string]int, key string) {
	// 声名 modify 函数，并接受一个 map 类型的数据和字符串类型的数据作为参数
	a[key] = 1000 // 通过参数名修改传入的 map 数据中 key 为传入 key 字符串的元素
}

func main() {
	var b map[string]int = map[string]int{
		"stu01": 100,
		"stu02": 200,
		"stu03": 300,
	}
	fmt.Printf("初始 map 数据:%#v\n", b)
	// 将 map 数据 b 赋值给 c，由于 map 数据存储的数据指向的都是内存地址，所以赋值操作是引用类型的拷贝
	c := b
	// 声名一个变量，用来存储粗腰操作 map 数据中的 key
	var key string = "stu02"
	c[key] = 800
	fmt.Printf("通过 c 引用 b map 数据修改内容后:%#v\n", b)
	modify(c, key)
	fmt.Printf("通过 modify 函数修改 b map 数据内容后:%#v\n", b)
}
