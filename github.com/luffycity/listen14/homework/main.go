package main

import (
	"fmt"
	"github.com/luffycity/listen14/homework/student_mgr"
	"os"
)

var (
	studentMgr = &student_mgr.StudentMgr{}
)

func NewStudent(username string, sex int, age int, score float32, grade string) *student_mgr.Students {
	var stu *student_mgr.Students = &student_mgr.Students{
		Username: username,
		Sex:      sex,
		Age:      age,
		Score:    score,
		Grade:    grade,
	}
	return stu
}

func Menu() {
	fmt.Println("========== 输入菜单编号执行相关操作 ===========")
	fmt.Println("	   1. 添加学生信息")
	fmt.Println("	   2. 修改学生信息")
	fmt.Println("	   3. 查看所有学生信息")
	fmt.Println("	   4. 退出程序")
	fmt.Println("===============================================\n")
}

func Start() {
	for {
		Menu()
		var ipt int
		_, _ = fmt.Scanf("%d\n", &ipt)
		switch ipt {
		case 1:
			stu := InputStudent()
			studentMgr.AddStudent(stu)
		case 2:
			stu := InputStudent()
			studentMgr.ModifyStudent(stu)
		case 3:
			studentMgr.ShowStudent()
		case 4:
			os.Exit(0)
		}
	}
}

func InputStudent() *student_mgr.Students {
	var (
		username string
		sex      int
		age      int
		score    float32
		grade    string
	)
	fmt.Println("Please input username>>:")
	_, _ = fmt.Scanf("%s\n", &username)
	fmt.Println("Please input sex[0|1]>>:")
	_, _ = fmt.Scanf("%d\n", &sex)
	fmt.Println("Please input score>>:")
	_, _ = fmt.Scanf("%v\n", &score)
	fmt.Println("Please input age[18-30]>>:")
	_, _ = fmt.Scanf("%d\n", &age)
	fmt.Println("Please input grade>>:")
	_, _ = fmt.Scanf("%s\n", &grade)
	Student := NewStudent(username, sex, age, score, grade)
	return Student
}

func main() {
	Start()
}
