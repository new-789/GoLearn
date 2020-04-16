package main

import "fmt"

// 定义一个 Animal 接口，并在其中定义下面三种方法
type Animal interface {
	Talk()
	Eat()
	Name() string
}

// 定义一个 Describe(描述) 接口，并在其中定义下面一种方法
type Describe interface {
	Describe() string
}

// 定义一个 AdvanceAnimal 接口，并将 Animal 接口和 Describe 接口嵌套在其中，
// 此时 AdvanceAnimal 接口中就同时包含了两个嵌套接口中的方法, 使用类型实现接口方法时则需要全部实现；
type AdvanceAnimal interface {
	Animal
	Describe
}

type Dog struct {
}

func (d *Dog) Eat() {
	fmt.Println("旺财在吃东西")
}

func (d *Dog) Talk() {
	fmt.Println("旺财在叫: 汪汪汪")
}

func (d *Dog) Name() string {
	fmt.Println("狗的名字叫: 旺财")
	return "狗的名字叫旺财"
}

func (d *Dog) Describe() string {
	fmt.Println("旺财告诉我们说:'这是接口嵌套方法示例'")
	return "接口嵌套方法示例"
}

func main() {
	var a AdvanceAnimal

	var d *Dog = &Dog{}
	a = d

	a.Name()
	a.Talk()
	a.Eat()
	a.Describe()
}
