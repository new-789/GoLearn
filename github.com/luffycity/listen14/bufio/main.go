package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	// 1. 声明一个 buio.Reader 类型的变量
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	// 2.使用 bufio 中的 NewReader 方法创建一个对象，接收一个终端文件输入方法为参数
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	// 3.使用创建的对象调用 ReadString 方法用来指定一个分隔符，返回获取到的数据及错误信息
	input, err = inputReader.ReadString('\n') // 注，此处为单引号
	// 判断 err 如果为空则打印获取到的内容
	if err == nil {
		fmt.Printf("The input was:%s\n", input)
	}
}
