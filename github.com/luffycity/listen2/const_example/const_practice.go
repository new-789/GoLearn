package main

import (
	"fmt"
)

//
const (
	A = iota
	B
	C
	D = 8
	E
	// 由于 iota 每隔一行它的默认值会增加一位，所以 e 在常量的第五行，所以 F 的值不会在上一个 iota 基础之上继续递增而是重新返回当前行的默认值所以值为 5
	F = iota
	G // G 的值会在它上一个常量之上继续递增
)

//
const (
	// 由于 iota 的值仅在同一个 const 中有效，所以在该 const 中 iota 的默认值值被重置为 0
	H1 = iota
	H2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H1)
	fmt.Println(H2)
}
