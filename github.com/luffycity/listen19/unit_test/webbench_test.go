package main

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// 压力测试方法，必须有 Benchmark 开头，后面的名称首字母必须大写
func BenchmarkAdd(b *testing.B) {
	// b.N， N为一个整数，具体的值是自动化框架确定的我们直接使用即可
	for i := 0; i < b.N; i++ {
		var a int = 30
		var c int = 20
		add(a, c)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a int = 30
		var b int = 10
		sub(a, b)
	}
}
