package main

import (
	"fmt"
	"os"
)

func main() {
	// 打印 os.Args 中的第零个元素，即该程序的名称
	fmt.Println("args[0]=", os.Args[0])
	// 如果 os.Args 切片的长度大于 1 则表示该用户传入了参数，
	if len(os.Args) > 1 {
		for index, val := range os.Args {
			if index == 0 {
				continue
			}
			fmt.Printf("用户传入的参数args[%d]：%s\n", index, val)
		}
	}
}
