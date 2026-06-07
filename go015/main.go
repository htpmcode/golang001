package main

import "fmt"

type Book1 struct {
	Title string
	Price float64
}

// 1. 值接收者：操作副本，原数据不变
func (b Book1) ChangeTitle(newTitle string) {
	b.Title = newTitle
	fmt.Println("方法内标题：", b.Title)
}

// 2. 指针接收者：操作原结构体
func (b *Book1) ChangePrice(newPrice float64) {
	b.Price = newPrice
}

func main() {
	b := Book1{Title: "Go编程", Price: 59.9}

	// 值接收者调用
	b.ChangeTitle("Java编程")
	fmt.Println("调用后原标题：", b.Title) // 原值不变

	// 指针接收者调用
	b.ChangePrice(49.9)
	fmt.Println("调用后原价格：", b.Price) // 原值被修改

	// 指针变量调用两种方法（语法糖兼容）
	b2 := &Book1{Title: "Python", Price: 39.9}
	b2.ChangeTitle("C++")
	b2.ChangePrice(29.9)
	fmt.Println(*b2)
}
