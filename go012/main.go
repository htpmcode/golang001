package main

import "fmt"

func sliceChange(s []int) {
	s[0] = 888          // 修改底层数组，外部生效
	s = append(s, 9999) // 扩容，指向新数组，后续修改不影响原切片
}

func main() {
	//1.从数组截取切片
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[1:4]   // low=1 high=4 len=3 cap=9
	s2 := arr[1:4:6] // low=1 high=4 max=6 len=3 cap=5
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	//2.字面量创建切片
	s3 := []int{11, 22, 33}
	fmt.Println(s3)

	//3.make创建切片
	s4 := make([]int, 3, 6) // len=3 cap=6，默认零值[0,0,0]
	fmt.Println(s4)

	//nil切片 vs 空切片
	var nilS []int
	emptyS := []int{}
	fmt.Println(nilS == nil, emptyS == nil) //true false

	//append追加与扩容
	s5 := make([]int, 2, 3)
	s5 = append(s5, 10) // cap用完 len=3 cap=3
	s5 = append(s5, 20) // 触发扩容 cap=6
	fmt.Println(len(s5), cap(s5))

	//浅拷贝：共用底层
	src := []int{1, 2, 3}
	shallow := src
	shallow[0] = 99
	fmt.Println(src) //[99 2 3]

	//深拷贝：copy+make
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 100
	fmt.Println(src, dst)

	//切片传参测试
	testS := []int{1, 2, 3}
	sliceChange(testS)
	fmt.Println(testS) //[888 2 3] append扩容不影响原切片
}
