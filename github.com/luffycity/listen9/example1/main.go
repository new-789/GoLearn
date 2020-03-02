package main

import (
	"fmt"
)

func main() {
	// make([]string, 5, 10) 表示创建的字符串切片包含5个空元素，10 则表示它所引用的字符串数组长度是等于10
	var sa []string = make([]string, 5, 10)
	fmt.Println(sa)
	for i := 0; i < 10; i++ {
		// fmt.Sprintf 格式化输出整数类型数据并将其转换为字符串类型的数据
		sa = append(sa, fmt.Sprintf("%d", i))
	}
	fmt.Println(sa)
}
