package main

import (
	"fmt"
)

// 声名 Teacher 接口
type Teacher interface {
	ReturnName() string
}

// 声名 School 接口
type School interface {
	ReturnSchoolName() string
}

// 声名一个结构体，定义为 teacher 类型
type Teach struct {
	Name   string
	School string
}

// 实现 Teacher 接口中的 ReturnName 方法，该方法属于 Teach 结构体
func (t *Teach) ReturnName() string {
	return t.Name
}

// 实现 School 接口中的 ReturnSchoolName 方法，该方法属于 Teach 结构体
func (t *Teach) ReturnSchoolName() string {
	return t.School
}

func main() {
	// 声名 Teacher 接口声名变量 t1
	var t1 Teacher
	// 声名 School 接口声名变量 s1
	var s1 School

	// 使用结构体初始化一个对象 t
	var t *Teach = &Teach{
		Name:   "易中天",
		School: "中南大学",
	}

	// 将类型变量存储在接口变量 t1 中
	t1 = t
	// 将类型变量存储在接口变量 s1 中
	s1 = t

	// 调用 Teacher 结构体中的 ReturnName 方法
	fmt.Println("teacher name is :", t1.ReturnName())
	// 调用 School 结构体中的 ReturnSchoolName 方法
	fmt.Println("teacher in school :", s1.ReturnSchoolName())
}
