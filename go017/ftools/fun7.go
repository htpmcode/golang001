package ftools

import "fmt"

// 1.defer修改命名返回值
func testDefer() (res int) {
	res = 10
	defer func() {
		res = 99 //return赋值完毕，defer修改返回变量
	}()
	return //裸返回，最终返回99
}

func Fun7() {
	//2.多个defer倒序
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")

	fmt.Println("函数主体代码运行")

	ret := testDefer()
	fmt.Println("defer修改后的返回值：", ret)
}

/*
defer 延迟执行：defer注册的语句不在当前立刻执行，函数即将结束 return 前统一执行。
多个 defer 倒序执行（后进先出）：先注册的后执行，后注册的先执行。
执行顺序规则：return 赋值 → defer 执行 → 函数真正返回
命名返回值是函数作用域内变量，defer内部可以修改它，最终返回修改后的值。

*/
