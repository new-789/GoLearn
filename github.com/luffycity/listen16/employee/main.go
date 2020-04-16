package main

import (
	"fmt"
)

// 定义一个接口，名称为 Employer(雇主)，并在其中为其定义一个 CalSalary 薪水计算方法
type Employer interface {
	CalSalary() float32 // 在接口中定义薪水计算方法，并返回一个 float32 类型的数据
}

// 定义一个结构体，名称为 ProRamer(程序员)，并在其中定义几个字段(名字，固定薪资，奖金系数)
type Programer struct {
	Name  string
	Base  float32
	Extra float32
}

// 为 Programer(程序员) 结构体定义一个 CalSalary 方法
func (p *Programer) CalSalary() float32 {
	return p.Base
}

// 为程序员职位定义一个构造函数，并返回一个 Programer 结构体对象的内存地址
func NewProgramer(name string, base float32, extra float32) *Programer {
	return &Programer{
		Name:  name,
		Base:  base,
		Extra: extra,
	}
}

// 定义一个结构体，名称为 Sale(销售)，并在其中定义几个字段(名字，固定薪资，奖金系数)
type Sale struct {
	Name  string
	Base  float32
	Extra float32
}

// 为 Sale(销售) 结构体定义一个 CalSalary 方法
func (s *Sale) CalSalary() float32 {
	return s.Base + s.Extra*s.Base*0.5
}

// 为销售邓毅一个构造方法，并返回一个 Salary 结构体对象的内存地址
func NewSale(name string, base float32, extra float32) *Sale {
	return &Sale{
		Name:  name,
		Base:  base,
		Extra: extra,
	}
}

// 定义一个 calcAll 方法接收一个地址类型对切片数据，用来计算总花费的薪水
func CalcAll(e *[]Employer) float32 {
	var cost float32
	for _, val := range *e {
		cost += val.CalSalary()
	}
	return cost
}

func main() {
	p1 := NewProgramer("搬砖工1", 13000.0, 0)
	p2 := NewProgramer("搬砖工2", 12000.0, 0)
	p3 := NewProgramer("搬砖工3", 11000.0, 0)

	s1 := NewSale("销售1", 8000.0, 2.5)
	s2 := NewSale("销售2", 9000.0, 3.0)
	s3 := NewSale("销售3", 7500.0, 2.0)

	// 由于员工都实现了 Employer 接口类型中的方法，所以此处定义一个接口类型的切片，用来存储所有员工结构体对象的内容
	var employeeList []Employer
	employeeList = append(employeeList, p1)
	employeeList = append(employeeList, p2)
	employeeList = append(employeeList, p3)

	employeeList = append(employeeList, s1)
	employeeList = append(employeeList, s2)
	employeeList = append(employeeList, s3)

	// 调用 CalcAll 方法用来计算总支出
	cost := CalcAll(&employeeList)
	fmt.Printf("这个月总人力成本为：%f 元\n", cost)
}
