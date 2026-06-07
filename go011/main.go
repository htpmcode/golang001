package main

import "fmt"

func main() {
	// 方法1：基础嵌套循环
	fmt.Println("=== 99乘法表 ===")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%2d  ", j, i, i*j)
		}
		fmt.Println()
	}

	// 方法2：使用前置语句
	fmt.Println("\n=== 使用前置语句 ===")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if result := i * j; result < 10 {
				fmt.Printf("%d×%d= %d  ", j, i, result)
			} else {
				fmt.Printf("%d×%d=%d  ", j, i, result)
			}
		}
		fmt.Println()
	}

	// 方法3：倒序版本
	fmt.Println("\n=== 倒序99乘法表 ===")
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%2d  ", j, i, i*j)
		}
		fmt.Println()
	}

	// 方法4：完整矩阵版（9×9）
	fmt.Println("\n=== 完整矩阵版 ===")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d×%d=%2d  ", j, i, i*j)
		}
		fmt.Println()
	}
}
