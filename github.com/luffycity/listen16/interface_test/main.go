package main

import "fmt"

// 定义一个接口，接口的名字为动物 Animal ，并在接口中定义三个方法，即表明要实现这三种方法才能算得上一个完整的动物
type Animal interface {
	Talk()        // 定义一个叫的方法
	Eat()         // 定义一个吃东西的方法
	Name() string // 定义一个 Name 方法，并返回一个字符串
}

// 定义一个 Dog 结构体
type Dog struct {
}

type Pig struct {
}

// 定义一个函数，使它属于 Dog 结构体，实现接口中的 Talk 方法
func (d *Dog) Talk() {
	fmt.Println("狗在叫：汪.汪.汪,是在告诉我们说它饿了要吃东西")
}

// 定义一个函数，使它属于 Dog 结构体，实现接口中的 Eat 方法
func (d *Dog) Eat() {
	fmt.Println("有人丢了一根骨头给狗吃")
}

// 定义一个函数，使它属于 Dog 结构体，实现接口中的 Name 方法
func (d *Dog) Name() string {
	fmt.Println("路上有一只狗，它的名字叫: 旺财")
	return "旺财"
}

//func testinterface1(){
//	var d *Dog = &Dog{}  // 通过结构体 Dog 声明一个实例 d
//	var a Animal  // 声名一个变量 a ，类型为定义的方法 Animal
//	a = d  // 将结构体型实例 d，赋值给方法类型的变量 a，使其可以通过结构体定义的方法可以调用接口中的方法
//	a.Name()  // 调用接口中的 Name 方法
//	a.Talk()  // 调用接口中的 Talk 方法
//	a.Eat()  // 调用接口中的 Eat 方法
//}

// 定义一个函数，使它属于 Dog 结构体，实现接口中的 Talk 方法
func (p Pig) Talk() {
	fmt.Println("猪在叫：嗯.嗯.嗯,是在告诉我们说它饿了要吃东西")
}

// 定义一个函数，使它属于 Dog 结构体，实现接口中的 Eat 方法
func (p Pig) Eat() {
	fmt.Println("有人丢了一根骨头给猪吃")
}

// 定义一个函数，使它属于 Dog 结构体，实现接口中的 Name 方法
func (p Pig) Name() string {
	fmt.Println("路上有一只猪，它的名字叫: 猪")
	return "旺财"
}

func testInterface(a Animal) {
	s, ok := a.(Pig)
	if ok {
		fmt.Println(s)
		return
	}
	switch v := a.(type) {
	case *Dog:
		v.Name()
		v.Talk()
		v.Eat()
	case Pig:
		v.Name()
		v.Talk()
		v.Eat()
	default:
		fmt.Println("not support")
	}
}

func main() {
	//testinterface1()
	//var d Dog  // 这样写会报错，因为 Dog 实现接口方法都是由指针类型实现的
	var d *Dog = &Dog{}
	testInterface(d)

	var p *Pig = &Pig{}
	testInterface(p)
	//var h Pig
	//testInterface(h)
}
