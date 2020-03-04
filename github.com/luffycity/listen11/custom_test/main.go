package main

import (
	"fmt"
	// 引入自定义的静态库文件(包),前面加上下划线表示我们没有导入该包中的任何变量或是函数，
	// 但是我们有需要该包中的 init 函数来为我们做初始化，此时我们就可以在前面加上 _ 下划线
	_ "github.com/luffycity/listen11/custom"
)

func init() {
	fmt.Println("程序启动时自定调用该 init 函数")
}

func init() {
	fmt.Println("程序启动时自定调用该 init2 函数")
}

func main() {
	// 调用自定义包中的 Add 函数并进行传参
	//var sum int = custom.Add(10, 70)
	fmt.Printf("sum:%d\n", sum)
}
