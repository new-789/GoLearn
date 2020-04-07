package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 使用 ioutil 工具包中的 ReadFile 方法读取文件内容方法，该方法返回两个值，
	// 第一个值为读取到的文件内容 **注意是二进制格式的， 第二个值为出错后的错误信息
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read file failed, error:", err)
		return
	}

	fmt.Println(string(content))
}
