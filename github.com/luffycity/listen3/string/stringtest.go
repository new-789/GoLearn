package main

import (
	"fmt"
)

func testString() {
	var str = "hello"                                       // 定义一个字符串变量
	fmt.Printf("str[0]=%c len(str)=%d\n", str[0], len(str)) // 获取定义字符串中的第一个字符

	// 遍历字符串操作
	for index, val := range str {
		fmt.Printf("str[%d]=%c\n", index, val)
	}

	// 修改字符换示例
	var byteSlice []byte                  // 定义一个切片类型的变量
	byteSlice = []byte(str)               // 将 str 类型数据转换为 byte 类型数据语法
	byteSlice[0], byteSlice[4] = 'o', 'h' // 修改 byte 类型中的数据,注：修改 byte 类型的字符串数据时只能用单引号
	str = string(byteSlice)               // 将 byte 类型的数据转换为 string 类型
	fmt.Printf("change str=%s\n", str)

	// 查看字符串的长度
	str1 := "hello world"
	fmt.Printf("len(str) = %d\n", len(str1))
	str2 := "中国"
	fmt.Printf("len(str2)=%d\n", len(str2))

	// rune 类型示例
	var b rune = '中' // 定义 rune(utf8) 字符,即指明该字符份编码类型为 utf8,只允许使用单引号('')给 rune 类型进行赋值
	fmt.Printf("b=%c\n", b)

	// 将 string 字符串转换为 rune 类型字符
	var runeSlice []rune     // 定义 rune 类型的切片
	runeSlice = []rune(str2) // 将字符串强制转换为 rune 切片类型的字符
	fmt.Printf("str1[]rune_length=%d\n", len(runeSlice))
}

// func main() {
// 	testString()
// }
