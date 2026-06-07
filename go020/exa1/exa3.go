package exa1

import "fmt"

type lab interface{}

type lab1 struct {
}

func (lab1 lab1) Add(a, b int) int {
	return a + b
}

func Sum() {
	l1 := lab1{}
	result := l1.Add(2, 3)
	fmt.Println(result)

}
