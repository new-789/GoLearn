package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (wirtten int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open read file failed, err:", err)
		return
	}

	dst, er := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0666)
	if er != nil {
		fmt.Println("open write file failed, err:", er)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	var opFimeName = "opfilename.txt"
	var dstFimeName = "dstfilename.txt"
	wirtten, err := CopyFile(dstFimeName, opFimeName)
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Printf("本次共拷贝了【%d】内容\n", wirtten)
}
