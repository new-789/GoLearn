package main

import (
	"fmt"
)

func main() {
	var a map[string]int = map[string]int{
		"stu01": 100,
		"stu02": 200,
		"stu03": 300,
	}
	fmt.Printf("map a is len:%d\n", len(a))
}
