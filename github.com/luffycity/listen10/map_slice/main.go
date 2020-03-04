package main

import (
	"fmt"
)

// 数据类型 []map[string]int = [{key:value, key:value} {key:value, key:value}]
func sliceMap() {
	// 声名一个 map 类型的切片，并进行初始化长度为 5，容量为 10
	var slicemap []map[string]int = make([]map[string]int, 5, 10)
	for index, val := range slicemap {
		fmt.Printf("slice[%d]=%v\n", index, val)
	}
	fmt.Println()
	// 对切片中的第 0 个 map 类型的元素进行初始化容量为 5
	slicemap[0] = make(map[string]int, 5)
	// 往切片中的第 0 个 map 类型的元素中插入数据
	slicemap[0]["stuo1"] = 1000
	slicemap[0]["stuo2"] = 2000
	slicemap[0]["stuo3"] = 3000
	for index, val := range slicemap {
		fmt.Printf("slice[%d]=%#v\n", index, val)
	}
}

// 数据类型：map[strint][]int = {key:[], key:[]}
func mapSlice() {
	// 声名一个 map 类型 key 为 string 值为切片类型的数据并进行初始化容量为 10
	var ms map[string][]int = make(map[string][]int, 10)
	// 声名 key 变量用来保存需要判断的 key
	key := "stu01"
	// 判断 key 是否存在返回一个 key 的状态(ok表示)，和当前key对应的切片值(value接收的值)
	value, ok := ms[key]
	// 判断 key 返回的状态如果不为 true
	if !ok {
		// 往 map 中增加一个 key 为 stu01 值，并初始化该 key 所对应类型为切片的value值
		ms[key] = make([]int, 0, 10)
		// 将初始化好的空切片赋值给 value
		//value = ms[key]
	}
	// 给 map 中 value 为切片类型中添加数据
	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	// 往 map 中为 stu01 的 key 添加一条切片数据的值
	ms[key] = value
	fmt.Printf("mapslice:%v\n", ms)
}

func main() {
	//sliceMap()
	mapSlice()
}
