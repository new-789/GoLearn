package main // 用来表示该程序是一个可执行程序包的名字一定是 main ，Go 语言规定的语法

// 导入包语法
import (
	"fmt"  // 1. fmt 包：输入输出标准包
	"time" // 2. time 包：时间处理模块包
)

func main() { // 程序执行时的入口函数
	fmt.Println("hello world")   // 打印输出 hello world 内容到终端
	time.Sleep(time.Second * 10) //  执行完之后停止 10 秒钟之后结束程序
}
