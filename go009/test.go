package main

import "fmt"

func main() {
	//	//7. 标签 break（跳出多层循环）
	//outer:
	//	for i := 0; i < 3; i++ {
	//		for j := 0; j < 3; j++ {
	//			if i == 1 && j == 1 {
	//				fmt.Println("\n跳出外层循环")
	//				break outer
	//			}
	//			fmt.Printf("(%d,%d) ", i, j)
	//		}
	//	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("\n跳出外层循环")
				break
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
}
