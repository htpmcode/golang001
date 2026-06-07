package main

import (
	"fmt"
	"github.com/gookit/goutil"
	"github.com/kamalyes/go-toolbox/pkg/stringx"
)

func main() {
	// convert type
	str := goutil.String(23)     // "23"
	iVal := goutil.Int("-2")     // 2
	i64Val := goutil.Int64("-2") // -2
	u64Val := goutil.Uint("2")   // 2

	// 输出每个变量的类型和值
	fmt.Printf("str:   type=%T, value=%v\n", str, str)
	fmt.Printf("iVal:  type=%T, value=%v\n", iVal, iVal)
	fmt.Printf("i64Val: type=%T, value=%v\n", i64Val, i64Val)
	fmt.Printf("u64Val: type=%T, value=%v\n", u64Val, u64Val)

	//	github第二个

	// 🧵 字符串处理：hello world → Hello World → Hello Go
	result := stringx.New("hello go")
	fmt.Printf("result: %s\n", result) // 输出: Hello Go

}
