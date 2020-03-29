package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Test"
	app.Usage = "使用方法"
	app.Action = func(c *cli.Context) error {
		var cmd string
		// cli 框架会将用户传入的参数保存在 c.NArg() 方法中，该方法返回一个切片长度类型的数据，可用该函数来判断用户是否传入了参数
		// 此处用来判断保存用户参数切片中的数据如果大于零则获取用户传入的参数
		if c.NArg() > 0 {
			// 通过 c.Args().Get(index) 方法可获取到用户传入的多个参数
			cmd = c.Args().Get(0)
		}
		fmt.Println("hello friend cmd:", cmd)
		return nil
	}
	app.Run(os.Args)
}
