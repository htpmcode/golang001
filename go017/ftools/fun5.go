package ftools

import "fmt"

// 生成计数器闭包
func CreateCounter() func() int {
	// cnt是外层局部变量
	cnt := 0
	// 匿名函数捕获cnt，构成闭包
	return func() int {
		cnt++
		return cnt
	}
}

func Fun5() {
	// 第一个独立计数器
	c1 := CreateCounter()
	fmt.Println(c1()) //1
	fmt.Println(c1()) //2
	fmt.Println(c1()) //3

	// 第二个独立计数器，变量cnt全新
	c2 := CreateCounter()
	fmt.Println(c2()) //1
	fmt.Println(c2()) //2
}

/*
闭包本质：匿名函数引用外层函数的局部变量，形成闭包。
生命周期特点：外层函数执行完毕正常销毁局部变量，但被闭包捕获的变量不会回收，生命周期被延长。
数据隔离：每调用一次外层函数，生成独立变量与独立闭包，多个闭包之间数据互不干扰。
常用场景：计数器、状态保存。

CreateCounter函数运行结束本该销毁cnt，但因为匿名闭包持有引用，cnt常驻内存；c1、c2各自绑定专属cnt，计数互不影响。
*/
