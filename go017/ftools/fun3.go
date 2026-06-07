package ftools

import "fmt"

// 1.多返回：(结果,错误)
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// 2.命名返回 + 裸return
func Max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}
	return // 裸return，自动return res
}

func Fun3() {
	// 常规接收双返回
	val, err := Div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("除法结果：", val)
	}

	// 3._忽略不需要的第二个返回值error
	val2, _ := Div(20, 4)
	fmt.Println("忽略错误：", val2)

	// 调用命名返回函数
	fmt.Println("最大值：", Max(12, 25))
}

/*
知识点总结
多返回值：Go 支持一个函数返回多个结果，业务标准写法 (返回数据, error)，用来返回结果 + 错误信息。
命名返回值（裸 return）：在函数声明处预先定义返回变量名，函数内直接使用变量，return 后面不写值（裸返回），自动返回已赋值的命名变量。
空白标识符 _：接收多返回值时，不需要的返回项用下划线接收，直接丢弃数据。
*/
