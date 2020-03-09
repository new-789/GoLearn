package main

import (
	"encoding/json" // 导入 json 标准包
	"fmt"
)

type Users struct {
	// json 后面的 username 既是通过 json 序列化反射后得到的字段名,注：tag 的 key 常设为 json
	Username string  `json:"username"`
	Sex      string  `json:"sex"`
	Score    float32 `json:"score"`
	age      int32   // 由于首字母是小写所以在进行 json 反射序列化时无法获取到 age 数据
}

func main() {
	user := &Users{
		Username: "user01",
		Sex:      "男",
		Score:    88.88,
		age:      88,
	}
	// 使用 json 包对 user 对象进行 json 序列化
	data, _ := json.Marshal(user)
	fmt.Printf("json_data_str:%s\n", string(data))
}
