package main

//go modules 是golang1.11新加的特性，用来管理模块中包的依赖关系

//go mod 使用方法
//初始化模块 go mod init <项目名称>

//依赖关系处理，根据go.mod文件 go mod tidy

//将依赖包复制到项目的vendor目录 go mod vendor
//如果包被屏蔽（墙），可以用这个指令，随后使用go build -mod=vendor编译

//显示依赖关系 go list -m all

//显示详细依赖关系 go list -m -json all

//下载依赖  go mod download [path@version](非必写）
