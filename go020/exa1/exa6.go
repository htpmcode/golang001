package exa1

import "fmt"

// 写法都是带有er后缀的接口，
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (f File) Read() string {
	return "Reading data"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}

func Exa6() {
	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("Hello, Go!")
}
