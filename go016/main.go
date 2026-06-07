package main

import (
	"encoding/json"
	"fmt"
)

// 嵌套子结构体，添加json tag
type Address struct {
	City string `json:"city"` // json字段名叫city
	area string `json:"area"` // 小写私有：序列化直接丢失
}

// 主结构体，全部配置json tag
type Person struct {
	Name    string        `json:"username"`      // json别名username
	Age     int           `json:"age,omitempty"` // omitempty：age=0则json不显示该字段
	Address `json:"addr"` // 匿名内嵌，json嵌套到addr对象里
	Weight  int           `json:"-"` // "-"：永久忽略，不会序列化
}

// 沿用之前的方法（保留前两节内容）
func (a *Address) ModifyCity(newCity string) {
	a.City = newCity
}
func (p Person) ShowInfo() {
	fmt.Printf("姓名:%s,年龄:%d,城市:%s\n", p.Name, p.Age, p.City)
}
func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func main() {
	// 1.构造结构体数据
	p1 := Person{
		Name: "张三",
		Age:  20,
		Address: Address{
			City: "北京",
			area: "海定", // 小写私有，无法序列化
		},
		Weight: 68, // tag:"-" 直接忽略
	}

	// ========== ① 序列化：结构体 → JSON []byte ==========
	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}
	fmt.Println("序列化JSON：", string(jsonData))
	/*输出：{"username":"张三","age":20,"addr":{"city":"西安"}}
	  说明：area(小写)、Weight(tag:-) 全部消失
	*/

	// 测试omitempty：Age=0，序列化自动删掉age字段
	p2 := Person{Name: "空年龄用户", Address: Address{City: "北京"}}
	json2, _ := json.Marshal(p2)
	fmt.Println("age为0省略：", string(json2))

	// ========== ② 反序列化：JSON字符串 → 结构体 ==========
	jsonStr := `{"username":"李四","age":25,"addr":{"city":"成都"}}`
	var p3 Person
	err = json.Unmarshal([]byte(jsonStr), &p3)
	if err != nil {
		fmt.Println("反序列化失败：", err)
		return
	}
	fmt.Println("反序列化结构体：", p3)
	p3.ShowInfo()

	// ========== ③ 匿名结构体快速序列化（临时JSON） ==========
	tempJson, _ := json.Marshal(struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}{"成功", 200})
	fmt.Println("匿名结构体JSON：", string(tempJson))
}
