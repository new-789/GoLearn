package main

import (
	"fmt"
	"reflect"
)

func reflect_example(a interface{}) {
	/*声名一个函数，该函数的参数为一个空接口*/
	x := reflect.TypeOf(a) // 通过 reflect 包判断空接口 a 接收了一个什么类型的数据
	fmt.Printf("type of a is :%v\n", x)
	// 通过返回的 x 对象获取空接口 a 中存储变量的类型
	varType := x.Kind()
	// 通过断言明确类型
	switch varType {
	case reflect.Int64:
		fmt.Printf("varType is int64 ")
	case reflect.String:
		fmt.Printf("varType is int32 ")
	case reflect.Float64:
		fmt.Printf("varType is float64 ")
	}
}

func reflectValue(a interface{}) {
	// 从空接口 a 中通过反射获取值类型信息
	v := reflect.ValueOf(a)
	// 通过值类型信息返回的对象获取它的类型信息，等同于 reflect.TypeOf()
	v.Type()
	// 通过值类型信息返回的对象获取它的值类型，然后可以通过断言获取具体的值
	k := v.Kind()
	// 通过断言获取值类型信息中存储的具体的值
	switch k {
	case reflect.Int64:
		// v.Int() 获取值类型对象中类型为 Int 类型的值，以同时获取 Int32 和 Int64 的值
		fmt.Printf("a is int64, store value is %d\n", v.Int())
	case reflect.Float64:
		// v.Float() 获取值类型对象中类型为浮点数类型的值,Float 可以同时获取 Float32 和 Float64 的值
		fmt.Printf("a is float64, store value is :%F\n", v.Float())
	}
}

func reflectSetValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		v.SetInt(88)
	case reflect.Float64:
		v.SetFloat(88.88)
	case reflect.Ptr:
		// 通过 v.Elem() 方法获取 v 指向的内存地址，然后对其进行赋值操作
		v.Elem().SetFloat(88.88)
	}
}

func main() {
	var x float64 = 3.4
	//reflect_example(x)
	//reflectValue(x)
	reflectSetValue(&x)
	fmt.Println("x change value is :", x)
}
