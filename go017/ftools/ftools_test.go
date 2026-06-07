package ftools

import "testing"

//func TestFun6(t *testing.T) {
//	Fun6()
//}

//func TestFun8(t *testing.T) {
//	Fun7()
//}

//func TestFun8(t *testing.T) {
//	Fun8()
//}

//func TestFun9(t *testing.T) {
//	Fun9()
//}

func TestFun10(t *testing.T) {
	Fun10()
}

//func TestFun11(t *testing.T) {
//	Fun11()
//}

/*

测试文件命名：xxx_test.go
测试函数格式：func TestXxx(t *testing.T)，首字母大写
函数参数固定：t *testing.T
测试代码和源码同包。
*/

/*

//测试Add函数，Test开头大写
func TestAdd(t *testing.T){
	res:=Add(2,3)
	want:=5
	if res!=want{
		//测试失败打印信息
		t.Errorf("Add(2,3) want=%d,got=%d",want,res)
	}
}


*/

/*
1.go test -v -run=TestAdd    只测试单个函数

2.测试所有函数
go test -v

3.常用测试 方法
t.Errorf()：报错，标记用例失败，继续执行后续用例
t.Fatalf()：报错直接终止当前测试函数
*/
