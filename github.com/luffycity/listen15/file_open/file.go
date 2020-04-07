package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 以只读方式打开文件方法
	inputFile, err := os.Open("./file.go")
	if err != nil {
		fmt.Println("open file:", err)
	}
	// 程序执行完后，使用打开的文件对象关闭文件方法
	defer inputFile.Close()

	// 声名一个 byte 类型的数组变量,用来确定每次读取文件内容的大小
	var buf [512]byte
	// 声名一个变量用来保存获取到的文件内容
	var messages []byte
	// 通过 for 循环读取整个文件的内容
	for {
		// 读取文件内容,并将 buf 更改为切片类型的数据传给 Read 方法
		n, err := inputFile.Read(buf[:])
		// 整个文件读取完毕后返回一个 io.EOF 错误，
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file:", err)
			return
		}
		// 将一个切片追加到另一个切片中方法, buf[:n]... 表示展开该切片方法，否则无法追加
		messages = append(messages, buf[:n]...)
	}
	fmt.Println(string(messages))

	// 表示读取 buf[:] 整个数组指定的内容，并且从 256 字节开始读取
	//n, err := inputFile.ReadAt(buf[:], 256)
	//if err != nil{
	//	return
	//}
	//fmt.Println(string(buf[:n]))
}
