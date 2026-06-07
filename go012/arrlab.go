package main

import "fmt"

func arrCopyTest(arr [3]int) {
	arr[0] = 99 // 值拷贝，修改副本不影响原数组
}

func arrPointTest(arr *[3]int) {
	arr[0] = 99 // 指针修改原数组
}

func main() {
	// 写法1：默认零值初始化
	var arr1 [4]int
	fmt.Println(arr1) // [0 0 0 0]

	// 写法2：字面量全赋值
	arr2 := [3]int{10, 20, 30}
	fmt.Println(arr2)

	// 写法3：部分赋值，剩余零填充
	arr3 := [5]int{1, 2}
	fmt.Println(arr3) // [1 2 0 0 0]

	// 写法4：...自动推导长度
	arr4 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(len(arr4)) //5

	// 数组传参：值拷贝测试
	arrParam := [3]int{1, 2, 3}
	arrCopyTest(arrParam)
	fmt.Println(arrParam) //[1 2 3] 原值不变
	arrPointTest(&arrParam)
	fmt.Println(arrParam) //[99 2 3] 指针修改生效

	// 下标遍历
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("%d ", arr2[i])
	}
	fmt.Println()
	// range遍历
	for idx, val := range arr2 {
		fmt.Printf("下标:%d,值:%d\n", idx, val)
	}

	// 二维数组
	var twoArr [2][3]int
	twoArr = [2][3]int{{1, 2}, {3, 4}}
	fmt.Println(twoArr)
}
