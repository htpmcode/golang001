package main

import "fmt"

func main() {
	// ==================== switch 语句 ====================

	// 1. 基础 switch
	fmt.Println("--- 基础 switch ---")
	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("优秀")
	case 'B', 'C': // 多个值匹配同一个 case
		fmt.Println("良好")
	default:
		fmt.Println("差")
	}

	// 2. switch 前置语句
	fmt.Println("\n--- switch 前置语句 ---")
	switch score := 75; {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
	// fmt.Println(score) // 报错：score 作用域仅在 switch 内

	// 3. fallthrough 穿透
	fmt.Println("\n--- fallthrough 穿透 ---")
	num := 1
	switch num {
	case 1:
		fmt.Println("匹配到 1")
		fallthrough // 强制进入下一个 case（不使用 break）
	case 2:
		fmt.Println("穿透到 2")

	case 3:
		fmt.Println("匹配到 3")
	}

	// 4. 无表达式 switch（等价 if-else 链）
	fmt.Println("\n--- 无表达式 switch ---")
	age := 25
	switch {
	case age < 12:
		fmt.Println("儿童")
	case age < 18:
		fmt.Println("青少年")
	case age < 60:
		fmt.Println("成年人")
	default:
		fmt.Println("老年人")
	}

	// 5. break 提前退出 switch
	fmt.Println("\n--- break 提前退出 ---")
	value := 2
	switch value {
	case 1:
		fmt.Println("值是 1")
	case 2:
		fmt.Println("值是 2")
		//break // 显式退出（其实 Go 默认就会退出，这里演示用法）
	case 3:
		fmt.Println("值是 3")
	}
	fmt.Println("switch 结束")

	// 6. 类型 switch（type switch）
	fmt.Println("\n--- 类型 switch ---")
	//var data interface{} = "Hello Go"
	var data interface{} = 12
	switch v := data.(type) {
	case int:
		fmt.Printf("整数：%d\n", v)
	case string:
		fmt.Printf("字符串：%s\n", v)
	case float64:
		fmt.Printf("浮点数：%f\n", v)
	default:
		fmt.Printf("未知类型：%T\n", v)
	}

	// 7. case 中使用表达式
	fmt.Println("\n--- case 表达式 ---")
	x := 10
	switch {
	case x > 0 && x < 5:
		fmt.Println("0-5 之间")
	case x >= 5 && x < 15:
		fmt.Println("5-15 之间")
	case x%2 == 0:
		fmt.Println("偶数")
	}

	// 8. 空 switch（编译错误，演示用）
	// switch {} // 报错：至少需要一个 case 或 default
}
