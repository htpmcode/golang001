package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	src := " Hello GoLang 中文编程 "
	fmt.Println("原始字符串:[", src, "]")

	// 1. 去首尾空格 Trim
	s1 := strings.TrimSpace(src)
	fmt.Println("Trim去空格:", s1)

	// 2. 前后指定字符裁剪
	s2 := strings.Trim(s1, "H")
	fmt.Println("Trim裁剪H:", s2)

	// 3. 大小写转换
	upper := strings.ToUpper(s1)
	lower := strings.ToLower(s1)
	fmt.Println("大写:", upper)
	fmt.Println("小写:", lower)

	// 4. 包含、前缀、后缀判断
	fmt.Println("包含GoLang:", strings.Contains(s1, "GoLang"))
	fmt.Println("前缀Hello:", strings.HasPrefix(s1, "Hello"))
	fmt.Println("后缀编程:", strings.HasSuffix(s1, "编程"))

	// 5. 查找下标 Index
	idx := strings.Index(s1, "Go")
	fmt.Println("Go起始下标:", idx)
	lastIdx := strings.LastIndex(s1, "g")
	fmt.Println("最后一个g下标:", lastIdx)

	// 6. 字符串替换 Replace
	rep := strings.ReplaceAll(s1, "GoLang", "Golang语言")
	fmt.Println("替换后:", rep)

	// 7. 分割 Split → 切片
	arr := strings.Split(s1, " ")
	fmt.Println("按空格分割切片:", arr)

	// 8. 切片拼接 Join
	joinStr := strings.Join(arr, "|")
	fmt.Println("切片拼接:", joinStr)

	// 9. 重复字符串 Repeat
	repeat := strings.Repeat("A-", 3)
	fmt.Println("重复3次:", repeat)

	// 10. 字符串截取（切片语法，左闭右开）
	sub := s1[6:11]
	fmt.Println("下标6~11截取:", sub)

	// 11. 数字 ↔ 字符串转换
	numStr := "12345"
	num, _ := strconv.Atoi(numStr)
	fmt.Println("字符串转int:", num+1)

	n := 6789
	strN := strconv.Itoa(n)
	fmt.Println("int转字符串:", strN+"666")

	// 12. 统计子串数量 Count
	cnt := strings.Count(s1, "o")
	fmt.Println("字符o出现次数:", cnt)

	// 13. range遍历字符串（适配中文、多字节）
	fmt.Println("\nrune遍历每个字符：")
	for index, char := range s1 {
		if index > 8 {
			break
		}
		fmt.Printf("%d:%c ", index, char)
	}
}
