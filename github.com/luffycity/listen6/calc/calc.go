package calc

var a int = 100 // 小写变量，表示为私有，仅在包内可使用
var A int = 200 // 大写变量，表示公有，可在任意地方使用

func Add(a, b int) int { // 大写开头表示公有，可在任意包内调用
	return a + b
}

func sub(a, b int) int { // 小写开头表示私有，仅在当前包内可调用
	return a - b
}
