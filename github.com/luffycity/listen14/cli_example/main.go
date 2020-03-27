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
		fmt.Println("hello cli test")
		return nil
	}
	app.Run(os.Args)
}
