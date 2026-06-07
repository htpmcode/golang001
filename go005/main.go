package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Go语言"
	// 字节长度
	fmt.Println("字节长度：", len(s))

	// 索引取字节
	fmt.Println("首字节：", s[0])

	// 字符串切片
	sub := s[0:2]
	fmt.Println("切片：", sub)

	// 拼接
	s1 := "Hello" + " Go"
	s2 := fmt.Sprintf("num=%d==%v", 666, true)
	fmt.Println(s1, s2)

	// 高性能拼接 Builder
	var builder strings.Builder
	builder.WriteString("a")
	builder.WriteString("b")
	fmt.Println(builder.String())

	// strings工具
	fmt.Println(strings.Contains(s, "语"))
	fmt.Println(strings.Split("1,2,3", ","))
	fmt.Println(strings.Join([]string{"x", "y"}, "-"))
	fmt.Println("========================")
	fmt.Println(strings.ReplaceAll("xxxxxx", "x", "y"))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.Trim("   xxx   ", " "))

	fmt.Println(strings.HasPrefix(s, "G"))
	fmt.Println(strings.HasSuffix(s, "语言"))
	fmt.Println(strings.Index(s, "语言"))
	fmt.Println(strings.LastIndex(s, "语言"))
	fmt.Println(strings.ReplaceAll(s, "Go", "C"))

	// strconv 数字字符串互转
	strNum := "123"
	n, _ := strconv.Atoi(strNum)
	fmt.Println("字符串转数字：", n)
	strOut := strconv.Itoa(456)
	fmt.Println("数字转字符串：", strOut)
}
