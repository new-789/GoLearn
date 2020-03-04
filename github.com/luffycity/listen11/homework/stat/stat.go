package main

import (
	"fmt"
	"strings"
)

func statWordCount(str string) map[string]int {
	// 声名一个 map 类型的数据，并进行初始化
	var result map[string]int = make(map[string]int, 128)
	// 对字符串进行按照 "空格" 切割，返回一个切片类型的数据
	words := strings.Split(str, " ")
	// 循环切片内容，获取切片中的每一个数据
	for _, val := range words {
		// 根据 key 获取对应的值与 key 存在与否是状态码
		count, state := result[val]
		// 根据获取的状态码判断当前 key 是否存在
		if !state {
			// 不存在则往 map 类型中添加一条数据
			result[val] = 1
		} else {
			// 存在则更新该 key 的值
			result[val] = count + 1
		}
	}
	return result
}

func word() {
	var str string = "how do you do ? do you like mi ?"
	mapTypeResult := statWordCount(str)
	fmt.Printf("result:%#v\n", mapTypeResult)
}

func main() {
	word()
}
