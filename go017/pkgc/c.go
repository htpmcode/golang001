package pkgc

import "fmt"

var CVal = initC()

func initC() int {
	fmt.Println("【C】全局变量初始化")
	return 100
}

// 包内两个init，从上到下执行
func init() {
	fmt.Println("【C】第一个init执行")
}
func init() {
	fmt.Println("【C】第二个init执行")
}
