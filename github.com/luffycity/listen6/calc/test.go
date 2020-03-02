package calc

func Test() int {
	return a // 在此处调用 calc/calc.go 同一包内的私有变量 a 并返回
}
