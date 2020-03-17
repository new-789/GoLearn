package main

import (
	"fmt"
)

// 使用 int 包定义一个 integre 类型,注使用 int 包定义 integre 类型并不是说 integre 类型就等同于 int 类型
type Integre int

// 定义一个 Print 方法，将 Integre 设置为接受者，表名该方法属于 Integre
func (i Integre) Print() {
	fmt.Println(i)
}

func main() {
	// 声名一个变量为自定义的 integre 类型
	var integre Integre
	integre = 1000
	integre.Print() // 使用声名的变量 integre 调用 Print 方法

	var b int = 9888
	// 使用我们定义 Integre 类型强制将 int 类型的变量 b 转换为 Integre 类型
	integre = Integre(b)
	integre.Print()
}
