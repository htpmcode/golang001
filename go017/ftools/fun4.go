package ftools

import (
	"fmt"
	"time"
)

func Fun4() {
	// 1.自执行匿名函数
	func(a, b int) {
		fmt.Printf("自执行：%d+%d=%d\n", a, b, a+b)
	}(3, 5)

	// 2.匿名函数赋值变量
	f := func(x int) int {
		return x * x
	}
	fmt.Println("变量调用：", f(6))

	//3.协程匿名函数
	go func() {
		fmt.Println("协程异步运行")
	}()

	time.Sleep(time.Millisecond * 20)
	//time.Sleep(time.Second *2)
}

/*
知识点总结
匿名函数：无函数名称的函数，定义格式 func(入参)(返回){}。
自执行匿名函数：定义末尾加()，声明完立刻运行，func(){代码}()。
赋值给变量：把匿名函数存到变量，后续通过变量名反复调用。
协程匿名函数：go func(){}()，开启 goroutine 异步执行。
*/
