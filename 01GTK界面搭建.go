package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)  // 初始化

	// 通过 GTK 创建界面
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	// 设置窗体大小
	win.SetSizeRequest(480, 320)
	// 设置窗体标题
	win.SetTitle("GTK 窗口")
	// 将界面显示在屏幕中
	win.Show()
	// 调用 gtk 下的 main 主函数运行程序
	gtk.Main()
}
