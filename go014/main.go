package main

import (
	"fmt"
	"unsafe"
)

// 1. 基础结构体 + 结构体标签
type User struct {
	Username string `json:"username"` // 结构体标签
	Age      int    `json:"age"`
	IsVip    bool   `json:"is_vip"`
}

// 2. 嵌套结构体
type Address struct {
	Province string
	City     string
}

// 具名嵌套
type Person struct {
	Name string
	Addr Address // 嵌套结构体
}

// 匿名内嵌结构体（提升字段访问便捷性）
type Student struct {
	string  // 匿名字段
	Age     int
	Address // 内嵌结构体
}

func main() {
	// 基础结构体初始化
	u := User{
		Username: "张三",
		Age:      20,
		IsVip:    true,
	}
	fmt.Println(u)
	fmt.Println("用户名：", u.Username)

	// 具名嵌套访问
	p := Person{
		Name: "李四",
		Addr: Address{
			Province: "广东",
			City:     "深圳",
		},
	}
	fmt.Println("城市：", p.Addr.City)

	// 匿名内嵌结构体访问
	s := Student{
		string: "王五",
		Age:    18,
		Address: Address{
			Province: "北京",
		},
	}
	fmt.Println("姓名：", s.string, "省份：", s.Province)

	// 3. 内存对齐演示与优化
	// 顺序混乱：内存空洞大
	type Demo1 struct {
		a bool  // 1字节
		b int64 // 8字节
		c bool  // 1字节
	}
	// 顺序优化：大类型在前，减少对齐空洞
	//占用字节大的类型放结构体前面，小字段集中放末尾，减少内存空洞
	//从大到小：int64 → int32 → int16 → bool
	type Demo2 struct {
		b int64
		a bool
		c bool
	}
	fmt.Printf("Demo1 内存大小：%d 字节\n", unsafe.Sizeof(Demo1{}))
	fmt.Printf("Demo2 内存大小：%d 字节\n", unsafe.Sizeof(Demo2{}))
}
