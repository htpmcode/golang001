package main

import "fmt"

func main() {
	// 1. var声明，零值
	var a int
	var s string
	var b bool
	fmt.Println("零值：", a, s, b) // 0 "" false

	// 2. 声明并赋值，类型推断
	var num int = 100
	var str = "go"
	fmt.Println(num, str)
	fmt.Println("=========")
	// 3. 短变量 :=
	age := 22
	fmt.Println(age)

	// 4. 单行多变量
	var x, y int = 1, 2
	m, n := "abc", 3.14
	fmt.Println(x, y, m, n)

	// 5. 块级批量声明
	var (
		floatNum float64 = 100.12
		intNum   int     = 199

		ptr = &intNum
	)
	fmt.Println("float零值：", floatNum, "指针零值：", ptr)

	//     6.格式化字符串

	// Go 语言中使用 fmt.Sprintf 或 fmt.Printf 格式化字符串并赋值给新串：

	//     Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
	//     Printf 根据格式化参数生成格式化的字符串并写入标准输出。
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var flag = false
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s===>%v"
	targetURL := fmt.Sprintf(url, stockcode, enddate, flag)
	fmt.Println(targetURL)

}
