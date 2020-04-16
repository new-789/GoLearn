package main

import (
	"fmt"
)

// 声名一个函数，并接收一个接口类型的参数
func assert1(a interface{}) {
	i, ok := a.(int) // 用来断言传入的值是否为 int 类型，是则将正确类型的值存储在 i 变量中
	if ok {
		fmt.Println(i)
		return
	}

	s, ok := a.(string)
	if ok {
		fmt.Println(s)
		return
	}
	// 对没有断言类型的数据则抛出提示信息
	fmt.Println("不知道是什么类型的数据")
}

func testInterface1() {
	var a int = 100
	assert1(a)

	var b string = "hello"
	assert1(b)

	var c float32 = 5.233
	assert1(c)
}

func assert2(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Printf("a is string, value:%s\n", a.(string))
	case int:
		fmt.Printf("a is int, value:%d\n", a.(int))
	case map[string]int:
		fmt.Printf("a is map, value:%v\n", a.(map[string]int))
	default:
		fmt.Println("sorry not support type")
	}
}

func assert3(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Printf("a is %T, value:%d\n", v, v)
	case string:
		fmt.Printf("a is %T, value:%s\n", v, v)
	case map[string]int:
		fmt.Printf("a is %T, value:%v\n", v, v)
	default:
		fmt.Println("sorry not support type")
	}
}

func testInterface3() {
	var a int = 1000
	assert3(a)
	var b string = "hello"
	assert3(b)
	var c map[string]int = make(map[string]int, 100)
	c["test1"] = 200
	c["test2"] = 300
	assert3(c)
}

func testInterface2() {
	var a int = 1000
	assert2(a)
	var b string = "hello"
	assert2(b)
	var c map[string]int = make(map[string]int, 100)
	c["test1"] = 200
	c["test2"] = 300
	assert2(c)
}

func main() {
	//testInterface1()
	//testInterface2()
	testInterface3()
}
