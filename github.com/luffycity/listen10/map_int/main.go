package main

import (
	"fmt"
)

func main() {
	// 声名一个 map 数据类型的变量，并对其进行初始化操作,
	// 也可以直接使用 a := map[string]int{key:value} 语法
	var a map[string]int = map[string]int{
		"stu01": 100,
		"stu02": 200,
		"stu03": 300,
	}
	// 修改 map 数据中 key 为 sut01 对应的 value
	a["stu01"] = 888
	// 往 map 类型的变量 a 中插入一条数据
	a["stu04"] = 400

	var key string = "stu04" // 定义一个变量并初始化内容为 map 数据中存在的 key 值
	// 通过定义的变量 key 访问 map 类型总的数据
	fmt.Printf("the's key[%s] value %d\n", key, a[key])
	// 直接通过 map 变量[key] 语法访问 map 数据中的元素
	fmt.Printf("the's key[stu01] value %d\n", a["stu01"])
}
