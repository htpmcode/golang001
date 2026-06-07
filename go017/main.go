package main

import (
	"17/pkgb"
	"fmt"
)

var mainVal = initMain()

func initMain() int {
	fmt.Println("【main】全局变量初始化，B包值：", pkgb.BVal)
	return 999
}

func init() {
	fmt.Println("【main】init执行")
}

func main() {

	//add := ftools.Add(10, 5)
	////
	//fmt.Println("加法：", add)

	//sub := ftools.sub(10, 5)

	fmt.Println("====main函数启动，程序运行====")
	fmt.Println("main全局变量 =", mainVal)

}

/*
关于init事项

知识点总结（多包、多 init 完整执行规则）
init()特性：无参无返回、禁止手动调用，包被导入时自动执行。
单包规则：同一个.go文件多个init()从上往下执行；同一包多个文件，按文件名字典序依次执行各文件init。
跨包导入执行顺序：
导入包全局变量初始化 → 导入包所有init → 当前包全局变量初始化 → 当前包所有init → main()
导入依赖链式：A 导入 B、B 导入 C：C全局→C init → B全局→B init → A全局→A init → main。
空白导入import _ 包路径：仅执行目标包 init，不暴露包内标识符，常用于驱动注册。

*/

/*

【C】全局变量初始化
【C】第一个init执行
【C】第二个init执行
【B】全局变量初始化，C包值： 100
【B】init执行
【main】全局变量初始化，B包值： 200
【main】init执行
====main函数启动，程序运行====
main全局变量 = 999
*/

/*


 */
