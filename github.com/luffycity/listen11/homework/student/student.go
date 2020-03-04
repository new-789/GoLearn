package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//func testInterface() {
//	// 声名一个 interface 类型的变量，注 interface 类型可以用来存储任何类型的数据
//	var a interface{}
//	var b int = 100
//	var c float64 = 123.123
//	var d string = "test"
//	a = b
//	fmt.Printf("a:%v\n", a)
//	a = c
//	fmt.Printf("a:%v\n", a)
//	a = d
//	fmt.Printf("a:%v\n", a)
//}

// 存储学生信息的数据结构map:{id:{id: 1, name: stu01, ...}, ...}
func studentStore() {
	var stuMap map[int]map[string]interface{} = make(map[int]map[string]interface{}, 10)
	// 声名学生信息变量
	var id int = 1
	var name string = "stu01"
	var score float64 = 98.9
	var age int = 33
	//
	value, state := stuMap[id]
	if !state {
		// 初始化第二层的 map 类型
		value = make(map[string]interface{}, 4)
	}
	// 更新学生信息
	value["id"] = id
	value["name"] = name
	value["age"] = age
	value["score"] = score

	stuMap[id] = value

	// 使用 for 循环批量添加学生信息
	for i := 0; i < 5; i++ {
		val, ok := stuMap[i]
		// 判断如果 key 存在则进入下一次循环，不存在则初始化对应的 map 类型的 value 值
		if ok {
			continue
		} else {
			// 初始化 map 类型的 value 值
			val = make(map[string]interface{}, 4)
		}
		// 更新学生信息
		val["id"] = i
		val["name"] = fmt.Sprintf("stu%d", i)
		val["age"] = rand.Intn(100)
		val["score"] = rand.Float64() * 100
		stuMap[i] = val
	}

	// 对 map 中的数据按照 key 进行排序
	var keys []int = make([]int, 0, 10)
	for k, _ := range stuMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 循环打印输出
	for _, key := range keys {
		fmt.Printf("id=%d stu info : %v\n", key, stuMap[key])
	}
}

func main() {
	//testInterface()
	studentStore()
}
