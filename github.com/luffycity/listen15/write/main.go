package main

import (
	"fmt"
	"os"
)

func main() {
	var fileName = "test.txt"
	// 打开一个文件，如果文件不存在则进行创建操作，如果存在则清空改文件然后在写入，
	// 指定权限为 0666 表示所有用户均有读写操作的权限
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed. err:", err)
		return
	}
	defer fileObj.Close()

	// 声名一个变量，用来定义需要写入文件的数据
	str := `
		func main() {
			fmt.Println("这是 Write 方法写入的内容")
		}` + "\n\n"
	// 通过 Write 方法将内容写入文件
	WriteN, err := fileObj.Write([]byte(str))
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
	fmt.Println("wirte file success, write is len:", WriteN)

	// 通过 WriteAt 方法将内容写入文件
	var s = "这是 WriteAt 方法写入的内容\n\n"
	WriteAtn, e := fileObj.WriteAt([]byte(s), int64(WriteN))
	if e != nil {
		fmt.Println("write file failed, err:", e)
		return
	}
	fmt.Println("wirte file success, write is len:", WriteAtn)

	// 通过 WriteString 方法将内容写入文件
	//var st = "测试 WriteString 方法写入数据"
	//WriteStr, er := fileObj.WriteString(st)
	//if e != nil{
	//	fmt.Println("write file failed, err:", er)
	//	return
	//}
	//fmt.Println("wirte file success, write is len:", WriteStr)
}
