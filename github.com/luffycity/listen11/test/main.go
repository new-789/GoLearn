package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "早上好!下午好!中午好!晚上好!"
	result := strings.Split(str, "!")
	fmt.Printf("result:%v", result)
}
