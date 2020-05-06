package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	Sex    int
	Score  float32
	School string // 小写开头为私有类型的字段
}

func main() {
	var s Student
	// 通过 reflect.ValueOf 获取结构你变量 s 中类型的相关信息
	v := reflect.ValueOf(s)
	// 通过值类型信息获取类型相关的信息,通过该类型信息可以获取反射结构体中的字段名,语法：type_obg.Field(i int).Name
	t := v.Type() // 等同于 t := reflect.TypeOf()
	//t := reflect.TypeOf(s)  // == t := v.Type() 只不过此方法更容易获取到类型信息
	// 通过 Kind 获取结构体变量中存储的为什么类型的字段
	kind := t.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Println("s is int64")
	case reflect.Float64:
		fmt.Println("s is Float64")
	// reflect.Struct 通过反射断言是否为一个结构体
	case reflect.Struct:
		fmt.Println("s is struct")
		// v.NumField() 用来在映射中获取结构体中的字段数量，用于方法 valueType_obj.NumField()
		fmt.Printf("field num of s is %d\n", v.NumField())
		// 遍历结构体中的所有字段
		for i := 0; i < v.NumField(); i++ {
			/* valueType_obj.Field(i int) 获取反射结构体中的字段信息, 该方法接收一个 int 类型的参数
			通过该方法返回的对象可以获取字段中存储的类型信息以及值信息，
			获取结构体中字段类型信息方法：field_obj.Type()
			获取结构体中字段存储的值内容方法： field_obj.Interface()
			*/
			field := v.Field(i)
			fmt.Printf("name:%s, type: %v, value: %v\n",
				t.Field(i).Name, field.Type(), field.Interface())
		}
	default:
		fmt.Println("default")
	}

}
