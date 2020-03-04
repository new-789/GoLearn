package main

import (
	"fmt"
)

func mapSlice() {
	var mapslice []map[string]int = make([]map[string]int, 5, 10)
	for index, val := range mapslice {
		fmt.Printf("slice[%d]=%v\n", index, val)
	}
	fmt.Println()
	mapslice[0] = make(map[string]int, 5)

}

func main() {
	mapSlice()
}
