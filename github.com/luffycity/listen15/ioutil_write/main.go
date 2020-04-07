package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var opFileName = "openfile.txt"
	var wrFileName = "writefile.txt"
	// 通过 ioutil 工具包中的 ReadFile 方法读取整个文件内容
	readObj, err := ioutil.ReadFile(opFileName)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}

	// 调用 ioutil 工具包中的 WriteFile 方法将内容整个保存在文件中, readObj = []byte("需要写的内容")
	er := ioutil.WriteFile(wrFileName, readObj, 0755)
	if er != nil {
		fmt.Println("write file failed, err:", er)
		return
	}
}
