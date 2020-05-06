package main

import (
	"fmt"
	"reflect"
)

type S struct {
	A int
	B string
}

// 为 S 结构体添加一个 SetB 方法
func (s *S) SetB(b string) {
	s.B = b
}

func (s *S) Print() {
	fmt.Printf("通过反射调用该方法示例：%#v\n", s)
}

func main() {
	var s S
	v := reflect.ValueOf(&s) // 通过 valueOf 获取结构体类型值相关类型信息，
	t := v.Type()            // 通过值类型相关信息获取结构体类型相关信息

	// t.NumMethod() 通过获取到的结构体类型信息调用 NumMethod() 方法获取该结构体中存在多少个方法,
	// 注：该方法只能获取到大写字母开头的方法
	fmt.Printf("struct S have %d methods\n", t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		// 通过结构体类型信息调用 Method(int下标) 来获取该结构体中的方法对象，然后可以通过该方法对象可以获取该方法的名字与类型
		method := t.Method(i)
		// method.Name 通过方法对象获取方法名, method.Type 通过方法多想获取方法类型
		fmt.Printf("struct %d method, name: %s, type:%v\n", i, method.Name, method.Type)
	}

	// 通过 reflect.ValueOf 获取结构体中的方法调用
	/*
			一、获取方法的两种方式两种方式均返回一个方法对象：
				方式一：Method(下标int)
				方式二：MethodByName("方法名")
			二、通过方法对象调用 call([]value) 调用结构体中的方法，注：call 方法接收一个切片类型的参数,
		       如果调用的方法中没有参数也需要声明一个 []reflect.Value 类型的空切片传入
	*/
	// 方式二示例：通过 MethodByName() 方法调用无参的方法
	m := v.MethodByName("Print")
	var test []reflect.Value
	m.Call(test)
	// 方式二示例：通过 MethodByName() 方法调用有参数的方法
	var test1 []reflect.Value
	mm := v.MethodByName("SetB")
	name := "test"
	nameVal := reflect.ValueOf(name) // 将字符串类型的数据转换为 reflect 类型的数据
	test1 = append(test1, nameVal)   // 将需要传入的参数存入声明的 slice 切片中
	mm.Call(test1)

	m.Call(test)
}

/*
由结果可以看出，结构体的方法是通过一个匿名函数来实现的，匿名海曙的第一个参数则为方法本身(指针类型)，
第二个参数则为结构体方法中接收的参数，如下:
S 结构体 Test 方法: func (s *S)Test(a string){...}
底层实现原理： func(*main.S, string)
并且方法会按照首字母的顺序先后存入到一个 slice 数据中的下标，并不是按照代码的先后顺序，然后按先进先出的原则进行调用
*/
