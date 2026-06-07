package main

import "fmt"

// struct
type User struct {
	ID       uint64
	Username string
	Birthday string
}

// ... existing code ...
type Order struct {
	OrderNo string
	UID     uint64
	Money   float64
}

func main() {
	// 1. 自定义全新类型
	type Age int
	var a Age = 20
	var raw int = 30
	//a = raw // 编译报错，类型不匹配
	a = Age(raw)
	fmt.Println("自定义Age：", a)

	// 2. 类型别名 =
	type IntAlias = int
	var x IntAlias = 100
	var y int = x // 直接赋值，无需转换
	fmt.Println("别名互通：", x, y)

	// 对比区分
	type UniqueInt int
	type AliasInt = int
	var u UniqueInt = 1
	var al AliasInt = 2
	var origin int = al
	// origin = u // 报错
	fmt.Println(u, al, origin)

	// User 示例
	user := User{
		ID:       1,
		Username: "张三",
		Birthday: "2000-01-01",
	}
	fmt.Println("用户信息：", user)
	fmt.Println("用户名：", user.Username)

	// Order 示例
	order := Order{
		OrderNo: "ORD20250604001",
		UID:     user.ID,
		Money:   99.129456,
	}
	fmt.Println("订单信息：", order)
	fmt.Printf("订单号：%s, 金额：%.2f\n", order.OrderNo, order.Money)

	// 修改字段
	user.Username = "李四"
	order.Money = 199.50
	fmt.Println("修改后用户：", user)
	fmt.Println("修改后订单金额：", order.Money)
}
