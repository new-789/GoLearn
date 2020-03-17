package main

import "fmt"

// 定义一个类型 Animal 为结构体
type Animal struct {
	Name string
	Sex  string
}

// 定义一个类型 PuruAnimal 为结构体
type PuruAnimal struct {
}

// 为 Animal 类定义一个 Talk 方法
func (a *Animal) Talk() {
	fmt.Printf("I'm Talk, I'm is %s\n", a.Name)
}

// 为 PuruAnimal 类定义一个 Talk 方法
func (p *PuruAnimal) Talk() {
	fmt.Println("I'm is PuruAnimal")
}

// 定义一个类型 Dog 为结构体
type Dog struct {
	Feet string
	// 在 Dog 类型中继承 Animal 类型中的所有字段
	*Animal
	// 在 Dog 类型中继承 PuruAnimal 类型中的所有字段
	*PuruAnimal
}

// 为 Dog 类型定义一个方法
func (d *Dog) Eat() {
	fmt.Println("Dog is eat")
}

func main() {
	// 通过 Dog 类型实例化 d 为指针类型对象，并在其中对继承的两个类型进行初始化
	var d *Dog = &Dog{
		Feet: "四条腿",
		// 实例化 Animal 类型
		Animal: &Animal{},
		// 实例化 PuruAnimail 类型
		PuruAnimal: &PuruAnimal{},
	}
	d.Name = "dog"
	d.Sex = "公狗"
	// 通过 Dog 实例化的 d 调用 Animal 类型中的字段并赋值
	d.Eat()
	// 通过 Dog 实例化的 d 加类型名调用 Animal 类型的 Talk 方法
	d.Animal.Talk()
	// 通过 Dog 实例化的 d 加类型名调用 PuruAnimal 类型的 Talk 方法
	d.PuruAnimal.Talk()
}
