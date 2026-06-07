package main

import (
	"fmt"
	"time"
	"unsafe"
)

// Logger 类型定义（包级别）
type Logger struct {
}

func (l Logger) Info(msg string) {
	fmt.Println("[INFO]", msg)
}

func (l Logger) Error(msg string) {
	fmt.Println("[ERROR]", msg)
}

func main() {
	// 1. struct{} 零大小特性
	var a struct{}
	fmt.Println("空结构体大小：", unsafe.Sizeof(a)) // 输出 0

	// 2. 实现 Set 集合（Go 没有原生 Set，用 map 模拟）
	set := make(map[string]struct{})
	// 添加元素
	//
	set["张三"] = struct{}{} // 创建一个空结构体的值，作为 map 的 value
	set["李四"] = struct{}{}
	// 判断存在
	_, ok := set["张三"]
	fmt.Println("张三是否存在：", ok)
	_, ok = set["王五"]
	fmt.Println("王五是否存在：", ok)
	// 删除元素
	delete(set, "张三")
	_, ok = set["张三"]
	fmt.Println("删除后张三是否存在：", ok)

	// 3. 通道做信号通知（只发信号，不带数据）
	done := make(chan struct{})

	go func() {
		fmt.Println("后台任务执行中,需要等待...")
		// 任务执行完发送信号
		time.Sleep(5 * time.Second)
		done <- struct{}{}
	}()

	<-done // 阻塞等待任务结束
	fmt.Println("任务结束")

	// 4. 结构体内嵌占位、仅实现方法、标识属性

	// 示例1：内嵌空结构体作为标记
	type Config struct {
		Name    string
		Enabled bool
		_       struct{} // 占位符
	}

	cfg := Config{
		Name:    "测试配置",
		Enabled: true,
	}
	fmt.Printf("配置信息：%+v, 大小：%d\n", cfg, unsafe.Sizeof(cfg))

	// 示例2：仅用于实现方法的类型
	logger := Logger{}
	logger.Info("系统启动")
	logger.Error("发生错误")

	// 示例3：多 goroutine 同步
	workers := 3
	allDone := make(chan struct{})

	for i := 1; i <= workers; i++ {
		go func(id int) {
			fmt.Printf("Worker %d 开始工作\n", id)
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("Worker %d 任务完成\n", id)
			allDone <- struct{}{}
		}(i)
	}

	// 等待所有 worker 完成
	for i := 0; i < workers; i++ {
		<-allDone
	}
	fmt.Println("所有 Worker 完成任务")
}

//
//1. struct{} - 类型
//这是类型声明，表示"空结构体类型"
//
//2. struct{}{} - 值
//这是空结构体的值，大小为0字节，用于占位。
//第一个 struct{} → 类型
//第二个 {} → 字面量初始化（虽然是空的）

//set["张三"] = struct{}{}  // 创建一个空结构体的值，作为 map 的 value
//done <- struct{}{}        // 发送一个空结构体的值到通道

////=================
//总结
//struct{} → 类型（Type）
//struct{}{} → 值（Value），是 struct{} 类型的实例
//当需要使用值的地方（赋值、发送通道等），就要用 struct{}{}
//当需要声明类型的地方（变量声明、类型定义），就用 struct{}
//
