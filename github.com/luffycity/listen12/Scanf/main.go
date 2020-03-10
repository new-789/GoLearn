package main

import "fmt"

func main() {
	var set string
	fmt.Println("请出入内容:")
	val, err := fmt.Scanf("%s\n", &set)
	if err != nil {
		fmt.Printf("输入出现错误")
	}
	fmt.Println(val)
	fmt.Printf("set:%s\n", set)
}
