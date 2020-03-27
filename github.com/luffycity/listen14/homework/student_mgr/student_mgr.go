package student_mgr

import (
	"fmt"
)

type Students struct {
	Username string
	Sex      int
	Age      int
	Score    float32
	Grade    string
}

type StudentMgr struct {
	AllStudent []*Students
}

func (s *StudentMgr) AddStudent(Stu *Students) (err error) {
	for index, val := range s.AllStudent {
		if val.Username == Stu.Username {
			s.AllStudent[index] = Stu
			fmt.Printf("user %s success update\n", val.Username)
			return
		}
	}
	s.AllStudent = append(s.AllStudent, Stu)
	fmt.Printf("user %s success insert\n", Stu.Username)
	return
}

func (s *StudentMgr) ModifyStudent(Stu *Students) (err error) {
	for index, val := range s.AllStudent {
		if Stu.Username == val.Username {
			s.AllStudent[index] = Stu
			fmt.Printf("user %s modify success\n", Stu.Username)
			return
		}
	}
	fmt.Printf("your input %s is not found\n", Stu.Username)
	// 返回一个错误码语法
	return fmt.Errorf("Yous input [%s] is not found", Stu.Username)
}

func (s *StudentMgr) ShowStudent() {
	for _, val := range s.AllStudent {
		fmt.Printf("user [%s] is value: %#v\n", val.Username, val)
	}
	fmt.Println()
}
