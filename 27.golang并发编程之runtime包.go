package main

import (
	"fmt"
	"runtime"
	"time"
)

//runtime包里面定义了一些写成管理相关的api

//runtime.Gosched() 让出CPU时间，重新等待安排任务

func showGoSched(s string) {
	for i := 0; i < 2; i++ {
		println(s)
	}
}

//runtime.Goexit() 退出当前协程
func showGoExit() {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit()
		}
		fmt.Printf("i:%v\n", i)
	}
}

//runtime.GOMAXPROCS 最大的cpu核心数（1.5之前是默认使用1个，之后是使用全部的）
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	go showGoSched("java")
	//主协程
	for i := 0; i < 2; i++ {
		//切一下，再次分配任务
		runtime.Gosched() //让出cpu不代表每次都让出，如果正在执行则不能让，只能在下一次之前让一下，如果主协程任务少可能一下子就先执行完了
		fmt.Println("golang")
	}

	go showGoExit()
	time.Sleep(time.Second)

	fmt.Printf("runtime.NumCPU():%v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2) //修改为1查看效果
	go a()
	go b()
	time.Sleep(time.Second)
}
