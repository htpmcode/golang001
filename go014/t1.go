package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	Username string `json:"username"` // json别名
	Age      int    `json:"age,omitempty"`
	IsVip    bool   `json:"is_vip"`
	secret   string // 小写私有字段，无法序列化
}

type User2 struct {
	Username string
	Age      int
	IsVip    bool
	secret   string
}

func main() {
	u2 := User1{
		Username: "张三",
		Age:      20,
		IsVip:    true,
	}

	// ========== 五、JSON序列化&反序列化 ==========
	// 序列化：结构体→json
	jsonByte, err := json.MarshalIndent(u2, "", "  ")
	if err == nil {
		fmt.Println("\n用户json：")
		fmt.Println(string(jsonByte))
	}

	// 反序列化：json→结构体
	//	jsonStr := `{"username":"反序列化用户","age":25,"is_vip":true}`
	//	var unUser User
	//	_ = json.Unmarshal([]byte(jsonStr), &unUser)
	//	fmt.Println("反序列化结果：", unUser)
	//
}
