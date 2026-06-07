package ftools

import "fmt"

// RecSum 计算1+2+3+...+n
func RecSum(n int) int {
	// 终止条件：n==1 不再递归
	if n == 1 {
		return 1
	}
	// 递归调用自身
	return n + RecSum(n-1)
}

func Fun9() {
	res := RecSum(5)
	fmt.Println("1~5累加结果：", res)
}

/*
递归定义：函数内部调用自身就是递归。
终止条件（出口）：必须设置结束判断，满足条件直接 return，防止无限递归栈溢出。
执行逻辑：层层递归入栈 → 触发终止条件逐层回溯返回结果。
*/

/*

执行拆解：
RecSum(5)=5+RecSum(4)
RecSum(4)=4+RecSum(3)
RecSum(3)=3+RecSum(2)
RecSum(2)=2+RecSum(1)
RecSum(1)=1（终止）
回溯：2+1 →3+3 →4+6 →5+10 =15
*/
