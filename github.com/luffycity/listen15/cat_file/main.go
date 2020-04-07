package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		// 使用 bufio.ReadBytes 方法读取传入的文件内容，并使用 \n 回车作为分割符
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
		//return
	}
}

func main() {
	// 通过 flag 包获取用户输入的命令中需要打开的文件名
	flag.Parse()
	// 如果用户输入的内容中不带有文件名，则返回到命令行界面
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	// 根据 flag.NArg 中的文件数量，循环打开文件
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
