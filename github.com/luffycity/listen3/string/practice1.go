// 写一个程序，对英文字符串进行逆序
package main

import "fmt"

func testReverStringv1(str string) (changeStr string) {
	// var str1 = str
	fmt.Printf("len_str = %d\n", len(str))
	var bytes []byte = []byte(str)

	// len(str)/2 表示只需要循环整个字符串长度的一半即可
	for i := 0; i < len(str)/2; i++ {
		tmp := bytes[len(str)-i-1]     // 获取字符串中后面的字符存到变量 tmp 中
		bytes[len(str)-i-1] = bytes[i] // 从后往前逐个替换为前面的字符
		bytes[i] = tmp                 // 逐个从前往后替换字符
	}
	changeStr = string(bytes)
	return changeStr
}

func testReverStringv2(str string) (changeStr string) {
	fmt.Printf("len_rune=%d\n", len([]rune(str)))
	var RUNE []rune = []rune(str)

	for i := 0; i < len(RUNE)/2; i++ {
		tmp := RUNE[len(RUNE)-i-1]
		RUNE[len(RUNE)-i-1] = RUNE[i]
		RUNE[i] = tmp
	}

	changeStr = string(RUNE)
	return changeStr
}

func testHuiWen(str1 string) (str2, Huiwen string) {
	fmt.Printf("原文:%s\n", str1)
	var runeStr []rune = []rune(str1)

	for i := 0; i < len(runeStr)/2; i++ {
		tmp := runeStr[len(runeStr)-i-1]
		runeStr[len(runeStr)-i-1] = runeStr[i]
		runeStr[i] = tmp
	}
	var HuiWen = string(runeStr)
	if str1 == HuiWen {
		str2 = "这个字符串是回文\n"
	} else {
		str2 = "这个字符串不是回文\n"
	}
	return str2, HuiWen
}

func main() {
	// 字符串逆序
	var changeStr = testReverStringv1("helloworld")
	fmt.Printf("change str = %s\n", changeStr)
	// 中文逆序
	var changeStr1 = testReverStringv2("中国人民万岁")
	fmt.Printf("change rune=%s\n", changeStr1)
	// 回文示例
	var str, HuiWen = testHuiWen("上海自来水来自海上")
	fmt.Printf("回文:%s\n", HuiWen)
	fmt.Printf("是否是回文：%s\n", str)
}
