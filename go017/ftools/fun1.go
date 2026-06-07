package ftools

import "fmt"

// 1.大写导出函数：跨包可访问
func Add(a, b int) int {
	return a + b
}

// 2.小写私有函数：仅本包可用
func sub(a, b int) int {
	return a - b
}

// main：程序入口
func f1() {
	res1 := Add(10, 5)
	res2 := sub(10, 5)
	fmt.Println("加法：", res1)
	fmt.Println("减法：", res2)

	// 错误：Go不支持重载，不能重复定义同名Add函数
	// func Add(a int)int{return a}
}

/*
函数语法格式
go
运行
func 函数名(形参列表)(返回值列表){函数体}
func是函数声明关键字。
2. 访问权限（首字母大小写）
函数名大写开头：导出函数，跨包 import 后可以调用；
函数名小写开头：私有函数，仅当前包内部使用，外部包无法访问。
main () 特殊入口函数
func main(){}，无参数、无返回值，程序唯一执行入口，不能手动调用。
Go 函数规则：不支持函数重载、没有默认参数
同一个包内不能定义同名函数（形参不同也不行），不能设置参数默认值。
补充：结构体指针传参修改原数据，值传递仅拷贝副本。

*/
