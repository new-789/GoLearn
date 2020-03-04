package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 按照时间建立一个随机数种子
	rand.Seed(time.Now().UnixNano())
	// 创建 map 类型数据
	var a map[string]int = make(map[string]int, 10)
	// 通过 for 循环往 map 中插入数据
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("stu%d", i)
		// value 的值在 1000 之内随机生成
		value := rand.Intn(1000)
		// map 插入值
		a[key] = value
	}
	// 定义一个切片用来存储 map 数据中的key
	var keys []string = make([]string, 0, 5)
	// 遍历 map 数据，并将 key 添加存储在创建的切片 keys 中
	for key, _ := range a {
		//fmt.Printf("map a[%s]=[%d]\n", key, value)
		keys = append(keys, key)
	}
	// 对存储所有 key 的切片进行排序
	sort.Strings(keys)
	// 循环遍历切片 keys 然后按照已排好序的 key 获取 map 中的数据输出即为排好序的输出结果
	for _, key := range keys {
		fmt.Printf("map a[%s]=[%d]\n", key, a[key])
	}
}
