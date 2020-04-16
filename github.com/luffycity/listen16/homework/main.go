package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ListDir(dirPath string, deep int) (err error) {
	// 通过 ioutil 方法中的 ReadDir 方法获取整个目录下的所有文件以及文件夹，返回一个切片数据和一个错误信息
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	// 判断 deep 是否等于 1，等于 1 则表示当前在第一级目录，则输出 |----目录名
	if deep == 1 {
		// 通过 filepath 标准包中的 Base 方法获取当前目录中的所有文件？
		fmt.Printf("|____%s\n", filepath.Base(dirPath))
	}
	// 获取路径的分隔符，windows 下的目录分隔符为：\，Linux 系统下的目录分隔符为： / ；
	sep := string(os.PathSeparator)
	// 循环遍历整个目录结构,遍历整个子目录
	for _, filename := range dir {
		// 判断当前目录是否为下一级目录
		if filename.IsDir() {
			fmt.Printf("|")
			// 循环遍历的目录深度
			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			// 打印目录深层目录中的目录名或文件名
			fmt.Printf("____%s\n", filename.Name())
			// 循环自调用，不断的进入下一层目录，并增加层级计数变量
			ListDir(dirPath+sep+filename.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		// 循环遍历第二层目录中的目录名或文件名，并输出 | 符号
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("____%s\n", filename.Name())
	}
	return nil
}

func main() {
	// 处理相关命令方法
	app := cli.NewApp()
	app.Name = "tree"
	app.Usage = "list all files"

	app.Action = func(context *cli.Context) error {
		// 什么一个字符串乐行的变量，并设定默认目录为当前目录
		var dir string = "."
		if context.NArg() > 0 {
			dir = context.Args().Get(0)
		}
		// 调用 ListDir 方法，并将目录名及目录的深度传递给该函数
		ListDir(dir, 1)
		return nil
	}
	// 执行 app 方法
	app.Run(os.Args)
}
