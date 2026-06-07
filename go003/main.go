package main

import (
	"fmt"
	"unsafe"
)

// 枚举常量：性别类型
const (
	Unknown = iota // 0
	Female         // 1
	Male           // 2
)

// 字符串相关常量
const (
	str     = "abc"
	strLen  = len(str)
	strSize = unsafe.Sizeof(str)
	//
)

//在计算的是 Go 字符串类型（string）的运行时表示结构的大小，

// iota混合赋值演示
const (
	A = iota   // 0
	B          // 1
	C = "test" // 固定值
	D          // 继承"test"
	E = 100    // 固定值
	F          // 继承100
	FF
	G = iota // 6 (iota当前值为7)
	H
	HH
)

func main() {
	const LENGTH int = 10
	const WIDTH = 5
	const a, b, c = 20, true, "go语言"

	area := LENGTH * WIDTH

	fmt.Println("===== 基础常量用法 =====")
	fmt.Printf("矩形面积：%d\n", area)
	fmt.Println("多类型常量：", a, b, c)

	fmt.Println("\n===== iota枚举 =====")
	fmt.Printf("未知：%d 女性：%d 男性：%d\n", Unknown, Female, Male)

	fmt.Println("\n===== 常量内置函数 =====")
	fmt.Printf("字符串：%s 长度：%d 内存大小：%d字节\n", str, strLen, strSize)

	fmt.Println("\n===== iota混合赋值 =====")
	fmt.Printf("A=%d B=%d C=%s D=%s E=%d F=%d FF=%d G=%d H=%d HH=%d\n", A, B, C, D, E, F, FF, G, H, HH)

}
