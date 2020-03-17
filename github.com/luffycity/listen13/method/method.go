package main

import (
	"fmt"
)

type People struct {
	Name    string
	Country string
}

// 定义一个 Test 方法，前面加上接受者 People 结构体
func (p People) Test() {
	// 通过 p 来访问 People 类型中的字段内容
	fmt.Printf("Name : %s , Country : %s\n", p.Name, p.Country)
}

// 定义一个 set 方法，前面加上接受者 People 结构体，并且该方法同时接收两个参数
func (p People) set(name string, country string) {
	p.Name = name // 通过传入的参数修改 People 结构体重的字段内容
	p.Country = country
}

func (p *People) Set(name string, country string) {
	p.Name = name
	p.Country = country
}

func main() {
	// 通过结构体进行实例化一个 p1 类型为 People 的对象
	var p1 People = People{
		Name:    "冠状病毒",
		Country: "武汉",
	}
	// 通过 p1 对象调用 Test 方法，在调用时会将 p1 传递给 Test 方法中前面括号中的 p ，即 p 就是 p1 的一个拷贝
	p1.Test()
	//p1.set("新冠状病毒", "广州")
	//p1.Test()

	// 首先通过 &p1 获取类型结构体的内存地址，然后调用 Set 方法，编辑器会自动帮助我们将 p1 的内存地址传递给 Set 方法前面的接受者
	(&p1).Set("新冠状病毒", "好了")
	p1.Set("新冠状病毒", "好")
	p1.Test()
}
