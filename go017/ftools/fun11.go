package ftools

import (
	"fmt"
	"time"
)

// 普通函数
func Test(msg string) {
	fmt.Println("协程执行：", msg)
}

func Fun11() {
	//1.普通函数开启协程
	go Test("普通函数协程")

	//2.匿名函数开启协程
	go func(s string) {
		fmt.Println("匿名协程：", s)
	}("传入参数")

	//主线程等待子协程执行完毕
	time.Sleep(time.Millisecond * 50)
	fmt.Println("main主线程结束")
}

/*
goroutine 开启格式：go 函数/匿名函数()，开启子协程异步执行函数，主协程 (main) 和子协程并发运行。
main 是主线程：main 函数执行完毕程序直接退出，未跑完的子协程会被强制终止，可用time.Sleep临时等待。
支持普通命名函数、匿名函数两种形式启动协程；匿名协程最常用。
go 关键字只负责启动，不阻塞后续代码。

*/
