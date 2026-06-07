package ftools

import "fmt"

// init自动初始化全局map
func init() {
	fmt.Println("init执行：初始化全局Map")
	userMap["小明"] = 90
	userMap["小红"] = 95
}

// 全局map，特意放在这里，就是为了验证全局变量初始化，正常应该放在最上面
var userMap = make(map[string]int)

func Fun10() {
	fmt.Println("main启动，读取map数据：", userMap)
}

/*

func init(){} 无参数、无返回值，不能手动调用，包被加载时自动执行。
执行优先级：全局变量初始化 → init 函数 → main 函数。
单个包可定义多个init()，按代码从上到下顺序执行；常用于全局资源初始化（map、连接、配置）。
*/
