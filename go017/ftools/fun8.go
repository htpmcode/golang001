package ftools

import "fmt"

// 分数校验：负数触发panic，defer+recover捕获
func AddScore(score int) int {
	// 提前注册defer捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常：", err)
		}
	}()

	if score < 0 {
		panic("分数不能为负数") // 触发崩溃
	}
	fmt.Println("分数合法，加分成功")
	return score
}

func Fun8() {
	//AddScore(60)
	AddScore(-10) // 负数触发panic
	fmt.Println("主函数继续运行，程序未宕机")
}

/*
panic：主动抛出运行异常，终止当前函数后续代码，倒序执行已注册的 defer，若无捕获程序崩溃退出。
recover：只能放在defer包裹的匿名函数内部使用，捕获 panic 错误，阻断崩溃，函数继续向下执行。
使用规范：异常捕获逻辑统一写在函数最开头 defer，提前注册捕获。


没有 recover：panic 后程序直接终止，后面代码不再执行；
recover 必须在 defer 内，写在普通代码位置无法捕获 panic。
*/
