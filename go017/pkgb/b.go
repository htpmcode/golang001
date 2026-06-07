package pkgb

import (
	"17/pkgc"
	"fmt"
)

var BVal = initB()

func initB() int {
	fmt.Println("【B】全局变量初始化，C包值：", pkgc.CVal)
	return 200
}

func init() {
	fmt.Println("【B】init执行")
}
