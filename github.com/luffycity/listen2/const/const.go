package main

import "fmt"

// func main() {
// 	// 定义常量语法
// 	const a string = "hello world" // 定义字符串常量
// 	const b = "hello word"         // 定义字符串常量
// 	const Pi = 3.1415926           // 定义浮点类型常量
// 	const d = 9 / 3                // 定义数值类型常量

// 	fmt.Printf("a=%s b=%s c=%f d=%d\n", a, b, Pi, d)
// 	// 修改常量会报错
// 	// a = "test"
// }

// 优雅的写法
// func main() {
// 	// const (
// 	// 	a int    = 100
// 	// 	b string = "hello world"
// 	// )

// 	// 专业的常量写法,此种语法，第一个常量定义type 与 value后，若第二个常量没有定义则默认与它上一个常量类型及值一致
// 	const (
// 		a int    = 100
// 		b        // 默认与上一个 a 常量的类型及值一致
// 		c string = "hello world"
// 		d        // 默认与上一个 c 常量的类型及值一致
// 	)
// 	fmt.Printf("a=%d b=%d c = %s d=%s", a, b, c, d)
// }

// 专业写法
func main() {
	// 示例一
	const (
		a = iota // iota 语义即表示 a 常量的默认值为 0
		b        // 由于上一个常量存在 iota 语法，固 b 默认值为 1
		c        // 由于上面常量存在 iota 语法，固 c 默认值为 2
	)
	// 示例二
	// << iota 语义表示在当前值的基础上左移一位,若给一个常量指定了值则取代 iota 默认为 0 的值
	const (
		e = 2 << iota //
		// 常量 f 没有定义，默认等于 e = 2 << iota, 由于 iota 每隔一行递增一位且左移的位置同样递增，
		// 所以 f 的默认值在e的基础上左移两位为4
		f
		// 常量 g 没有定义，默认等于 e = 2 << iota, 由于 iota 每隔一行递增一位且左移的位置同样递增，所以 g 的默认值在 e 的基础上为第二次递增所以为6，
		// 而 << iota 表示左移一位g的位置较 e 的位置由于需要左移两位，固最后 g 的结果的值是 8
		g
	)

	fmt.Printf("a=%d b=%d c=%d\n", a, b, c)
	fmt.Printf("e=%d f=%d g=%d\n", e, f, g)
}
