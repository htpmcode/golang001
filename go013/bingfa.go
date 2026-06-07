package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup
	// 并发写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sm.Store(n, n*10)
		}(i)
	}
	wg.Wait()

	// 查询
	val, ok := sm.Load(3)
	if ok {
		fmt.Println(val)
	}
	// 遍历
	sm.Range(func(k, v any) bool {
		fmt.Printf("k:%v v:%v\n", k, v)
		return true
	})
}
