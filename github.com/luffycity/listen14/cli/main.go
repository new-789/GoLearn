package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	// 声名两个变量用来接收用户的参数
	var language string
	var recusive bool
	app := cli.NewApp()

	// app.Flags 为切片类型的数据，该切片中的元素为结构体类型的数据，我们可以在这个结构体中定制用户传入的选项
	app.Flags = []cli.Flag{
		// 在 StringFlag 结构体中用来定义字符串类型的选项, 注：在 cli 包中有函数接收指针类型的 StringFlag 为参数
		// 所以此处必须为内存地址
		&cli.StringFlag{
			Name:        "lang, l", // 用来定义选项名称，使用都改分割起来的为选项的简称(该示例不支持简称短选项不知道为啥)
			Value:       "english", // 该字段用来设置默认值
			Destination: &language, // 该字段用来绑定一个地址类型的变量
		},

		// 在 BoolFlag 结构体中用来定义布尔类型的选项，注：在 cli 包中有函数接收指针类型的 BoolFlag 为参数
		// 所以此处必须为内存地址
		&cli.BoolFlag{
			Name:        "recusive, r", // 该字段用来定义选项名称，使用都改分割起来的为选项的简称(该示例不支持简称短选项不知道为啥)
			Usage:       "这是使用说明内容",
			Value:       false,     // 该字段用来设置默认值为 false
			Destination: &recusive, // 该字段用来绑定一个地址类型的变量
		},
	}
	app.Action = func(context *cli.Context) error {
		var cmd string
		if context.NArg() > 0 {
			cmd = context.Args().Get(0)
			fmt.Printf("获取到的选项参数为:%s\n", cmd)
		}
		fmt.Println("recusive is", recusive)
		fmt.Println("language is", language)
		return nil
	}
	app.Run(os.Args)
}
