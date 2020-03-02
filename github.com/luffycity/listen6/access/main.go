package main

import (
	"fmt"

	"github.com/luffycity/listen6/calc"
)

func main() {
	var s1 int = 100
	var s2 int = 300
	sum := calc.Add(s2, s1) // 调用 calc 表中的 Add 函数
	fmt.Printf("s1+s2= %d\n", sum)
	// sub := calc.sub(s2, s1) // 在 calc 包外调用 sub 函数，提示错误
	// fmt.Printf("s2-s1= %d\n", sub)

	fmt.Printf("A=%d\n", calc.A) // 调用 calc 包中的 A 变量
	// fmt.Printf("a=%d\n", calc.a) // 调用 calc 包中的 a 变量，提示错误
	a := calc.Test() // 调用 calc 包中的 Test 函数并获取 calc 包中的变量 a
	fmt.Printf("a=%d\n", a)
}
