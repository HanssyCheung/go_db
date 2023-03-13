package main

import (
	"fmt"
	"sync"
)

//同步：两个协程之间相互等待，一个没执行完，另一个要等待

//添加WaitGroup和不添加WaitGroup的区别（协程中提供的函数）
//将线程放进group里面，等待执行完毕通知另一个线程
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() //goroutine结束就登记-1
	//fmt.Println("hello Goroutine", 1)
	fmt.Printf("i:%v\n", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //启动一个goroutine
		go hello(i)
	}
	fmt.Println("go") //主协程会等其他协程结束再结束，但并不代表主协程不先执行
	wg.Wait()         //等待所有登记的goroutine都结束
}
