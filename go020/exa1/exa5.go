package exa1

import "fmt"

// 空姐口断言
func Exa5() {

	var i interface{} = 42

	if str, ok := i.(string); ok {
		fmt.Println("String:", str)
	} else {
		fmt.Println("Not a string")
	}

}

type Persion struct {
	Name string
	Age  int
}

func Basic(val interface{}) {

	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)

	case []Persion:
		fmt.Println("Persons:", v)
	default:
		fmt.Println("Unknown type")
	}

}

func CallBasic() {
	//验证各种类型
	Basic(42)
	Basic("Hello")
	Basic(3.14)
	p1 := Persion{"Tom", 20}
	Basic([]Persion{p1})

}
