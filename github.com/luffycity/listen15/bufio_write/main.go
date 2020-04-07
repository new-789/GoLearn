package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fileName = "bufioWrite.txt"
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}

	// 通过 bufio 调用 NewWriter 方法构造一个对象
	writeObj := bufio.NewWriter(fileObj)
	// 声名一个变量，用来记录写入内容的字节大小
	var num int
	for i := 0; i < 10; i++ {
		var str = "使用 bufio 包写数据到文件测试\n"
		// 调用 WriteString 方法将内容写入到文件, 该方法同样返回两个值，一个为写入文件内容的长度，
		// 一个为写入文件出错后的错误提示信息
		n, er := writeObj.WriteString(str)
		if er != nil {
			fmt.Printf("the [%d] wirte  info to file failes, err:%s\n", i, er)
			return
		}
		num += n
	}
	fmt.Printf("本次共写入了 【%d】 字节的内容", num)
	e := writeObj.Flush()
	if e != nil {
		fmt.Println("write file failed, err:", e)
		return
	}
}
