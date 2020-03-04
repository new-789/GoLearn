package main

import (
	"fmt"
	// 引入自定义的静态库文件(包)
	"github.com/luffycity/listen11/custom"
)

func main() {
	// 调用自定义包中的 Add 函数并进行传参
	var sum int = custom.Add(10, 70)
	fmt.Printf("sum:%d\n", sum)
}
