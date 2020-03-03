package main

import (
	"fmt"
)

func main() {
	var a map[string]int
	fmt.Printf("未初始化 a:%v\n", a)
	if a == nil {
		// 使用 make 方法初始化声名的 map 类型变量 a ，参数一同样是 map 类型，参数二则为 map 容量大小
		a = make(map[string]int, 16)
		a["stu01"] = 100 // 初始化之后往 map 类型的数据中插入数据记录
		a["stu02"] = 200 // 初始化之后往 map 类型的数据中插入数据记录
		a["stu03"] = 300 // 初始化之后往 map 类型的数据中插入数据记录
		fmt.Printf("初始化插入数据之后 a:%#v\n", a)
	}
}
