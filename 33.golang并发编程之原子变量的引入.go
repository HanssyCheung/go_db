package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var i = 100

var lock1 sync.Mutex

func add3() {
	lock1.Lock()
	i++
	lock1.Unlock()
}
func sub() {
	lock1.Lock()
	i--
	lock1.Unlock()
}

//atomic原子操作：主要使用了 cas compare and swap old new 比较和交换，只有旧值和新值做比较没有变化才能进行下一步操作

var i32 int32 = 100

func atomicAdd() {
	atomic.AddInt32(&i32, 1)
}

func atomicSub() {
	atomic.AddInt32(&i32, -1)
}

func main() {
	for i := 0; i < 10000; i++ {
		go add3()
		go sub()
	}
	time.Sleep(time.Second * 3)
	println(i)
}
