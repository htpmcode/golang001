module go018

go 1.25.10

require (
	github.com/gookit/goutil v0.7.6
	github.com/kamalyes/go-toolbox v0.15.3
)

//间接引用
require (
	github.com/kamalyes/go-argus v0.2.1 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/text v0.37.0 // indirect
)

//排除某个版本
//exclude github.com/kamalyes/go-toolbox v0.15.3

//替换某个版本// 写法2：远端模块 → 指定版本的官方依赖
//go list -m github.com/gookit/goutil
replace github.com/gookit/goutil => github.com/gookit/goutil v0.7.2
