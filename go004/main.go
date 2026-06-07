package main

import "fmt"

func main() {
	// bool
	var flag bool = true
	fmt.Println("bool:", flag)

	// 整数、无符号、byte
	var i64 int64 = 999
	var bt byte = 'A'
	fmt.Println("int64:", i64, "byte字符A:", string(bt))

	// rune 中文unicode
	var ru1 rune = '中'
	var ru2 rune = '国'
	fmt.Println("rune中文：", string(ru1), string(ru2))

	// 浮点
	var f float64 = 2.718
	fmt.Println("float64:", f)
	fmt.Printf("float64:%.2f\n", f)

	// 显式类型转换
	var a int = 5
	var b float64 = 3.9
	sum := float64(a) + b
	trunc := int(b) // 截断小数
	fmt.Println("转换求和：", sum, "截断小数：", trunc)
}
