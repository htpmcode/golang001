package exa1

import "fmt"

func Exa7() {
	var i interface{} = 42
	fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)
}

/*
 接口变量实际上包含了两部分：

    动态类型：接口变量存储的具体类型。
    动态值：具体类型的值。

动态值和动态类型示例：

*/

//接口的零值
//
//接口的零值是 nil。
//
//当接口变量的动态类型和动态值都为 nil 时，接口变量为 nil。

func Labnil() {
	var i interface{}
	fmt.Println(i == nil) // 输出：true
}
