package main

import (
	"fmt"
)

func testIf() {
	var num int = 3
	if num == 1 {
		fmt.Printf("num = 1\n")
	} else if num == 2 {
		fmt.Printf("num = 2\n")
	} else if num == 3 {
		fmt.Printf(" num = 3\n")
	} else if num == 4 {
		fmt.Printf("num = 4\n")
	} else {
		fmt.Printf("num = 5\n")
	}
}

func testSwitch() {
	var num int = 3
	switch num {
	case 1:
		fmt.Printf("num = 1\n")
	case 2:
		fmt.Printf("num = 2\n")
	case 3:
		fmt.Printf("num = 3\n")
	case 4:
		fmt.Printf("num=4\n")
	case 5:
		fmt.Printf("num = 5\n ")
	}
}

func testSwitch2() {
	switch num := 10; num {
	case 1:
		fmt.Printf("num=1\n")
	case 2:
		fmt.Printf("num=2\n")
	case 3:
		fmt.Printf("num=3\n")
	case 4:
		fmt.Printf("num=4\n")
	default:
		fmt.Printf("num=%d\n", num)
	}
	// fmt.Printf(num) // 此处调用 num 变量则会提示不存在
}

func getValue() int {
	return 2
}

func testSwitch3() {
	switch getValue() { // 直接调用函数并进行判断
	case 1:
		fmt.Printf("value is %d\n", getValue())
	case 3:
		fmt.Printf("value is %d\n", getValue())
	default: // 上面的 case 都不满足则走该分支
		fmt.Printf("value is %d\n", getValue())
	}
}

func testSwitch4() {
	switch varstr := "i"; varstr {
	case "a", "b", "c", "d", "i", "o": //判断如果值在这其中的一个则执行该分支
		fmt.Printf("variable is %s\n", varstr)
	default:
		fmt.Printf("variable is %s\n", varstr)
	}
}

func testSwitch5() {
	switch vint := 75; {
	case vint >= 30 && vint <= 50:
		fmt.Printf("value greater than 30 and value less than 50")
	case vint >= 55 && vint <= 70:
		fmt.Printf("value greater than 55 and value less than 70")
	case vint >= 75 && vint <= 90:
		fmt.Printf("value greater than 75 and value less than 90")
	}
}

func testSwitch6() {
	switch vint := 75; {
	case vint >= 30 && vint <= 50:
		fmt.Printf("value greater than 30 and value less than 50\n")
	// 此处判断条件表达式判断正确，正常来自说执行到这里就应该直接跳出 switch 语句了
	case vint >= 55 && vint <= 75:
		fmt.Printf("value greater than 55 and value less than 70\n")
		fallthrough // 此处加了 fallthrough 关键字，则直接穿透到到下一个 case 中并继续执行代码
	case vint >= 80 && vint <= 95:
		fmt.Printf("value greater than 75 and value less than 90\n")
	}
}

func testMulti() {
	// 1*1 =1
	// 1*2 =2 2*2 =4
	// 1*3 =3 2*3 =6 3*3 =9
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// testIf()
	// testSwitch()
	// testSwitch2()
	// testSwitch3()
	// testSwitch4()
	// testSwitch5()
	// testSwitch6()
	testMulti()
}
