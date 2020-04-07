package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	// 声名一个变量用来存储需要读取的 gzip 文件名
	fileName := "test.txt.gz"

	// 使用 os.Open 方法打开文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("rerad file failed, error:", err)
		return
	}
	defer file.Close()
	// 调用 gzip.NewReader 方法创建一个对象，该方法返回一个读取到的文件对象，和一个错误信息
	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("read gzip file failed , error:", err)
		return
	}

	// 调用方式一：不使用 bufio 缓存的情况下直接通过 gzip 方法读取 gz 压缩包返回的文件对象对文件进行读取，
	// 注：若读取的文件中有中文存在时，由于一个占用三个UTF8字节的关系，所以在读取时可能会存在读不完整的情况；
	var content []byte
	var buf [128]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file:", err)
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))

	// 声名一个 *bufio.Reader 结构体指针类型的变量
	var r *bufio.Reader
	// 调用方式二：调用 bufio 包对 gzip 返回的文件对象进行再次封装，为读取文件创建缓存提高性能
	r = bufio.NewReader(reader)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Done not file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
