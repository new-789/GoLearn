package main

import (
	"fmt"
	"strings"

	// 导入自定义包
	"github.com/luffycity/listen2/access"
)

// bool 数据类型示例
func testBool() {
	var a bool // 定义布尔类型变量,默认为 false
	fmt.Println(a)
	a = true // 定义布尔类型变量并赋值
	fmt.Println(a)

	a = !a // 取反
	fmt.Println(a)
	var b bool = true
	if a == true && b == true { // && 与运算符示例
		fmt.Println("&& right")
	} else {
		fmt.Println("&& not right")
	}

	if a == true || b == true { // || 或运算符示例
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not right")
	}

	fmt.Printf("%t %t\n", a, b) // 格式化输出占位符示例
}

// 数值类型示例
func testInt() {
	// 正数类型示例
	var a int8 // int8 占一个字节，取值范围 -128 --- 127 之间
	a = 18
	fmt.Println("a=", a)
	a = -18
	fmt.Println("a=", a)
	var b int // int 类型占用字节取决于操作系统
	b = 23423423423
	fmt.Println("b=", b)

	// 浮点数类型示例,在定义浮点数类型的变量时类型 float32 和 float64 可以省略不写
	var c float32 // 浮点数默认值为 0
	fmt.Println("c=", c)
	var d float32 = 3.1415926
	fmt.Println("d=", d)
}

func add() {
	var a int = 18  // int 类型数据
	var b int8 = -3 // int8 类型数据
	var c int = -12
	var d float32 = 3.14
	c = int(d)
	fmt.Println(a+int(d), c)
	fmt.Printf("a=%d b=%x c=%x d=%f\n", a, b, c, d)
}

func teststr() {
	var a string // 定义空字符串
	fmt.Println(a)
	var b = "hello world" // 定义字符串并赋值
	fmt.Println(b)
	c := "c:hello" // 直接定义变量语法，等于 var c string = ":=符号"
	fmt.Println(c)
	d := 222                      // 使用 := 符号直接定义一个变量，等于 var d int = 222
	fmt.Printf("a=%v d=%v", a, d) // 万能占位符的使用
	u := "u=\nhello"
	s := `
		煮豆燃豆萁，
		豆在釜中泣，
	`
	i := `
		本是同根生，
		相煎何太急。
	`
	fmt.Println(u, s)
	// 查看字符串长度示例
	// slen := len(s)
	fmt.Printf("这是字符串的长度：%d", len(s))
	// 字符串拼接示例
	s = s + i // 拼接字符串方式一
	fmt.Println(s)
	s = fmt.Sprintf("%s%s", s, i) // 拼接字符串方式二
	fmt.Println(s)
	// 字符串分割示例
	ips := "192.169.0.33;192.168.1.23;192,168.1.35"
	ipArray := strings.Split(ips, ";")
	fmt.Println(ipArray)
	fmt.Println(ipArray[1])
	// 包含示例
	ipss := "192.169.0.33;192.168.1.23;192,168.1.35"
	result := strings.Contains(ipss, "192.168.1.23")
	if result {
		fmt.Printf("存在：%t\n", result)
	} else {
		fmt.Printf("不存在：%t\n", result)
	}
	// 前缀和后缀判断
	// 前缀判断
	str := "https://www.baidu.com"
	if strings.HasPrefix(str, "som") {
		fmt.Printf("str is  https url\n")
	} else {
		fmt.Printf("str is not https url\n")
	}
	// 判断后缀结尾
	if strings.HasSuffix(str, "baidu.com") {
		fmt.Printf("url is baidu web\n")
	} else {
		fmt.Printf("url is not baidu web\n")
	}
	// 判断子串中的位置
	str1 := "www.baidu baidu.com"
	index := strings.Index(str1, "baidu")
	fmt.Printf("baidu 第一次出现的位置是：%d\n", index)
	lastindex := strings.LastIndex(str1, "baidu")
	fmt.Printf("baidu 最后一次出现的位置：%d\n", lastindex)
	// join 操作示例
	var strArr []string = []string{"192.168.0.1", "192.168.1.1"}
	strArray := strings.Join(strArr, ";")
	fmt.Printf("strArray is:%s\n", strArray)
}

// 操作符示例
func testOperator() {
	a := 2
	if a == 2 {
		fmt.Print("is right\n")
	} else {
		fmt.Print("not right\n")
	}

	a = a + 98
	fmt.Printf("a=%d\n", a)

}

func testAccess() {
	// 其它包中小写变量导入使用报错
	// fmt.Printf("access_a=%d", access.a)
	fmt.Printf("access_A=%d", access.A)
}

func main() {
	// testBool()
	// testInt()
	// add()
	// teststr()
	// testOperator()
	testAccess()
}
