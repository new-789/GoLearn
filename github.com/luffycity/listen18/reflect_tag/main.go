package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	// 在结构体中声明一个字段，并为该字段定义 tag 信息,注意两个数之间没有 , 号；
	Name string `json:"name" db:"testDb"`
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) Print() {
	fmt.Printf("这是通过反射进行调用的结果: %#v\n", s)
}

func main() {
	var s Student
	// 获取 Student 结构体值类型信息
	v := reflect.ValueOf(&s)
	// 通过值类型信息对象获取 Student 结构体的类型信息
	t := v.Type()
	// 通过类型对象调用 Field(下标int) 方法获取第零个字段，返回一个字段对象
	field0 := t.Elem().Field(0)
	// 通过字段对象调用 Tag 结构体中的 Get(key名称) 方法获取反射对象中的 Tag 信息
	fmt.Printf("tag json = %s\n", field0.Tag.Get("json"))
	fmt.Printf("tag json = %s\n", field0.Tag.Get("db"))
}
