package exa1

import "fmt"

func printValue(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

func Space() {
	printValue(42)          // int
	printValue("hello")     // string
	printValue(3.14)        // float64
	printValue([]int{1, 2}) // slice
}

/*
空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。

任意类型都实现了空接口。
常用于需要存储任意类型数据的场景，如泛型容器、通用参数等。

*/
