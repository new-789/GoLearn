package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("./buf.go")
	if err != nil {
		fmt.Println("erad file:", err)
		return
	}
	defer inputFile.Close()

	// 使用 bufio.NewReader 方法构造一个对象，接收的参数为 os.Open 返回的文件对象
	inputReader := bufio.NewReader(inputFile)
	for {
		// 通过 bufio 创建的对象调用 ReadString 方法读取字符串形式的文件内容，该方法每次读取一行内容
		// 接收的参数为指定一行结束的符号，返回的内容直接为字符串类型的数据，和一个错误信息
		line, err := inputReader.ReadString('\n')
		// 当读取完毕后同样返回文件对象的 io.EOF 错误
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed:", err)
			return
		}
		fmt.Println(line)
	}
}
