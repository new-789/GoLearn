package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testForMap() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	// 定义一个 map 数据类型结构，并初始化
	var a map[string]int = make(map[string]int, 1024)
	// 通过循环往 map 类型中插入数据
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("%d", i)
		value := rand.Intn(1000)
		a[key] = value
	}

	// 遍历 map 并打印输出相应的数据
	for key, value := range a {
		fmt.Printf("map a[%s]=%d\n", key, value)
	}
}

func main() {
	testForMap()
}
