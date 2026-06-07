package exa1

import "fmt"

// 定义 Person 结构体
type Person struct {
	Name string
	Age  int
}

func IfaceSlice() {
	// 1. 定义空接口切片，存放任意类型（这里存 Person 结构体）
	//只要是定义interface的内容  后期都是要断言的，否则不能使用
	var data []interface{}

	// 往切片添加 Person 实例
	data = append(data, Person{Name: "张三", Age: 20})
	data = append(data, Person{Name: "李四", Age: 25})

	// 2. 遍历切片，使用 类型断言 取出 Person
	for idx, val := range data {
		// 类型断言：val.(Person)，ok 标记断言是否成功
		p, ok := val.(Person)
		if ok {
			fmt.Printf("第%d个元素：姓名=%s，年龄=%d\n", idx+1, p.Name, p.Age)
		} else {
			fmt.Printf("第%d个元素不是 Person 类型\n", idx+1)
		}
	}
	s1 := data[0]
	s2 := data[1]
	fmt.Println("切片第一个元素：", s1)
	fmt.Println("切片第二个元素：", s2)

	//	是否可以获取切片中关于interface{}的元素呢
	//name := s1.Name
	//fmt.Println("name:", name)

	name, ok := s1.(Person)
	if ok {
		fmt.Println("name:", name.Name)
	}
}
