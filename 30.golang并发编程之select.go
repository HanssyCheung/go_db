package main

import (
	"fmt"
	"time"
)

/*
1.select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中channel的读写操作，当case
中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作
    select中的case语句必须是一个channel操作
	select中的default子句总是可运行的
2.如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行
3.如果没有可运行地case语句，且有default语句，那么就会执行default的动作
4.如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行

*/

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		fmt.Printf("chanInt23:%p\n", &chanInt)
		chanInt <- 100
		chanStr <- "hello"
		close(chanInt)
		close(chanStr)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt:%v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr:%v\n", r)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
