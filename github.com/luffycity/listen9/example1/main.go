package main

import (
	"fmt"
)

func main() {
	var sa []string = make([]string, 5, 10)
	fmt.Printf("sa:%v\n", sa)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%d", i))
	}
	fmt.Println("sa:", sa)
}
