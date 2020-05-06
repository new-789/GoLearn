package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Sex   int
	Score float32
}

func main() {
	var s Student
	v := reflect.ValueOf(&s) // 获取值类型相关信息，由于是更改值操作所以只能传一个地址
	// 在反射中只能通过 Elem() 来获取传入的内存地址，然后通过 Field 传入下标来获取需要设置的字段，
	v.Elem().Field(0).SetString("stu01")
	// FieldName 方法直接获取到反射结构体中的字段名，参数为需要获取的字段名
	v.Elem().FieldByName("Age").SetInt(88)
	v.Elem().FieldByName("Sex").SetInt(0)         // 设置 Int 类型字段的值
	v.Elem().FieldByName("Score").SetFloat(99.99) // 设置浮点类型字段的值
	fmt.Printf("%#v\n", s)
}
