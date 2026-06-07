package ftools

import "fmt"

// 1.自定义Calc函数类型
type Calc func(int) int

// 2.高阶函数：第二个参数为Calc函数类型
func CalcHandle(num int, f Calc) int {
	return f(num)
}

// 普通函数，符合Calc签名
func Double(n int) int {
	return n * 2
}

func Fun6() {
	//方式1：传入已定义函数
	res1 := CalcHandle(10, Double)
	fmt.Println("翻倍：", res1)

	//方式2：直接传入匿名函数
	res2 := CalcHandle(10, func(n int) int {
		return n + 5
	})
	fmt.Println("加5：", res2)
}

/*
自定义函数类型：type 类型名 func(入参列表)返回值，把一类入参、返回一致的函数抽象成自定义类型。
高阶函数定义：可以接收函数作为参数，或把函数作为返回值的函数。
使用规则：凡是满足 Calc 签名格式的普通函数、匿名函数，都能作为实参传入高阶函数。

*/
