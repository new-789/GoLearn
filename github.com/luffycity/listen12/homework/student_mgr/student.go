package main

import (
	"fmt"
	"os"
)

// 使用结构体 StudentInfo 定义一个指针切片类型的全局变量,用来存储学生信息
var AllStudents []*StudentInfo

type StudentInfo struct {
	Username string
	Grade    string
	Sex      string
	Score    float32
}

// 菜单显示方法
func showMenu() {
	fmt.Println("========== 输入菜单编号执行相关操作 ===========")
	fmt.Println("	   1. 添加学生信息")
	fmt.Println("	   2. 修改学生信息")
	fmt.Println("	   3. 查看所有学生信息")
	fmt.Println("	   4. 退出程序")
	fmt.Println("===============================================\n\n")
}

// 用来将用户输入的数据使用 StudentInfo 结构体进行初始化操作，并返回一个指针类型的数据
func NewStudent(username string, grade string, sex string, score float32) (student *StudentInfo) {
	// 根据传入的参数通过 StudentInfo 结构体生成一个 student 对象，注意是获取的结构体内存地址，所以是指针类型的对象
	student = &StudentInfo{
		Username: username,
		Grade:    grade,
		Sex:      sex,
		Score:    score,
	}
	return
}

// 声名一个处理用户输入数据的函数
func inputStudent() *StudentInfo {
	var (
		username string
		grade    string
		sex      string
		score    float32
	)
	fmt.Println("please input username>>: ")
	// 通过 Scanf 标准输入方法获取用户输入的内容，并存储在变量中
	_, _ = fmt.Scanf("%s\n", &username)
	fmt.Println("please input grade [0-6]>>: ")
	_, _ = fmt.Scanf("%s\n", &grade)
	fmt.Println("please input sex [男|女]>>: ")
	_, _ = fmt.Scanf("%s\n", &sex)
	fmt.Println("please input score [0-100]>>: ")
	_, _ = fmt.Scanf("%f\n", &score)
	// 调用 NewStudent 函数并传参，该函数返回一个指针类型的数据
	student := NewStudent(username, grade, sex, score)
	return student
}

// 添加学生内容函数
func addStudent() {
	// 调用 inputStudent 函数处理用户输入内容
	student := inputStudent()
	// 循环切片用来判断当前用户是否存在并做相应处理
	for index, val := range AllStudents {
		// 判断如果用户存在则更新该学生的信息
		if val.Username == student.Username {
			fmt.Printf("user [%s] success update\n", student.Username)
			AllStudents[index] = student
			return
		}
	}
	AllStudents = append(AllStudents, student) // 存数据到切片
	fmt.Printf("user [%s] success insert\n", student.Username)
}

// 修改学生内容函数
func modifyStudent() {
	// 调用 inputStudent 函数处理用户输入内容
	student := inputStudent()
	// 循环切片用来判断当前用户是否存在并做相应处理
	for index, val := range AllStudents {
		// 判断学生如果存在则更新学生信息
		if val.Username == student.Username {
			fmt.Printf("user [%s] success modify\n", student.Username)
			AllStudents[index] = student
			return
		}
	}
	// 如果不存在则打印提示信息许需要修改的学生不存在,然后让其继续输入
	fmt.Printf("sorry~! your input [%s] is not found\n", student.Username)
}

// 显示学生信息内容函数
func showStudent() {
	for _, value := range AllStudents {
		fmt.Printf("user:%s info value:%#v\n", value.Username, value)
	}
}

func start() {
	// 不断循环，用来实现操作完一个功能继续输出菜单可持续操作
	for {
		showMenu()
		var userInput int
		_, _ = fmt.Scanf("%d\n", &userInput)
		// 根据用户输入的内容进行判断调用不同的函数方法
		switch userInput {
		case 1:
			addStudent()
		case 2:
			modifyStudent()
		case 3:
			showStudent()
		case 4:
			// os 标准包退出程序方法
			os.Exit(0)
		}
	}
}

func main() {
	start()
}
