package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a = 20
	var b = 20
	t.Logf("a = %d, b = %d\n", a, b)
	c := Add(a, b)
	if c != 30 {
		// Fatalf 表示测试失败用来输出测试时输出的错误行内容，并且结束执行测试代码
		t.Fatalf("invalid a + b, c = %d\n", c)
	}
	// Logf 用来输出测试通过后输出的日志内容,打印调试日之后会继续向下执行其它代码
	t.Logf("%d + %d = %d", a, b, c)
}

func TestSub(t *testing.T) {
	var a = 20
	var b = 30
	c := Sub(a, b)
	if c != -10 {
		t.Fatalf("invalid a - b, c = %d\n", c)
	}
	t.Logf("%d - %d = %d", a, b, c)
}
