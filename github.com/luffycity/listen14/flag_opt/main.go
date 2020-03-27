package main

import (
	"flag"
	"fmt"
)

// 声名全局变量，用来存储输入的参数
var (
	recusive bool
	test     string
	level    int
)

// 声名一个 init 函数，该函数是在入口 main 函数运行前自动执行并将函数中的内容假造到内存
func init() {
	// 定义参数类容
	flag.BoolVar(&recusive, "r", false, "recusive xxx")
	flag.StringVar(&test, "s", "default string", "string option")
	flag.IntVar(&level, "l", 1, "level option")
	// 使定义的参数生效
	flag.Parse()
}

func main() {
	fmt.Printf("recusive:%v\ntest:%s\nlevel:%d\n", recusive, test, level)
}
