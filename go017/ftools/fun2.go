package ftools

import "fmt"

type User struct {
	Name  string
	Score int
}

// 1.普通参数 + 2.可变参数
func Sum(base int, nums ...int) int {
	total := base
	for _, v := range nums {
		total += v
	}
	return total
}

// 指针传参：修改原结构体数据
func ChangeScore(u *User, newScore int) {
	u.Score = newScore
}

func fun2() {
	// 普通实参调用可变参数
	r1 := Sum(10, 1, 2, 3, 10)
	fmt.Println("常规传参结果：", r1)

	// 3.切片打散 arr...
	slice := []int{5, 6, 7}
	//r2 := Sum(20, slice...)
	r2 := Sum(20, slice...)
	fmt.Println("切片打散结果：", r2)

	// 4.指针传参
	u := User{Name: "小明", Score: 50}
	ChangeScore(&u, 99)
	//ChangeScore(u, 99)
	fmt.Println("指针修改后：", u)
}

/*
、知识点总结
普通参数：同类型参数可简写 a,b int；值类型实参为值拷贝，函数内修改副本不影响原值。
可变参数 nums ...int：可变参数必须放在形参最后一位，底层本质是 []int 切片，支持 0 个或多个实参传入。
切片打散 arr...：已有切片变量通过 切片名... 打散，逐个传入可变参数函数。
指针传参：传入变量地址，函数内通过指针直接操作原始内存，修改会影响外部原变量（结构体 / 数组常用）。

*/
