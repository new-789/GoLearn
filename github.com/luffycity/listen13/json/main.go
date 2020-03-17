package main

import (
	"encoding/json"
	"fmt"
)

// 声名一个 Student 类型的数据
type Students struct {
	Id   int
	Name string
	Sex  string
	Age  int
}

// 声名一个 Class 类型的数据
type Class struct {
	Name  string
	Count int
	// 定义为一个切片数据，类型为 Student 类型的指针
	Students []*Students
}

var rawJson = `{
"Name":"Go 语言全栈","Count":0,"Students":
	[
		{"Id":1000,"Name":"stu-100","Sex":"man","Age":18},
		{"Id":1001,"Name":"stu-101","Sex":"man","Age":19},
		{"Id":1002,"Name":"stu-102","Sex":"man","Age":20},
		{"Id":1003,"Name":"stu-103","Sex":"man","Age":21},
		{"Id":1004,"Name":"stu-104","Sex":"man","Age":22},
		{"Id":1005,"Name":"stu-105","Sex":"man","Age":23},
		{"Id":1006,"Name":"stu-106","Sex":"man","Age":24},
		{"Id":1007,"Name":"stu-107","Sex":"man","Age":25},
		{"Id":1008,"Name":"stu-108","Sex":"man","Age":26},
		{"Id":1009,"Name":"stu-109","Sex":"man","Age":27}
	]
}`

func main() {
	// 初始化 class 结构体
	c := &Class{
		Name:  "Go 语言全栈",
		Count: 0,
	}
	// 循环批量添加数据
	for i := 0; i < 10; i++ {
		study := &Students{
			Id:   1000 + i,
			Name: fmt.Sprintf("stu-%d", 100+i),
			Sex:  "man",
			Age:  18 + i,
		}
		// 给继承的 Student 指针类型的切片中添加数据
		c.Students = append(c.Students, study)
		c.Count = len(c.Students)
	}
	// 将 class 结构体对象 c 通过 json 包序列化为 json 格式，返回一个 bytes 类型的切片的数据和错误信息
	// 返回的数据 data 默认为 ASCLL 码类型的内容，所以还需要使用 string 进行转换才是人类易读的格式
	data, err := json.Marshal(c)
	// 判断如果返回的错误内容中有数据则打印错误结束执行
	if err != nil {
		fmt.Println("Json Marshal failed")
		return
	}
	fmt.Printf("jsonFormat:%s\n", string(data))

	// json 数据反序列化结构体
	fmt.Println()
	var c1 *Class = &Class{}
	// Unmarshal 方法返回一个 error 信息
	err = json.Unmarshal([]byte(rawJson), c1)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return
	}
	fmt.Printf("c1: %#v\n", c1)
	for _, val := range c1.Students {
		fmt.Printf("student:%#v\n", val)
	}
}
