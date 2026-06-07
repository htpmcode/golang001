package main

import "fmt"

type Book struct {
	Title string
	Price float64
}

// 值接收：拷贝副本
func (b Book) ChangeTitle(newTitle string) {
	b.Title = newTitle
	fmt.Println("方法内副本标题：", b.Title)
}

// 指针接收：操作原内存
func (b *Book) ChangePrice(newPrice float64) {
	b.Price = newPrice
}

func main() {
	// 情况1：普通结构体变量调用
	b := Book{Title: "Go编程", Price: 59.9}
	b.ChangeTitle("Java编程")       // 值接收：复制副本，改副本
	fmt.Println("原值标题：", b.Title) // Go编程，没变

	b.ChangePrice(49.9)           // 指针接收，语法糖自动&b，修改原数据
	fmt.Println("原值价格：", b.Price) //49.9，修改成功

	fmt.Println("-----------------")

	// 情况2：结构体指针调用
	b2 := &Book{Title: "Python", Price: 39.9}
	b2.ChangeTitle("C++") // 值接收：*b2做值拷贝，改副本，原数据不变
	b2.ChangePrice(29.9)  // 指针接收：直接操作原结构体
	fmt.Println(*b2)      // {Python 29.9} 标题没变、价格变了
}

//核心结论（你总结完全正确）
//能不能修改原结构体，只看「方法接收者是值还是指针」，和调用方是结构体变量 / 结构体指针无关；Go 自带语法糖自动适配。
//值接收者 (b Book)：永远拷贝一份结构体副本，方法内部修改的是副本，无法改动外面原始数据。
//指针接收者 (b *Book)：拿到原结构体地址，无论用值变量还是指针变量调用，全都能修改原数据。
