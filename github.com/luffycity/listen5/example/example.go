package main

import (
	"fmt"
)

// 质数判断逻辑函数
func primelogic(n int) bool {
	if n <= 1 { // 如果N这个数小于等于1，那么它肯定不是质数，直接返回 false
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 { // 如果n这个数能够被其它数整除则说明该数不是质数，返回 false
			return false
		}
	}
	return true // 循环结束
}

func primenum() {
	for i := 2; i < 100; i++ {
		if primelogic(i) == true {
			fmt.Printf("[%d] is prime\n", i)
		}
	}
}

func primelogic1(num int) bool {
	first := num % 10         // 取个位数上的数
	second := (num / 10) % 10 // 取十位上的数
	third := (num / 100) % 10 // 取百位上的数
	sum := first*first*first + second*second*second + third*third*third
	if sum == num {
		return true
	}
	return false
}

func example() {
	for i := 100; i < 1000; i++ {
		if primelogic1(i) == true {
			fmt.Printf("[%d] 是水仙花数\n", i)
		}
	}
}

func strCount(str string) (charcount int, numcount int, spacecount int, othercount int) {
	// 将字符串转换为 utf8 字符，默认字符按照 ascall 码算字节不同的内容字节数不同影响实际结果，转换后一个字符占一位，返回一个数组
	strChars := []rune(str)
	for i := 0; i < len(strChars); i++ {
		if strChars[i] >= 'a' && strChars[i] <= 'z' || strChars[i] >= 'A' && strChars[i] >= 'Z' { // 判断是否为大小或小写字符串
			charcount++
			continue
		}
		if strChars[i] >= '0' && strChars[i] <= '9' { // 判断是否为数字
			numcount++
			continue
		}
		if strChars[i] == ' ' { // 判断是否为空字符串
			spacecount++
			continue
		}
		othercount++
	}
	return
}

func example1() {
	var str string = "AddAADSkjhjkujhj      我是 22632434 4555552 中国人"
	charcount, numcount, spacecount, othercount := strCount(str)
	fmt.Printf("char count:[%d], num count:[%d], space count:[%d], otrher count:[%d]\n", charcount, numcount, spacecount, othercount)
}

func main() {
	// primenum()
	// stat()
	example1()
}
