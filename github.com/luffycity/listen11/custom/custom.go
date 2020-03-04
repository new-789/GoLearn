package custom

import "fmt"

var (
	Sum int
)

func Add(a int, b int) int {
	Sum = a + b
	return Sum
}

func init() {
	fmt.Println("这是 custom 包中的 init 函数")
}
