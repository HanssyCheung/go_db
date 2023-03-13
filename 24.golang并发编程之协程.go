package main

import (
	"fmt"
	"time"
)

//Golang中并发是函数相互独立运行的能力。Goroutines是并发运行的函数。Golang提供了Goroutines作为并发处理操作的一种方式

//创建协程方式,使用go关键字  go task()

func showMsg(msg string) {
	for i := 1; i < 5; i++ {
		fmt.Printf("msg:%v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showMsg("java") //启动了一个协程来执行，main函数里必须保证有主函数在执行，否则主协程结束，其他协程就结束了
	showMsg("golang")
}
