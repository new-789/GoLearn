package main

import (
	"fmt"
)

var (
	coins = 50 // 声名一个常量记录金币总数量
	// 声名一个切片，用来存放所有用户
	users []string = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Aaron", "Elizabeth",
	}
	// 声名一个 map 用来存放最后每个用户分到的金币数
	distribution = make(map[string]int, len(users))
)

// 声名一个函数接收字符串类型的用户名，用来遍历用户的名，并进行判断用户名中所包含的字母进行金币的分配
func calcCoin(username string) int {
	// 声名一个变量用来存储分得的金币总数，根据传入的人员名不同则 sumCoin 变量的内容也不同
	var sumCoin int
	// 遍历用户名中的每个字符
	for _, char := range username {
		// 类似 if ... else {} 语句，可在 switch 作用域中判断多个值
		switch char {
		case 'a', 'A': // 判断是否为 a 或者 A，是则记录在 sumCoin 变量中
			sumCoin += 1
		// 判断是否为 e 或者 E，是则覆盖 sumCoin 变量中的值，依次来记录当前用户名分得的金币
		case 'e', 'E':
			sumCoin += 1
		case 'i', 'I':
			sumCoin += 2
		case 'o', 'O':
			sumCoin += 3
		case 'u', 'U':
			sumCoin += 5
		}
	}
	return sumCoin
}

// 声名一个函数，遍历存储用户名的切片，并拿到分金币后的结果然后存储在定义的 distribution map 中,
// 并返回剩余的金币
func dispatchCoin() int {
	// 循环切片中的所有人名
	for _, username := range users {
		// 调用 calcCoin 函数，获取每个人的金币数
		allCoin := calcCoin(username)
		// 总金币减去每个人获得的金币数得到剩余的金币，每循环一次则覆盖掉一次 coins 变量的值，最终为分配完之后剩余的金币
		coins -= allCoin
		// 获取 map 中当前 key 用户名的值，以及根据 key 的存在与否返回一个状态 true/false
		value, ok := distribution[username]
		// 判断装备如果为 false 则往 map 中插入一条用户数据
		if !ok {
			distribution[username] = allCoin
		} else {
			distribution[username] = allCoin + value
		}
	}
	return coins
}

func main() {
	surplusCoin := dispatchCoin()
	for username, coin := range distribution {
		fmt.Printf("[%s] 分到 %d 枚金币\n", username, coin)
	}
	fmt.Printf("还剩下 [ %d ] 枚金币\n", surplusCoin)
}
