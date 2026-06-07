package main

import "fmt"

// 函数：通过指针修改外部变量
func changeNum(p *int) {
	// 解引用，修改原变量值
	*p = 888
}

// 定义结构体
type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. 基础用法：取地址 & 、解引用 *
	var num int = 100
	var p *int // 声明int类型指针，零值为 nil

	p = &num // &num：获取num的内存地址，赋值给指针p
	fmt.Println("num 的内存地址：", p)
	fmt.Println("指针指向的值：", *p) // *p：解引用取值

	// 通过指针修改原变量
	*p = 200
	fmt.Println("修改后 num =", num)

	// 2. nil 指针与判空
	var nilPtr *string
	fmt.Println("\n空指针值：", nilPtr)
	if nilPtr == nil {
		fmt.Println("当前是空指针，禁止解引用")
	}
	// *nilPtr  // 解除注释会触发 panic

	// 3. new() 在堆上分配内存，返回指针
	heapPtr := new(int)
	*heapPtr = 666
	fmt.Println("\nnew 创建堆内存指针，值：", *heapPtr)

	// 4. 指针作为函数参数，修改变量
	var a int = 10
	fmt.Println("调用函数前 a =", a)
	changeNum(&a) // 传入变量地址
	fmt.Println("调用函数后 a =", a)

	// 5. 结构体指针 + 语法糖
	var per Person = Person{Name: "张三", Age: 20}
	perPtr := &per
	// 标准写法：(*perPtr).Name
	// 语法糖简写：直接 perPtr.Name
	perPtr.Age = 21
	fmt.Printf("\n结构体指针访问：姓名=%s，年龄=%d\n", perPtr.Name, perPtr.Age)

	// 6. 验证：Go 不支持指针运算
	// p++  // 编译报错：invalid operation: p++ (non-numeric type *int)
}
