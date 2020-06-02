package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "hello,"
	ch <- "world!"

	s1 := <-ch
	s2 := <-ch
	s3 := <-ch // 取出的内容空了之后
	fmt.Println(s1, s2, s3)
}
