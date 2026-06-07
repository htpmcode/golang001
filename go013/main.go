package main

import (
	"fmt"
	"sort"
)

func main() {
	//=====================1.四种创建map方式=====================
	var nilMap map[string]int // nil map，未分配内存
	fmt.Printf("nilMap==nil：%t\n", nilMap == nil)
	// nilMap["a"] = 1 // 取消注释直接panic

	mapLiteral := map[string]int{"张三": 18, "李四": 22} //字面量初始化
	//emptyMap := make(map[int]string)                 //空map，已分配内存
	preMap := make(map[string]float64, 10) //预分配容量，优化扩容
	preMap["语文"] = 95.5
	fmt.Println("字面量map:", mapLiteral)

	//=====================2.增删改查操作=====================
	userAge := make(map[string]int, 5)
	//新增
	userAge["小明"] = 18
	userAge["小红"] = 20
	//修改
	userAge["小明"] = 19
	//查询（必须ok判断）
	if age, ok := userAge["小红"]; ok {
		fmt.Printf("小红年龄：%d\n", age)
	} else {
		fmt.Println("用户不存在")
	}
	//删除
	delete(userAge, "小红")
	fmt.Println("删除后长度len：", len(userAge))

	//=====================3.三种range遍历 + key排序实现有序遍历=====================
	score := map[string]int{"语文": 90, "数学": 95, "英语": 88}
	fmt.Println("原始无序遍历：")
	for k, v := range score {
		fmt.Printf("%s:%d ", k, v)
	}

	//提取key排序，实现有序输出
	keySlice := make([]string, 0, len(score))
	for k := range score {
		keySlice = append(keySlice, k)
	}
	sort.Strings(keySlice)
	fmt.Println("\n排序后有序遍历：")
	for _, k := range keySlice {
		fmt.Printf("%s:%d ", k, score[k])
	}

	//=====================4.map引用特性：浅拷贝、函数传参修改原map=====================
	fmt.Println("\n====引用测试====")
	funcModify := func(m map[string]int) {
		m["新增字段"] = 100
	}
	m1 := map[string]int{"a": 1}
	m2 := m1 //浅拷贝，共用底层哈希表
	m2["a"] = 99
	fmt.Println("m1:", m1)
	funcModify(m1)
	fmt.Println("传参修改后m1:", m1)

	//map深拷贝（手动新建循环赋值）
	mCopy := make(map[string]int, len(m1))
	for k, v := range m1 {
		mCopy[k] = v
	}
	mCopy["a"] = 999
	fmt.Println("深拷贝后原m1不受影响：", m1, "拷贝mCopy：", mCopy)

	//=====================5.map嵌套（value为map，业务常用）=====================
	fmt.Println("\n====嵌套Map====")
	studentScore := make(map[string]map[string]int)
	//内层map必须手动make初始化
	studentScore["小明"] = make(map[string]int)
	studentScore["小明"]["数学"] = 98
	studentScore["小明"]["语文"] = 92
	fmt.Println("嵌套map：", studentScore)

	//=====================6.map经典用法：切片去重（空结构体不占内存）=====================
	fmt.Println("\n====切片去重====")
	originArr := []int{1, 2, 2, 3, 3, 3, 5}
	uniqueMap := make(map[int]struct{})
	for _, val := range originArr {
		uniqueMap[val] = struct{}{}
	}
	//key转回切片
	resSlice := make([]int, 0, len(uniqueMap))
	for k := range uniqueMap {
		resSlice = append(resSlice, k)
	}
	fmt.Println("去重结果：", resSlice)

	//=====================7.易错点演示：nil与空map、禁止&m[key]=====================
	var testNil map[int]int
	emptyTest := make(map[int]int)
	fmt.Printf("\nnil map len=%d, 空map len=%d\n", len(testNil), len(emptyTest))
	// _ = &emptyTest[1] // 取消注释：编译报错，不能取map元素地址
}
